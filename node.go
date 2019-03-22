package main

type NodeRole int

const (
	NodeRoleRoot NodeRole = iota
	NodeRoleRegistrar
	NodeRoleWorker
)

type Node struct{
	Role NodeRole
	Address string
	Registrar *Node
	Services []interface{}
}

func NewNode() Node {
	return Node{
		Role: NodeRoleRoot,
		Address: "localhost",
	}
}

func (this *Node) Run() {
	api := NewApi(this)
	api.Run()
}

func (this *Node) RegisterService(service Service)