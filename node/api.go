package node

import (
	"github.com/commnerd/sw-ports"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"strconv"
	"log"
	"fmt"
	// "os"
)

type Api struct {
	node *Node
	router *mux.Router
    domain string       `json:"domain"`
    basePath string     `json:"base_path"`
    hostPort string     `json:"host_port"`
}

func InitApi(node *Node) *Api {
	api := Api{
		node: node,
		router: mux.NewRouter(),
		domain: "localhost",
		basePath: "/",
		hostPort: "80",
	}
	return &api
}

func (this *Api) Listen() {
	this.HandleFunc("/", Welcome)
	this.HandleFunc("/register", Register)
	this.HandleFunc("/node", WhoYou)
	this.HandleFunc("/ports/next", NextPort)

    log.Fatal(http.ListenAndServe(":"+this.hostPort, this.router))
}

func (this *Api) HandleFunc(path string, f func(http.ResponseWriter, *http.Request) *mux.Route) {
	this.router.HandleFunc(path, f)
}

func Welcome(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, "Welcome to SpiderWeb Master!")
}

func WhoYou(w http.ResponseWriter, request *http.Request) {
	e := json.Marshal()
    fmt.Fprintf(w, "Welcome to SpiderWeb Master!")
}

func Register(w http.ResponseWriter, request *http.Request) {
    // request.RemoteAddr
    fmt.Fprintf(w, "Registering stuff")
}

func NextPort(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, strconv.Itoa(ports.NextAvailPort()))
}
