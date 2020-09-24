package message_bus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockSender struct{
	fail bool
}

func (r *mockSender) GetLabel() string {
	return "Test Label"
}

func (r *mockSender) Send(msg interface{}) bool {
	return r.fail
}

func TestSenderLabel(t *testing.T) {
	s := &mockSender{}
	assert.Equal(t, s.GetLabel(), "Test Label")
}

func TestSenderFailSend(t *testing.T) {
	s := &mockSender{
		fail: false,
	}

	assert.True(t, !s.Send("foo"))
}

func TestSenderSuccessfulSend(t *testing.T) {
	s := &mockSender{
		fail: true,
	}

	assert.True(t, s.Send("bar"))
}