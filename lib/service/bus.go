package service

import (
	"./comm/in"
)

type Bus struct{
	Listener in.Server
}

type executor interface {
	Execute(Service)
}

func New(node executor) *Bus {
	bus := &Bus{
		Listener: in.New(),
	}
	return bus
}

func (bus *Bus) Triage(msg interface{}) {
}

func (bus *Bus) Send(msg interface{}) {
}