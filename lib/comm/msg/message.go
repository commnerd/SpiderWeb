package msg

import (
	"github.com/google/uuid"
	"../../node"
)

type Node interface{
	GetId() uuid.UUID
	GetMask() int
}

type Message interface{
	Id uuid.UUID
	Receipient node.Node
	Body interface{}
}

type Messageable interface{
	ToMessage() msg.Message
}