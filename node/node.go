package node

import (
    "github.com/google/uuid"
    "io/ioutil"
	"net/http"
	"log"
)

type Node struct {
	id string            `json:"id"`
	role string          `json:"role"`
	registry []Node 	 `json:"registry"`
	version string       `json:"version"`    
    api *api.Api         `json:"api"`
    instances []Instance `json:"instance"`
    volumes []Volume     `json:"volumes"`
}

func NewNode() Node {
	node := Node {
        id: uuid.New().String(),
       	role: "init",
       	registry: make([]Node)
        version: "0.0.1",
        api: &Api{},
        instances: []Instance{},
        volumes: []Volume{},
    }

    return node
}

func (this *Node) Execute() {
	this.Register()
	this.api.Listen()
}

func (this *Node) Register() {
	if this.role == "root" {
        publicNodes = append(publicNodes, this)
        return 
    }

    resp, err := http.Get("http://"+env["ROOT_ADDR"]+"/register")
    if err != nil {
        log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(string(body))
}