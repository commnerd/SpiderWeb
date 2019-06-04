package models

import (
    "./messages"
    "testing"
)

func TestNewMessage(t testing.T) {
    m := NewMessage(messages.LoginRequest)
    if messages.GetType() != messages.LoginRequest {
        t.Error('OOfta')
    }
}
