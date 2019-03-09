package node

import (
    "os/exec"
)

type Tunnel struct{
    Service
    node *Node
}

func NewTunnel(node *Node) *Tunnel {
    tunnel := &Tunnel{
        node: node,
    }

    return tunnel
}

func (this *Tunnel) Run() {
    if this.node.HostNode != nil {
        cmd := exec.Command("ssh", "-o ServerAliveInterval=300", "-N", "-R", this.node.HostNode.Api.HostPort+":localhost:22 root@"+this.node.HostNode.Ip)
        cmd.Start()
    }
}
