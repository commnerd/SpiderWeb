package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net"
	"fmt"
	"log"
)

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
}

func NewApi(node *Node) *Api {
	api := &Api{ Node: node }
	api.Routes = []Route{
		Route{ Method: "GET", Path: "/", Handler: api.Welcome, Name: "home"},
	}
	return api
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

	l, err := net.Listen("unix", "/tmp/spiderweb.sock")
	if err != nil {
		log.Fatalln(err)
	}

	err = http.Serve(l, r)
	if err != nil {
		log.Fatalln(err)
	}
}



func (this *Api) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SpiderWeb!")
}