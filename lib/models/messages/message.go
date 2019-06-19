package messages

type MessageType int

// Message types
const (
    InitRequest MessageType = iota
    InitResponse
	LoginRequest
	LoginResponse
	TunnelRequest
	TunnelResponse
	RegisterRequest
	RegisterResponse
	TunnelBreakNotification
	PublicCheck
	PublicResponse
	PublicTunnelResponse
	KeyUpdateNotification
)
