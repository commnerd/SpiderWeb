package models

import (
    "testing"
)

func TestNewMessage(t *testing.T) {
    m := NewMessage(LoginRequest)
    if m.GetType() != LoginRequest {
        t.Errorf("MessageType - Expected %v, Got %v", LoginRequest, m.GetType())
    }
}
