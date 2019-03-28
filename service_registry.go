package main

import 	(
	"./services"
	"log"
)

type ServiceRegistry struct {
    Node *Node
    Services []services.Service
    EventChannel chan services.ServiceNotification
}

func NewServiceRegistry(node *Node) *ServiceRegistry {
    return &ServiceRegistry{
        Node: node,
        Services: make([]services.Service, 0),
        EventChannel: make(chan services.ServiceNotification),
    }
}

func (this *ServiceRegistry) Add(service services.Service) {
	this.Services = append(this.Services, service)
}

func (this *ServiceRegistry) Remove(s services.Service) {
	for index, service := range this.Services {
		if service == s {
			this.Services = append(this.Services[:index], this.Services[index+1:]...)
			return
		}
	}
}

func (this *ServiceRegistry) GetEventChannel() chan services.ServiceNotification {
    return this.EventChannel
}

func (this *ServiceRegistry) Monitor() {
	for {
		notification := <-this.EventChannel
		msg := ""
		switch notification.Event {
		case services.ServiceInitialized:
			msg = notification.Service.GetLabel() + " was started."
		case services.ServiceRunning:
			msg = notification.Service.GetLabel() + " is running."
		case services.ServiceDied:
			msg = notification.Service.GetLabel() + " has died."
		case services.ServiceKilled:
			msg = notification.Service.GetLabel() + " was terminated."
		}
		log.Println(msg)
	}
}
