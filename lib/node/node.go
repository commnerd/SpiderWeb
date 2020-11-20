package node

import (
	"../id"
)

type Node struct{
	Registrants map[id.Id]interface{}
	Children map[id.Id]interface{}
	ChildCount uint64
}

func RegisterChild(node Node) bool {
	return true
}