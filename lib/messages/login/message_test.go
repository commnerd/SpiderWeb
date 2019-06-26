package login

import (
    "testing"
    "fmt"
)

func TestNewLogin(t testing.T) {
    m := New()
    if _, ok := m.(LoginRequest); !ok {
        t.Errorf("New message not of type LoginRequest")
    }
}
