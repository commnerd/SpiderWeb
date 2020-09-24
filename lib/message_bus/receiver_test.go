package message_bus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockReceiver struct{
	pipe chan string
}

func (r *mockReceiver) GetLabel() string {
	return "Test Label"
}

func (r *mockReceiver) Receive() interface{} {
	return <- r.pipe
}

func TestReceiverLabel(t *testing.T) {
	r := &mockReceiver{}
	assert.Equal(t, r.GetLabel(), "Test Label")
}

func TestReceiverChannel(t *testing.T) {
	r := &mockReceiver{
		pipe: make(chan string),
	}

	go func(rcv Receiver) {
		in := rcv.Receive()
		assert.Equal(t, in, "Foo")
	}(r)

	r.pipe <- "Foo"
}