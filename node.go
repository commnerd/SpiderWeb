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
	ServiceRegistry *ServiceRegistry
	NodeRegistry *NodeRegistry
}

type JsonNode struct{
	Id string           `json:"id,omitempty"`
	IdNetworkMask int   `json:"id_network_mask"`
    Role string         `json:"role"`
    Address string      `json:"address"`
    Registrar *JsonNode `json:"registrar,omitempty"`
}

func (this *Node) Run() {
	this.ServiceRegistry = NewServiceRegistry(this)
	services.Bootstrap(this)
	api := NewApi(this)
	go this.setRole()
	api.Run()
}

func (this *Node) GetNodeRegistry() *NodeRegistry {
	return this.NodeRegistry
}

func (this *Node) GetServiceRegistry() services.Registry {
	return this.ServiceRegistry
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
	resp, err := http.Post(this.Registrar.Address+"/register", "application/json", bytes.NewBuffer(this.MarshalJSON()))
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

	if node.Registrar != nil && this.Registrar.Address != node.Registrar.Address {
		this.changeRegistrar(node.Registrar)
		return
	}

	this.Id = node.Id
	this.IdNetworkMask = node.IdNetworkMask
	this.Role = node.Role
	this.Address = node.Address
}

func (this *Node) changeRegistrar(newReg *Node) {
	oldId := this.Id
	oldRegAddress := this.Registrar.Address
	this.Registrar = newReg
	this.setRole()
	this.dropOldRegistrar(oldRegAddress, oldId)
}

func (this *Node) dropOldRegistrar(oldRegAddress string, id string) {
	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", "http://"+oldRegAddress+"/registry/"+id, nil)

	client.Do(req)
}

func NewNode() Node {
	node := Node{
		Id: uuid.New().String(),
		IdNetworkMask: 0,
		Role: NodeRoleInit,
		Address: "localhost",
		Registrar: &Node{ Address: RootNodeAddress },
	}
	node.ServiceRegistry = NewServiceRegistry(&node)
	node.NodeRegistry = NewNodeRegistry(&node)
	return node
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
