package main

import (
	"./services"
	"log"
)

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
	CommChannel chan services.ServiceNotification
	Services []*services.Service
}

func NewNode() Node {
	return Node{
		Role: NodeRoleRoot,
		Address: "localhost",
		CommChannel: make(chan services.ServiceNotification),
		Services: make([]*services.Service, 0),
	}
}

func (this *Node) Run() {
	services.Bootstrap(this)
	api := NewApi(this)
	api.Run()
}

func (this *Node) RegisterService(service *services.Service) {
	service.Index = len(this.Services)
	this.Services = append(this.Services, service)
}

func (this *Node) GetCommChannel() chan services.ServiceNotification {
	return this.CommChannel
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

func (this *Node) monitorServices() {
	for {
		notification := <-this.CommChannel
		msg := ""
		switch notification.Event {
		case services.ServiceInitialized:
			msg = notification.Service.Label + " was started."
		case services.ServiceRunning:
			msg = notification.Service.Label + " is running."
		case services.ServiceDied:
			msg = notification.Service.Label + " has died."
		case services.ServiceKilled:
			msg = notification.Service.Label + " was terminated."
		}
		log.Println(msg)
	}
}