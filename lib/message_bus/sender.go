package message_bus

type Sender interface{
	GetLabel() string
	Send(msg interface{}) bool
}