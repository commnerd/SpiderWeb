package models

// Message structure
type msg struct {
    Type MessageType
    Payload interface{}
}

// MessageType declaration
type MessageType int

// Message types
const (
	LoginRequest MessageType = iota
	LoginResponse
	TunnelRequest
	TunnelResponse
	RegisterRequest
	RegisterResponse
	TunnelBreakNotification
	PublicCheck
	PublicResponse
	PublicTunnelResponse
	KeyUpdateNotifications
)
// The groundwork for all node communications
type Message interface {
	GetType() MessageType
	GetPayload() interface{}
    Send() error
}

// Craft a new message
func NewMessage(mt MessageType) Message {
    return &msg{
        Type: mt,
    }
}

// Get the message type
func (this msg) GetType() MessageType {
    return this.Type
}

// Get the message payload
func (this msg) GetPayload() interface{} {
    return "YAY!"
}

// Send the message
func (this msg) Send() error {
    return nil
}
