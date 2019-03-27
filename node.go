package main

import (
	"github.com/google/uuid"
	"encoding/json"
	"./services"
	"io/ioutil"
	"net/http"
	"bytes"
	"log"
	"fmt"
)

const RootNodeAddress = "http://spiderweb.com"

type NodeRole int

const (
	NodeRoleInit NodeRole = iota
	NodeRoleRoot
	NodeRoleRegistrar
	NodeRoleWorker
)

type Node struct{
	Id string
	IdNetworkMask int
	Role NodeRole
	Address string
	Registrar *Node
	ServiceChannel chan services.ServiceNotification
	Services []*services.Service
	Registry []*RegistryNode
}

type JsonNode struct{
	Id string           `json:"id,omitempty"`
	IdNetworkMask int   `json:"id_network_mask"`
    Role string         `json:"role"`
    Address string      `json:"address"`
    Registrar *JsonNode `json:"registrar,omitempty"`
}

type RegistryNode struct{
	Id string         `json:"id,omitempty"`
	IdNetworkMask int `json:"id_network_mask"`
    Address string    `json:"address"`
}

func (this *Node) Run() {
	services.Bootstrap(this)
	api := NewApi(this)
	go this.setRole()
	api.Run()
}

func (this *Node) RegisterService(service *services.Service) {
	service.Index = len(this.Services)
	this.Services = append(this.Services, service)
}

func (this *Node) GetServiceChannel() chan services.ServiceNotification {
	return this.ServiceChannel
}

func (this *Node) GetRegistrar() services.Node {
	return this.Registrar
}

func (this *Node) GetAddress() string {
	return this.Address
}

func (this *Node) GetRole() NodeRole {
	return this.Role
}

func (this *Node) GetRoleLabel() string {
	switch this.Role {
	case NodeRoleRoot:
		return "root"
	case NodeRoleRegistrar:
		return "registrar"
	}
	return "worker"
}

func (this *Node) MarshalJSON() []byte {
	jNode := JsonNode{
        Id: this.Id,
        Role: this.GetRoleLabel(),
        Address:  this.Address,
    }

	data, err := json.Marshal(jNode)
    if(err != nil) {
		fmt.Println(data)
    	log.Fatalln(err)
    }

    return data
}

func (this *Node) UnmarshalJSON(contents []byte) {
	var jNode JsonNode
	err := json.Unmarshal(contents, &jNode)
	if(err != nil) {
		log.Fatalln(err)
	}

	this.Id = jNode.Id
	this.IdNetworkMask = jNode.IdNetworkMask
	this.Role = GetRoleFromLabel(jNode.Role)
	this.Address = jNode.Address
}

func (this *Node) setRole() {
	resp, err := http.Post(RootNodeAddress, "application/json", bytes.NewBuffer(this.MarshalJSON()))
	if(err != nil) {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if(err != nil) {
		log.Fatalln(err)
	}

	var node Node
	node.UnmarshalJSON(body)

	this.Id = node.Id
	this.IdNetworkMask = node.IdNetworkMask
	this.Role = node.Role
	this.Address = node.Address
	if this.Registrar.Id != node.Registrar.Id {
		this.changeRegistrar(node.Registrar)
	}
}

func (this *Node) monitorServices() {
	for {
		notification := <-this.ServiceChannel
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

func (this *Node) changeRegistrar(newReg *Node) {
	currentId := this.Id
	if this.registerWithNewRegistrar(newReg) {
		this.Registrar = newReg
		this.dropOldRegistrar(currentId)
	}
}

func (this *Node) contactNewRegistrar(newReg *Node) bool {
	res, err := http.Post(newReg.Address+"/register", "application/json", this.MarshalJSON())
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if(err != nil) {
		return false
	}

	var node Node
	node.UnmarshalJSON(body)

	this.Id = node.Id
	this.IdNetworkMask = node.IdNetworkMask
	this.Role = node.Role
	this.Address = node.Address
	if this.Registrar.Id != node.Registrar.Id {
		this.changeRegistrar(node.Registrar)
	}

}

func (this *Node) dropOldRegistrar(id string) {
	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", "http://"+this.Registrar.Address+"/registry/"+id)

	client.Do(req)
}

func NewNode() Node {
	return Node{
		Id: uuid.New().String(),
		IdNetworkMask: 0,
		Role: NodeRoleInit,
		Address: "localhost",
		Registrar: &Node{ Address: RootNodeAddress },
		ServiceChannel: make(chan services.ServiceNotification),
		Services: make([]*services.Service, 0),
		Registry: make([]*RegistryNode, 0),
	}
}

func GetRoleFromLabel(label string) NodeRole {
	switch label {
	case "root":
		return NodeRoleRoot
	case "registrar":
		return NodeRoleRegistrar
	}
	return NodeRoleWorker
}
