package in

import "../msg"

type receiver struct{
	Channel chan msg.Message
}

type Receiver interface{
	Receive() msg.Message
}

func NewReceiver() Receiver {
	return &receiver{}
}

func (r *receiver) Receive() msg.Message {
	return <- r.Channel
}