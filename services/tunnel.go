package services

import "os/exec"

const (
	TunnelPrivateKey string = "/var/run/spider_web/keys/id_rsa"
	TunnelPublicKey  string = "/var/run/spider_web/keys/id_rsa.pub"
)

type Tunnel ServiceStruct

func NewTunnel(node Node) *Tunnel {
	return &Tunnel{ node, "tunnel", 0}
}

func (this *Tunnel) GetNode() Node {
	return this.Node
}

func (this *Tunnel) GetLabel() string {
	return this.Label
}

func (this *Tunnel) GetIndex() int {
	return this.Index
}

func (this *Tunnel) SetIndex(index int) {
	this.Index = index
}

func (this *Tunnel) Run() {
	go func() {
		commChannel := this.GetNode().GetCommChannel()
		commChannel <- ServiceNotification{ this, ServiceInitialized }
		cmd := exec.Command("ssh", "-i", TunnelPrivateKey, this.Node.GetRegistrar().GetAddress())
		cmd.Run()
		commChannel <- ServiceNotification{ this, ServiceDied }
	}()
}
