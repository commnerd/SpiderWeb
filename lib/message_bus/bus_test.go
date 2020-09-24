package message_bus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type sendReceive struct{
	sent bool
}

var sender Sender
var receiver Receiver
var senderReceiver *sendReceive

func (sr *sendReceive) GetLabel() string {
	return "Sender/Receiver"
}

func (sr *sendReceive) Send(msg interface{}) bool {
	return sr.sent
}

func (sr *sendReceive) Receive() interface{} {
	return "Message"
}

var bus *Bus

func init() {
	sender = &mockSender{}
	receiver = &mockReceiver{}
	senderReceiver = &sendReceive{}
	bus = &Bus{
		Receivers: make(map[string]Receiver),
		Senders: make(map[string]Sender),
	}
}

func TestInit(t *testing.T) {
	assert.IsType(t, &Bus{}, instance)
}

func TestNew(t *testing.T) {
	assert.IsType(t, &Bus{}, New())
}

func TestGet(t *testing.T) {
	assert.Equal(t, instance, Get())
}

func TestInstanceRegister(t *testing.T) {
	instance = New()

	Register(senderReceiver)

	assert.Equal(t, instance.Receivers["Sender/Receiver"], senderReceiver)
	assert.Equal(t, instance.Senders["Sender/Receiver"], senderReceiver)
}

func TestVariableRegister(t *testing.T) {
	sndRcv := New()

	sndRcv.Register(senderReceiver)

	assert.Equal(t, sndRcv.Receivers["Sender/Receiver"], senderReceiver)
	assert.Equal(t, sndRcv.Senders["Sender/Receiver"], senderReceiver)
}