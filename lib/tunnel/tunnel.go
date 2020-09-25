package tunnel

import (
	"golang.org/x/crypto/ssh"
	"../config"
	"io/ioutil"
	"../port"
	"fmt"
	"log"
	"net"
	"io"
)

type Endpoint struct {
	Host string
	Port port.Port
}

type Conn net.Conn

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

func handleClient(client Conn, remote Conn) {
	defer client.Close()
	chDone := make(chan bool)

	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(client, remote)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy remote->local: %s", err))
		}
		chDone <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(remote, client)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy local->remote: %s", err))
		}
		chDone <- true
	}()

	<-chDone
}

func publicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Cannot read SSH private key file %s... %s", file, err.Error()))
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Cannot parse SSH private key file %s... %s", file, err.Error()))
		return nil
	}
	return ssh.PublicKeys(key)
}

func Open(localPort port.Port, remoteEndpoint Endpoint) {
	// local service to be forwarded
	var localEndpoint = Endpoint{
		Host: "127.0.0.1",
		Port: localPort,
	}

	// refer to https://godoc.org/golang.org/x/crypto/ssh for other authentication types
	sshConfig := &ssh.ClientConfig{
		// SSH connection username
		User: config.GetString("user"),
		Auth: []ssh.AuthMethod{
			publicKeyFile(config.GetString("id_rsa_pub_path")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to SSH remote server using remoteEndpoint
	serverConn, err := ssh.Dial("tcp", remoteEndpoint.String(), sshConfig)
	if err != nil {
		log.Fatalln(fmt.Printf("Dial INTO remote server error: %s", err))
	}

	// Listen on remote server port
	listener, err := serverConn.Listen("tcp", remoteEndpoint.String())
	if err != nil {
		log.Fatalln(fmt.Printf("Listen open port ON remote server error: %s", err))
	}
	defer listener.Close()

	// handle incoming connections on reverse forwarded tunnel
	for {
		// Open a (local) connection to localEndpoint whose content will be forwarded so remoteEndpoint
		local, err := net.Dial("tcp", localEndpoint.String())
		if err != nil {
			log.Fatalln(fmt.Printf("Dial INTO local service error: %s", err))
		}

		client, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		handleClient(client, local)
	}

}