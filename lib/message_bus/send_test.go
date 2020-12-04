package message_bus

// import (
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// )

// type mockSender struct{
// 	success bool
// }

// func (r *mockSender) GetLabel() string {
// 	return "Test Label"
// }

// func (r *mockSender) Send(msg interface{}) interface{} {
// 	return r.success
// }

// func TestSenderLabel(t *testing.T) {
// 	s := &mockSender{}
// 	assert.Equal(t, s.GetLabel(), "Test Label")
// }

// func TestSenderFailSend(t *testing.T) {
// 	s := &mockSender{
// 		success: false,
// 	}

// 	sent, ok := s.Send("foo").(bool)
// 	assert.True(t, ok)
// 	assert.True(t, !sent)
// }

// func TestSenderSuccessfulSend(t *testing.T) {
// 	s := &mockSender{
// 		success: true,
// 	}

// 	sent, ok := s.Send("foo").(bool)
// 	assert.True(t, ok)
// 	assert.True(t, sent)
// }