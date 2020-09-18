package node

import (
	"github.com/google/uuid"
	"../service"
	"../util"
	"../ids"
)

type Node struct{
	Id uuid.UUID
	Mask ids.Mask
	Parent *Node
	Children map[string]*Node
	ServiceManager service.Manager
}

func New(parent *Node) *Node {
	id := uuid.New()
	mask := ids.Mask(-1)

	if(parent != nil) {
		mask = getNextMask(parent.Mask)
		id = ids.Create(parent.Id, mask)
	}

	node := &Node {
		Id: id,
		Mask: mask,
		Parent: parent,
		Children: make(map[string]*Node),
	}
	node.ServiceManager = service.NewManager(node)

	return node
}

func (parent *Node) AddChild(node *Node) {
	for i := 0; i <= int(parent.Mask); i++ {
		node.Id[i] = parent.Id[i]
	}
	parent.Children[node.Id.String()] = node
}

func (node *Node) GetId() uuid.UUID {
	return node.Id
}

func (node *Node) GetMask() ids.Mask {
	return node.Mask
}

func getNextMask(mask ids.Mask) ids.Mask {
	mask++
	exists, _ := util.InArray(mask, ids.BadMasks)
	if(exists) {
		return mask + 1
	}
	return mask
}