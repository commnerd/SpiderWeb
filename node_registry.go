package main

type NodeRegistry struct {
    Node *Node
    Nodes map[string]*RegistryNode
}

type RegistryNode struct{
	Id string         `json:"id,omitempty"`
	IdNetworkMask int `json:"id_network_mask"`
    Address string    `json:"address"`
}

func NewNodeRegistry(node *Node) *NodeRegistry {
    return &NodeRegistry{
        Node: node,
        Nodes: make(map[string]*RegistryNode),
    }
}

func (this *NodeRegistry) All() map[string]*RegistryNode {
    return this.Nodes
}

func (this *NodeRegistry) Add(node *Node) {
    this.Nodes[node.Id] = &RegistryNode{
        Id: node.Id,
        IdNetworkMask: node.IdNetworkMask,
        Address: node.Address,
    }
}

func (this *NodeRegistry) Remove(id string) {
    var nodes map[string]*RegistryNode = make(map[string]*RegistryNode)

    for nodeId, node := range this.Nodes {
        if id != nodeId {
            nodes[nodeId] = node
        }
    }

    this.Nodes = nodes
}
