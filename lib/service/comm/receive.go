package comm

import "./comm/msg"

type Receiver interface{
	Execute(msg.Message)
}

func (r Receiver) Receive(m msg.Message) {
	r.Execute(m)
}