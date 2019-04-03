package services

import "os/exec"

const (
	TunnelPrivateKey string = "/root/.ssh/id_rsa"
	TunnelPublicKey  string = "/root/.ssh/id_rsa.pub"
)

type Tunnel Service

func NewTunnel(node Node) Tunnel {
	return Tunnel{ node, "tunnel", 0}
}

func (this *Tunnel) Run() {
	go func() {
		service := Service(*this)
		serviceChannel := this.Node.GetServiceChannel()
		serviceChannel <- ServiceNotification{ service, ServiceInitialized }
		cmd := exec.Command("ssh", "-i", TunnelPrivateKey, "-o", "ServerAliveInterval=300", "-NR", "8080:localhost:80", this.Node.GetRegistrar().GetAddress())
		cmd.Run()
		serviceChannel <- ServiceNotification{ service, ServiceDied }
	}()
}
