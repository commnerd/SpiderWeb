package message_bus

import (
	"reflect"
	"errors"
	"fmt"
)

var die = false

const (
	channel_exists_error = "Channel \"%s\" already exists."
	channel_not_exists_error = "Channel \"%s\" does not exist."
)

type Channel chan interface{}

func (c Channel) Send(msg interface{}) {
	go func() { c <- msg }()
}

func (c Channel) Receive(args ...interface{}) interface{} {

	if len(args) < 1 {
		return <- c
	}

	if len(args) > 1 {
		panic("OH SHIT")
	}

	go func(fn interface{}, c Channel) {
		callback := reflect.ValueOf(fn)
		if callback.Kind() == reflect.Func {
			for {
				inInterface := <- c
				if str, ok := inInterface.(string); ok && str == "die" {
					return
				}
				in := []reflect.Value{reflect.ValueOf(inInterface)}

				callback.Call(in)
			}
		}
	}(args[0], c)

	return nil
}

func MakeChannel(channel string) error {
	if instance.channels[channel] != nil {
		return errors.New(fmt.Sprintf(channel_exists_error, channel))
	}
	instance.channels[channel] = make(Channel)
	return nil
}

func GetChannel(channel string) (Channel, error) {
	if instance.channels[channel] == nil {
		return nil, errors.New(fmt.Sprintf(channel_not_exists_error, channel))
	}
	return instance.channels[channel], nil
}