package node

import (
	"github.com/google/uuid"
	"../util"
	"../ids"
)

type Node struct{
	Id uuid.UUID
	Mask ids.Mask
	Parent *Node
}

func New(parent *Node) Node {
	id := uuid.New()
	mask := ids.Mask(-1)

	if(parent != nil) {
		mask = getNextMask(parent.Mask)
		id = ids.Create(parent.Id, mask)
	}

	return Node {
		Id: id,
		Mask: mask,
		Parent: parent,
	}
}

func getNextMask(mask ids.Mask) ids.Mask {
	mask++
	exists, _ := util.InArray(mask, ids.BadMasks)
	if(exists) {
		return mask + 1
	}
	return mask
}