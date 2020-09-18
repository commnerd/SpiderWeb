package comm

import (
	"./msg"
	"./out"
	"./in"
)

type bus struct{
	Receiver in.Receiver
	Sender out.Sender
}

type Bus interface{
	Receive() msg.Message
}

func NewBus() Bus {
	return &bus{
		Receiver: in.NewReceiver(),
		Sender: out.NewSender(),
	}
}

func (b *bus) Receive() msg.Message {
	return msg.Message{}
}