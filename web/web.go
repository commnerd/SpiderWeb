// The web package provides all necessary endpoints for messaging between nodes
package main 

import (
    "github.com/gorilla/mux"
    "net/http"
    "log"
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
type server struct{
    router *mux.Router
}

// The web service held open for node communications
type Server interface{
    Start()
}

func NewServer() *server {
    return &server{ mux.NewRouter() }
}

// Start the web server
func (this *server) Start() {
    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", this.router))
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

func main() {
	s := NewServer()
	s.Start()
}
