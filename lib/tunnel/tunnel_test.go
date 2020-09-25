package tunnel

import (
	"github.com/stretchr/testify/assert"
	sshSvr "github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"../config"
	"../keys"
	"testing"
	"strconv"
	"../port"
	"time"
	"net"
	"fmt"
	"io"
	"os"
)

type testConn struct{
	toSend string
	buffer []byte
	addr net.Addr
	remote net.Addr
}

var mockEndpoint *Endpoint
var mockClientConn net.Conn
var mockRemoteConn net.Conn

func TestMain(m *testing.M) {
	config.Set("id_rsa_pub_path", "/tmp/id_rsa.pub")
	config.Set("id_rsa_path", "/tmp/id_rsa")

	keys.Generate()

	ioutil.ReadFile(config.GetString("id_rsa_pub_path"))
	ioutil.ReadFile(config.GetString("id_rsa_path"))

	mockEndpoint = &Endpoint{
		Host: "127.0.0.1",
		Port: 12345,
	}
	m.Run()

	os.Remove(config.GetString("id_rsa_path"))
	os.Remove(config.GetString("id_rsa_pub_path"))
}

func TestTunnelString(t *testing.T) {
	assert.Equal(t, "127.0.0.1:12345", mockEndpoint.String())
}

func TestMockReader(t *testing.T) {
	content := make([]byte, 5)
	reader := &testConn{toSend: "tested"}
	reader.Read(content)
	assert.Equal(t, "teste", string(content))
}

func TestMockWriter(t *testing.T) {
	writer := &testConn{buffer: make([]byte, 0)}
	writer.Write([]byte("tested"))
	assert.Equal(t, "tested", string(writer.buffer))
}

func TestHandleClient(t *testing.T) {
	client := &testConn{
		toSend: "Foo",
		buffer: make([]byte, 0),
		addr: &mockAddr{
			addr: "client:4321",
		},
		remote: &mockAddr{
			addr: "remote:1234",
		},
	}
	remote := &testConn{
		toSend: "Bar",
		buffer: make([]byte, 0),
		addr: &mockAddr{
			addr: "remote:1234",
		},
		remote: &mockAddr{
			addr: "client:4321",
		},
	}

	go handleClient(client, remote)
	time.Sleep(time.Millisecond)

	assert.Equal(t, "Foo", string(remote.buffer))
	assert.Equal(t, "Bar", string(client.buffer))
}

func TestPublicKeyFile(t *testing.T) {

	ret := publicKeyFile(config.GetString("id_rsa_pub_path"))

	_, ok := ret.(ssh.AuthMethod)

	assert.True(t, ok)
}

func TestOpen(t *testing.T) {

	sshSvr.Handle(func(s sshSvr.Session) {
		authorizedKey := ssh.MarshalAuthorizedKey(s.PublicKey())
		io.WriteString(s, fmt.Sprintf("public key used by %s:\n", s.User()))
		s.Write(authorizedKey)
	})

	publicKeyOption := sshSvr.PublicKeyAuth(func(ctx sshSvr.Context, key sshSvr.PublicKey) bool {
		return true // allow all keys, or use ssh.KeysEqual() to compare against known keys
	})

	go func() {
		err := sshSvr.ListenAndServe("127.0.0.1:" + strconv.Itoa(int(port.MAX)), nil, publicKeyOption)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	time.Sleep(time.Millisecond)

	endpoint := Endpoint{
		Host: "127.0.0.1",
		Port: port.MAX,
	}
	go Open(port.MIN, endpoint)

	assert.True(t, !port.Available(port.MAX))
}

func (c *testConn) Read(b []byte) (n int, err error) {
	n = len(c.toSend)
	for i := 0; i < cap(b) && i < n; i++ {
		b[i] = c.toSend[i]
	}
	err = io.EOF
	return
}
func (c *testConn) Write(b []byte) (n int, err error) {
	n = 0
	for i := 0; i < len(b); i++ {
		c.buffer = append(c.buffer, b[i])
		n++
	}
	err = nil
	return
}
func (c *testConn) Close() error {
	return nil
}
func (c *testConn) LocalAddr() net.Addr {
	return c.addr
}
func (c *testConn) RemoteAddr() net.Addr {
	return c.remote
}
func (c *testConn) SetDeadline(t time.Time) error {
	return nil
}
func (c *testConn) SetReadDeadline(t time.Time) error {
	return nil
}
func (c *testConn) SetWriteDeadline(t time.Time) error {
	return nil
}

type mockAddr struct{
	addr string
}
func (a *mockAddr) Network() string {
	return "tcp"
}
func (a *mockAddr) String() string {
	return a.addr
}