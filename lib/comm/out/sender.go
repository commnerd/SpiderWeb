package out

import (
	"github.com/google/uuid"
	"../msg"
)

type Bus interface{
	Send(msg.Message)
}

type Sender interface{
	Send(msg.Message)
}

type sender struct{
	Id uuid.UUID
	Channel chan msg.Message
	Bus Bus
}

func NewSender(bus Bus) Sender {
	return &sender{
		Id: uuid.New(),
		Channel: make(chan msg.Message),
		Bus: bus,
	}
}

func (s *sender) Send(m msg.Message) {
	fun := func() {
		for {
			m := <- s.Channel
			s.Send(m)
		}
	}
	go fun()
}