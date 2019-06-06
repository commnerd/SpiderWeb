package web

import (
    "testing"
)

// Ensure we have a server to run
func TestNewServer(t *testing.T) {
    var s Server = NewServer()
    if _, ok := s.(Server); !ok {
        t.Errorf("Expected true, Got %v", ok)
    }
}

// Test server start
func TestServerStart(t *testing.T) {
    var s Server = NewServer()

    s.Start()
}
