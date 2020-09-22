package in

import (
	"github.com/google/uuid"
	"../msg"
)

type mockReceiver struct{
	Id uuid.UUID
	Channel chan msg.Message
	Bus Bus
}

func ReceiverMock(bus Bus) Receiver {
	return &mockReceiver{
		Id: uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da"),
		Channel: make(chan msg.Message),
		Bus: bus,
	}
}

func (r *mockReceiver) GetChannel() chan msg.Message {
	return r.Channel
}

func (r *mockReceiver) Listen() {

}