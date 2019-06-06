// The web package provides all necessary endpoints for messaging between nodes
package web

import (
    "net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

// The structure for maintaining route state
type route struct {
    path string
    handler Handler
}

// The route interface for the rest of the application
type Route interface{
    GetPath() string
    GetFunc() func(http.ResponseWriter, *http.Request)
}

// The structure for maintaining server state
type server struct{}

// The web service held open for node communications
type Server interface{
    Start()
}

func NewServer() *server {
    return &server{}
}

// Start the web server
func (s *server) Start() {
}

// The structure definition of a request
type request http.Request

// The request structure passed to a server
type Request interface{
    GetVerb() string
    GetReferer() string
    GetBody() string
}

// The structure definition for a response
type response struct{}

// The response structure returned from a server
type Response interface{
    GetFormat() string
    GetBody() string
}
