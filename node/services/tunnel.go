package services

import (
    "os/exec"
    "fmt"
)

type Tunnel struct{
    node *node.Node
}

func NewTunnel(node *node.Node) *Tunnel {
    tunnel := &Tunnel{
        node: node,
    }

    return tunnel
}

func (this *Tunnel) GetLabel() string {
    return "tunnel"
}

func (this *Tunnel) IsRunning() boolean {
    return true
}

func (this *Tunnel) Run() {
    if this.node.HostNode != nil {
        fmt.Println("Starting tunnel.")
        fmt.Println("ssh -o ServerAliveInterval=300 -N -R"+this.node.HostNode.Api.HostPort+":localhost:22 root@"+this.node.HostNode.Addr)
        cmd := exec.Command("ssh", "-o ServerAliveInterval=300", "-N", "-R", this.node.HostNode.Api.HostPort+":localhost:22 root@"+this.node.HostNode.Addr)
        cmd.Start()
    }
}
