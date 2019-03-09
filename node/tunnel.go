package node

import (
    "os/exec"
)

type Tunnel struct{
    node *Node
}

func NewTunnel(node *Node) *Tunnel {
    return &Tunnel{
        node: node,
    }
}

func (this *Tunnel) Run() {
    cmd := exec.Command("ssh", "-o ServerAliveInterval=300", "-N", "-R", this.node.HostNode.Api.HostPort+":localhost:22 root@"+this.node.HostNode.Ip)
    cmd.Start()
}
