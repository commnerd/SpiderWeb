package node

import (
	"math/rand"

	_ "../message_bus"
	"../config"
	"../keys"
	"../api"
	"../id"
)

var privKeyMap map[id.Id]string

func (node *Node) RegisterChild(child Node) *api.RegisterResponse {
	privKey, pubKey := keys.Generate()
	adjustedId, err := id.Derive(child.Id, node.Mask)
	if err != nil {
		return craftForwardedNodeResponse(*node, child)
	}

	child.Id = adjustedId

	child.SshKeys = keys.Keys{
		Priv: privKey,
		Pub: pubKey,
	}

	node.registrants[child.Id] = child

	return craftRegisteredNodeResponse(*node, child)
}

func craftRegisteredNodeResponse(node, child Node) *api.RegisterResponse {
	privKey, pubKey := keys.Generate()
	privKeyMap[child.Id] = privKey
	return &api.RegisterResponse{
		Status: api.Success,
		Version: config.GetString("project_version"),
		AdjustedId: child.Id.String(),
		Mask: node.Mask + 1,
		Ip: node.Ip,
		Port: node.SshPort,
		PublicRsa: pubKey,
	}
}

func craftForwardedNodeResponse(node, child Node) *api.RegisterResponse {
	newParent := assignNewParent(node, child)

	return &api.RegisterResponse{
		Status: api.Forward,
		Version: config.GetString("project_version"),
		Ip: newParent.Ip,
		Port: newParent.SshPort,
	}
}

func assignNewParent(node, child Node) Node {
	var slice []id.Id = make([]id.Id, 0)

	for id, _ := range node.children {
		slice = append(slice, id)
	}

	id := slice[rand.Int() % len(slice)]
	return node.children[id].(Node)
}