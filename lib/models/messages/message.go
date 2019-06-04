// The messages package provides message constructs to communicate between nodes
package messages

type MessageType int

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
}
