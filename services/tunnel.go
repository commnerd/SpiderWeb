package services

import "os/exec"

const (
	TunnelPrivateKey string = "/var/run/spider_web/keys/id_rsa"
	TunnelPublicKey  string = "/var/run/spider_web/keys/id_rsa.pub"
)

type Tunnel struct{
	Node Node
}

func NewTunnel(node Node) *Tunnel {
	return &Tunnel{ node }
}

func (this *Tunnel) GetLabel() string {
	return "Tunnel"
}

func (this *Tunnel) GetNode() Node {
	return this.Node
}

func (this *Tunnel) Run() {
	go func() {
		service := Service(this)
		serviceChannel := this.GetNode().GetServiceRegistry().GetEventChannel()
		serviceChannel <- ServiceNotification{ service, ServiceInitialized }
		cmd := exec.Command("ssh", "-i", TunnelPrivateKey, this.GetNode().GetRegistrar().GetAddress())
		cmd.Run()
		serviceChannel <- ServiceNotification{ service, ServiceDied }
	}()
}
