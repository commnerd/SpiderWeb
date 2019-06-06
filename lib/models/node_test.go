package models

import (
    "testing"
    "net/url"
)

func TestNewNode(t *testing.T) {
    loc, err := url.Parse("test.com:3030")
    if err != nil {
        panic(err)
    }
    n := NewNode(loc)

    if n.Url.String() != "test.com:3030" {
        t.Errorf("Expecting \"test.com:3030\", Got \"%s\"", n.Url.String())
    }
}
