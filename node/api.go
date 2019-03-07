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
    Domain string       `json:"domain"`
    BasePath string     `json:"base_path"`
    HostPort string     `json:"host_port"`
}

func InitApi(node *Node) *Api {
	api := Api{
		node: node,
		router: mux.NewRouter(),
		Domain: "localhost",
		BasePath: "/",
		HostPort: env["API_PORT"],
	}
	return &api
}

func (this *Api) Listen() {
	this.HandleFunc("/", this.Welcome)
	this.HandleFunc("/env", this.Env)
	this.HandleFunc("/register", this.Register)
	this.HandleFunc("/node", this.Node)
	this.HandleFunc("/ports/next", this.NextPort)

	fmt.Println("Welcome to SpiderWeb!")
    log.Fatal(http.ListenAndServe(":"+this.HostPort, this.router))
}

func (this *Api) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	this.router.HandleFunc(path, f)
}

func (this *Api) Env(w http.ResponseWriter, request *http.Request) {
	for k,v := range(env) {
		fmt.Fprintf(w, k + ": " + v + "\n")
	}
}

func (this *Api) Welcome(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, "Welcome to SpiderWeb Master!")
}

func (this *Api) Node(w http.ResponseWriter, request *http.Request) {
	e, _ := json.Marshal(this.node)
    fmt.Fprintf(w, string(e))
}

func (this *Api) Register(w http.ResponseWriter, request *http.Request) {
	var node Node
	request.ParseForm()
	for key,_ := range request.Form {
		err := json.Unmarshal([]byte(key), &node)
		if err != nil {
			log.Println(err.Error())
		}
	}
	this.node.Registry = append(this.node.Registry, &node)
}

func (this *Api) NextPort(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, strconv.Itoa(ports.NextAvailPort()))
}
