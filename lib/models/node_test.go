package models

import (
    "testing"
    "net/url"
)

func NewNodeTest(t *testing.T) {
    loc, err := url.Parse("test.com:3030")
    if err != nil {
        panic(err)
    }
    n := NewNode(loc)

    if n.Url.RequestURI() != "http://test.com:3030" {
        t.Errorf("Expecting \"http://test.com:3030\", Got \"%s\"", n.Url.RequestURI())
    }
}
