package service

import (
	"../comm"
	"fmt"
	"os"
)

type Manager interface{}

type manager struct{
	Node interface{}
	Bus comm.Bus
	Services []Service
}

func NewManager(node interface{}) Manager {
	manager := &manager{
		Node: node,
		Bus: comm.NewBus(),
		Services: make([]Service, 0),
	}
	manager.Run()
	return manager
}

func (manager *manager) Run() {
	for {
		msg := manager.Bus.Receive()
		fmt.Println(msg)
		os.Exit(0)
	}
}