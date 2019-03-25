package services

import "os/exec"

const (
	TunnelPrivateKey string = "/var/run/spider_web/keys/id_rsa"
	TunnelPublicKey  string = "/var/run/spider_web/keys/id_rsa.pub"
)

type Tunnel Service

func NewTunnel(node Node) Tunnel {
	return Tunnel{ node, "tunnel", 0}
}

func (this *Tunnel) Run() {
	go func() {
		service := Service(*this)
		commChannel := this.Node.GetCommChannel()
		commChannel <- ServiceNotification{ service, ServiceInitialized }
		cmd := exec.Command("ssh", "-i", TunnelPrivateKey, this.Node.GetRegistrar().GetAddress())
		cmd.Run()
		commChannel <- ServiceNotification{ service, ServiceDied }
	}()
}