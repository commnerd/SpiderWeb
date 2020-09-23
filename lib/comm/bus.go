package comm

import (
	"../node"
	"./msg"
	"./out"
	"./in"
)

type Bus interface{
	Send(msg.Message)
	Receive() msg.Message
}

type bus struct{
	Node node.Node
	Receiver in.Receiver
	Sender out.Sender
}

func NewBus(n node.Node) Bus {
	b := &bus{
		Node: n
	}
	b.Receiver = in.NewReceiver(b)
	b.Sender = out.NewSender(b)

	return b
}

func (b *bus) Send(m msg.Message) {
	b.Sender.Send(m)
}

func (b *bus) Receive() msg.Message{
	return <- b.Receiver.GetChannel()
}