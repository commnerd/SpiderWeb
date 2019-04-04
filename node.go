package main

import (
	"github.com/google/uuid"
	"encoding/json"
	"./services"
	"io/ioutil"
	"net/http"
	"reflect"
	"bytes"
	"log"
	"fmt"
)

const RootNodeAddress = "http://spiderweb.com/register"

type NodeRole int

const (
	NodeRoleInit NodeRole = iota
	NodeRoleRoot
	NodeRoleRegistrar
	NodeRoleWorker
)

type Node struct{
	Id string
	Role NodeRole
	Address string
	Registrar *Node
	ServiceChannel chan services.ServiceNotification
	Services []*services.Service
	Registry []*RegistryNode
	PubKey string
}

type JsonNode struct{
	Id string           `json:"id,omitempty"`
    Role string         `json:"role"`
    Address string      `json:"address"`
    Registrar *JsonNode `json:"registrar,omitempty"`
    PubKey string 		`json:"id_rsa_pub,omitempty"`
}

type RegistryNode struct{
	Id string       `json:"id,omitempty"`
    Address string  `json:"address"`
}

func (this *Node) Run() {
	services.Bootstrap(this)
	api := NewApi(this)
	go this.setRole()
	api.Run()
}

func (this *Node) RegisterService(service *services.Service) {
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
        PubKey: this.PubKey,
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
	this.Role = GetRoleFromLabel(jNode.Role)
	this.Address = jNode.Address
	this.PubKey = jNode.PubKey
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

	this.Role = node.Role
}

func (this *Node) monitorServices() {
	for {
		notification := <-this.ServiceChannel
		msg := ""
		switch notification.Event {
		case services.ServiceInitialized:
			msg = reflect.TypeOf(notification.Service).String() + " was started."
		case services.ServiceRunning:
			msg = reflect.TypeOf(notification.Service).String() + " is running."
		case services.ServiceDied:
			msg = reflect.TypeOf(notification.Service).String() + " has died."
		case services.ServiceKilled:
			msg = reflect.TypeOf(notification.Service).String() + " was terminated."
		}
		log.Println(msg)
	}
}

func NewNode() Node {
	pubKey, _ := GenerateKeys()
	return Node{
		Id: uuid.New().String(),
		Role: NodeRoleInit,
		Address: "localhost",
		Registrar: &Node{ Address: RootNodeAddress },
		ServiceChannel: make(chan services.ServiceNotification),
		Services: make([]*services.Service, 0),
		Registry: make([]*RegistryNode, 0),
		PubKey: string(pubKey),
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
