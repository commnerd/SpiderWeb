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
		hostPort: env["API_PORT"],
	}
	return &api
}

func (this *Api) Listen() {
	this.HandleFunc("/", this.Welcome)
	this.HandleFunc("/env", this.Env)
	this.HandleFunc("/register", this.Register)
	this.HandleFunc("/node", this.WhoYou)
	this.HandleFunc("/ports/next", this.NextPort)

    log.Fatal(http.ListenAndServe(":"+this.hostPort, this.router))
}

func (this *Api) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	this.router.HandleFunc(path, f)
}

func (this *Api) Env(w http.ResponseWriter, request *http.Request) {
	for k,v := range(env) {
		fmt.Fprintf(w, k + ": " + v)
	}
}

func (this *Api) Welcome(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, "Welcome to SpiderWeb Master!")
}

func (this *Api) WhoYou(w http.ResponseWriter, request *http.Request) {
	e, _ := json.Marshal(this.node.id)
    fmt.Fprintf(w, string(e))
}

func (this *Api) Register(w http.ResponseWriter, request *http.Request) {
    // request.RemoteAddr
    fmt.Fprintf(w, "Registering stuff")
}

func (this *Api) NextPort(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, strconv.Itoa(ports.NextAvailPort()))
}
