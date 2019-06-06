package models

import (
    "testing"
)

func NewMessageTest(t testing.T) {
    m := NewMessage(LoginRequest)
    if m.GetType() != LoginRequest {
        t.Errorf("MessageType - Expected %v, Got %v", LoginRequest, m.GetType())
    }
}
