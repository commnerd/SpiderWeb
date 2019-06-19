package models

import (
    "net/http"
    "testing"
)

type testMessage struct{}

func (this *testMessage) BuildRequest() *http.Request {
    req, _ := http.NewRequest("GET", "http://example.com", nil)
    return req
}

func TestMessageBuildRequest(t *testing.T) {
    msg := &testMessage{}
    if _, ok := interface{}(msg).(Message); !ok {
        t.Errorf("Not a message.")
    }
}
