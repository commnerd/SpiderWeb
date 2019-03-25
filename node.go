package main

import "./services"

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
	Services []services.Service
}

func NewNode() Node {
	return Node{
		Role: NodeRoleRoot,
		Address: "localhost",
	}
}

func (this *Node) Run() {
	services.Bootstrap(this)
	api := NewApi(this)
	api.Run()
}

func (this *Node) RegisterService(service services.Service) {
	this.Services = append(this.Services, service)
}

func (this *Node) GetRegistrar() services.Node {
	return this.Registrar
}

func (this *Node) GetAddress() string {
	return this.Address
}

func (this *Node) GetRole() string {
	switch this.Role {
	case NodeRoleRoot:
		return "root"
	case NodeRoleRegistrar:
		return "registrar"
	}
	return "worker"
}