package node

import (
    "github.com/google/uuid"
    "io/ioutil"
	"net/http"
	"log"
    "fmt"
    "os"
)

type Node struct {
    environment string   `json:"environment"`
	id string            `json:"id"`
	role string          `json:"role"`
	registry []*Node 	 `json:"registry"`
	version string       `json:"version"`
    api *Api             `json:"api"`
    instances []*Instance `json:"instance"`
    volumes []*Volume     `json:"volumes"`
}

func NewNode() Node {
    initEnv()
    
	node := Node {
        environment: env["ENVIRONMENT"],
        id: uuid.New().String(),
       	role: env["NODE_ROLE"],
       	registry:make([]*Node, 256),
        version: "0.0.1",
        api: new(Api),
        instances: make([]*Instance, 1),
        volumes: make([]*Volume, 1),
    }

    node.api = InitApi(&node)

    return node
}

func (this *Node) Execute() {
	this.Register()
	this.api.Listen()
}

func (this *Node) Register() {
	if this.role == "root" {
        registry := this.registry
        this.registry = append(registry, this)
        return
    }
    fmt.Println(this.role)
    os.Exit(1)

    resp, err := http.Get("http://"+env["ROOT_NODE_URL"]+"/register")
    if err != nil {
        log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(string(body))
}
