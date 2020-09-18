package in

import (
	"github.com/google/uuid"
	"../msg"
	"fmt"
)

type Bus interface{
	Receive() msg.Message
}

type Receiver interface{
	Listen()
	GetChannel() chan msg.Message
}

type receiver struct{
	Id uuid.UUID
	Channel chan msg.Message
	Bus Bus
}

func NewReceiver(bus Bus) Receiver {
	return &receiver{
		Id: uuid.New(),
		Channel: make(chan msg.Message),
		Bus: bus,
	}
}

func (r *receiver) Listen() {
	fun := func() {
		for {
			m := <- r.Channel
			fmt.Println(m)
		}
	}
	go fun()
}

func (r *receiver) GetChannel() chan msg.Message {
	return r.Channel
}
