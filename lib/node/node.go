package node

import (
	"github.com/google/uuid"
	"../ids"
)

type Node struct{
	Id uuid.UUID
	Mask int
	Parent *Node
}

func New(parent *Node) Node {
	id := uuid.New()
	mask := -1

	if(parent != nil) {
		mask = parent.Mask + 1
		id = ids.Create(parent.Id, mask)
	}

	return Node {
		Id: id,
		Mask: mask,
		Parent: parent,
	}
}