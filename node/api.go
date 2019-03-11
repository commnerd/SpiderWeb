package node

import (
	"github.com/commnerd/sw-ports"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"net"
	"log"
	"fmt"
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

func (this *Api) Run() {
	this.HandleFunc("/", this.Welcome)
	this.HandleFunc("/hello", this.Hello)
	this.HandleFunc("/env", this.Env)
	this.HandleFunc("/register", this.Register)
	this.HandleFunc("/node", this.Node)
	this.HandleFunc("/ports/next", this.NextPort)

	fmt.Println("Welcome to SpiderWeb on port "+this.HostPort+"!")
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

func (this *Api) Hello(w http.ResponseWriter, request *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var node Node
	err = json.Unmarshal(b, &node)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	ip, _, err := net.SplitHostPort(request.RemoteAddr)
    if err != nil {
        //return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)

        log.Fatalln(w, "userip: %q is not IP:port", request.RemoteAddr)
    }

	node.Addr = ip
	node.HostNode = this.node
	if(isPublic(node)) {
		node.Role = "registry"
	}

	output, err := json.Marshal(node)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (this *Api) Welcome(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, "Welcome to SpiderWeb Master!")
}

func (this *Api) Node(w http.ResponseWriter, request *http.Request) {
	e, _ := json.Marshal(this.node)
    fmt.Fprintf(w, string(e))
}

func (this *Api) Register(w http.ResponseWriter, request *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var node Node
	err = json.Unmarshal(b, &node)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(this)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
	this.node.Registry = append(this.node.Registry, &node)
}

func (this *Api) NextPort(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(w, strconv.Itoa(ports.NextAvailPort()))
}

func isPublic(node Node) bool {
	addr := node.Addr
	port := node.Api.HostPort
	base := node.Api.BasePath
	responseUrl := "http://"+addr+":"+port+base
	fmt.Println("Checking "+responseUrl)
    client := http.Client{}

    _, err := client.Get(responseUrl)
	if(err != nil) {
		fmt.Println("Ugh...")
		return false
	}
	fmt.Println("Yay!")
	return true
}
