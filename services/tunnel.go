package services

import "os/exec"

const (
	TunnelPrivateKey string = "/var/run/spider_web/keys/id_rsa"
	TunnelPublicKey  string = "/var/run/spider_web/keys/id_rsa.pub"
)

type Tunnel ServiceStruct

func NewTunnel(node Node) *Tunnel {
	return &Tunnel{ node, "tunnel", ServiceStatusInit }
}

func (this *Tunnel) GetServiceLabel() string {
	return this.Label
}

func (this *Tunnel) GetStatus() ServiceStatus {
	return this.Status
}

func (this *Tunnel) Run() {
	go func() {
		this.Status = ServiceStatusRunning
		cmd := exec.Command("ssh", "-i", TunnelPrivateKey, this.Node.GetRegistrar().GetAddress())
		cmd.Run()
		this.Status = ServiceStatusDead
	}()
}


	
	