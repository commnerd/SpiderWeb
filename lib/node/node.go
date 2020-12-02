package node

import (
	"../config"
	"../keys"
	"../id"
)

type Node struct{
	Id id.Id
	Mask id.Mask
	Ip string
	SshPort int
	SshKeys keys.Keys
	registrants map[id.Id]interface{}
	children map[id.Id]interface{}
}

func (node *Node) GetId() id.Id {
	return node.Id
}

func init() {
	if config.IsSet("NODE") {
		panic("You already have one node running in this process.")
	}
	config.Set("NODE", &Node{
		Id: id.New(),
		Mask: id.Mask(0),
		Ip: "",
		SshPort: 22,
		registrants: make(map[id.Id]interface{}, 0),
		children: make(map[id.Id]interface{}, 0),
	})
}