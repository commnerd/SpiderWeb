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
    Environment string    `json:"environment"`
	Id string             `json:"id"`
	Role string           `json:"role"`
	Registry []*Node 	  `json:"registry"`
	Version string        `json:"version"`
    Api *Api              `json:"api"`
    Instances []*Instance `json:"instance"`
    Volumes []*Volume     `json:"volumes"`
}

func NewNode() Node {
    initEnv()
    
	node := Node {
        Environment: env["ENVIRONMENT"],
        Id: uuid.New().String(),
       	Role: env["NODE_ROLE"],
       	Registry:make([]*Node, 1),
        Version: "0.0.1",
        Api: new(Api),
        Instances: make([]*Instance, 1),
        Volumes: make([]*Volume, 1),
    }

    node.Api = InitApi(&node)

    return node
}

func (this *Node) Execute() {
	this.Register()
	this.Api.Listen()
}

func (this *Node) Register() {
	if this.Role == "root" {
        registry := this.Registry
        this.Registry = append(registry, this)
        return
    }
    fmt.Println(this.Role)
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
