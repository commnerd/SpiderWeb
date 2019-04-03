package main

import (
	"github.com/gorilla/mux"
	"os/signal"
	"io/ioutil"
	"net/http"
	"syscall"
	"strconv"
	"net"
	"fmt"
	"log"
	"os"
)

const UnixSocketLocation = "/tmp/spiderweb.sock"

type Query struct {
	Label string
	Value string
}

type Route struct {
	Method string
	Path string
	Handler func(w http.ResponseWriter, r *http.Request)
	Name string
}

type Api struct{
	Node *Node
	Routes []Route
	Listener net.Listener
}

func NewApi(node *Node) *Api {
	port := 0
	role := node.GetRole()
	if(role == NodeRoleInit || role == NodeRoleRoot) {
		port = 80
	}

	api := &Api{
		Node: node,
		Listener: NewListener(port),
	}
	api.Routes = []Route{
		Route{ Method: "GET", Path: "/", Handler: api.Welcome, Name: "home"},
		Route{ Method: "GET", Path: "/node", Handler: api.GetNode, Name: "node"},
		Route{ Method: "POST", Path: "/register", Handler: api.RegisterNode, Name: "register"},
	}

	return api
}

func NewListener(port int) net.Listener {

	var listener net.Listener
	var err error

	if port < 1 {
		listener, err = net.Listen("unix", UnixSocketLocation)
	} else {
		listener, err = net.Listen("tcp", ":" + strconv.Itoa(port))
	}

	if err != nil {
		log.Fatalln(err)
	}

	return listener
}

func (this *Api) HydrateRoutes(r *mux.Router) {
	for _, route := range this.Routes {

		f := r.HandleFunc(route.Path, route.Handler)

		if route.Method != "" {
			f.Methods(route.Method)
		}

		if route.Name != "" {
			f.Name(route.Name)
		}
	}
}

func (this *Api) Run() {
	r := mux.NewRouter()
	this.HydrateRoutes(r)
	fmt.Println("Welcome to SpiderWeb!")

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func(c chan os.Signal) {
		// Wait for a SIGINT or SIGKILL:
	    sig := <-c
	    log.Printf("Caught signal %s: shutting down.", sig)
	    // Stop listening (and unlink the socket if unix type):
	    this.Kill()

	    os.Exit(0)
	}(sigc)

	log.Fatalln(http.Serve(this.Listener, r))
}

func (this *Api) Listen(port int)  {
	this.Listener = NewListener(port)
}

func (this *Api) Kill() {
    this.Listener.Close()
}

func (this *Api) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SpiderWeb!")
}

func (this *Api) GetNode(w http.ResponseWriter, r *http.Request) {
	json := this.Node.MarshalJSON()
	fmt.Fprintf(w, string(json))
}

func (this *Api) RegisterNode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if(err != nil) {
		log.Fatalln(err)
	}

	var node Node
	node.UnmarshalJSON(body)

	f, err := os.OpenFile("/root/.ssh/authorized_keys", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
	    log.Fatalln(err)
	}
	defer f.Close()

	if _, err = f.WriteString(node.PubKey); err != nil {
        log.Fatalln(err)
    }

	if node.Id == this.Node.Id {
		node.Role = NodeRoleRoot
		w.Write(node.MarshalJSON())
		return
	}

	// Established we are not root, what about registrar?

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        log.Fatalln(w, "userip: %q is not IP:port", r.RemoteAddr)
    }

	_, err = http.Get("http://"+ip)
	if err != nil {
		node.Role = NodeRoleRegistrar
	}

	_, err = http.Get("https://"+ip)
	if err != nil {
		node.Role = NodeRoleRegistrar
	}

	if node.Role == NodeRoleRegistrar {
		w.Write(node.MarshalJSON())
		return
	}

	node.Role = NodeRoleWorker
	w.Write(node.MarshalJSON())
	return
}
