package in

import (
	"net/http"
)

type Handler interface{
	Path() string
	Handler() func(http.ResponseWriter, *http.Request)
}

type Server interface{
	SetHandler(handler)
	Serve()
}

type server struct{
	Handlers
}

func NewServer() Server {
	return &server{}
}

func triage(w http.ResponseWriter, req *http.Request) {
}

func (svr *server) Serve() {
	http.HandleFunc("/", triage)
	http.ListenAndServe(":8090", nil)
}