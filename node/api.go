package node

import (
	"github.com/commnerd/sw-ports"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"log"
	"fmt"
	// "os"
)

type Api struct {
	node *Node
    domain string       `json:"domain"`
    basePath string     `json:"base_path"`
    hostPort string     `json:"host_port"`
}

func InitApi(node *Node) *Api {
	api := Api{
		node: node,
		domain: "localhost",
		basePath: "/",
		hostPort: "80",
	}
	return &api
}

func (this *Api) Listen() {
	r := mux.NewRouter()

	r.HandleFunc("/", Welcome)
	r.HandleFunc("/register", Register)
	r.HandleFunc("/ports/next", NextPort)


    log.Fatal(http.ListenAndServe(":"+this.hostPort, r))
}

func Welcome(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, "Welcome to SpiderWeb Master!")
}

func Register(w http.ResponseWriter, request *http.Request) {
    // request.RemoteAddr
    fmt.Fprintf(w, "Registering stuff")
}

func NextPort(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, strconv.Itoa(ports.NextAvailPort()))
}
