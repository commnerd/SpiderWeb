package models

// Node structure
type node struct{
    Config Config
    Services []Service
}

// Node interface
type Node interface{
    StartService(Service)
}

// Craft and return a new Node
func NewNode() *node {
    return &node{
        Services: make([]Service, 0),
    }
}

// Start a service on the node
func (this *node) StartService(s Service) {
    this.Services = append(this.Services, s)
}
