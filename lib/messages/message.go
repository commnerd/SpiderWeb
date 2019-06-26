// Package messages : The messages package provides message constructs to communicate between nodes
package messages

// Message structure
type msg struct {
    Type MessageType
    Payload interface{}
}

// MessageType declaration
type MessageType int

// Type constants
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

// Message : The groundwork for all node communications
type Message interface{
	getType() MessageType
	getPayload() interface{}
}
