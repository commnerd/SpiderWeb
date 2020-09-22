package in

import (
	"net/http"
)

type Server interface{
	Serve()
}

type server struct{
	MsgHandler Receiver
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