// Package login : Login message logic
package login

import (
    "../../models"
    ".."
)

type message struct{}

type payload struct{
    Node models.Node
}

// New : Craft a new login message
func New() messages.Message {
    return &message{}
}

// GetType : Get the message type
func (msg *message) GetType() messages.MessageType {
    return messages.LoginRequest
}

// GetPayload : Get the message payload
func (msg *message) GetPayload() interface{} {
    return payload{
        models.Node
    }
}
