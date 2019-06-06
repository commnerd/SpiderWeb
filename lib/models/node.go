package models

import (
    "net/url"
)

// Node structure
type node struct{
    Url *url.URL
}

// Node interface
type Node interface{
    GetUrl() *url.URL
}

// Craft and return a new Node
func NewNode(loc *url.URL) *node {
    return &node{
        Url: loc,
    }
}

// Get the Node's url
func (this *node) GetUrl() *url.URL {
    return this.Url
}
