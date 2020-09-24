package message_bus

type Receiver interface{
	GetLabel() string
	Receive() interface{}
}