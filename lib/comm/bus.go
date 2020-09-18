package comm

import (
	"./msg"
	"./out"
	"./in"
)

type Bus interface{
	Send(msg.Message)
	Receive() msg.Message
}

type bus struct{
	Receiver in.Receiver
	Sender out.Sender
}

func NewBus() Bus {
	b := &bus{}
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