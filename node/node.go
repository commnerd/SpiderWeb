package node

import (
    "github.com/google/uuid"
    "encoding/json"
    "net/http"
    "bytes"
)

type Node struct {
    Environment string    `json:"environment"`
    Id string             `json:"id"`
    Role string           `json:"role"`
    Registry []*Node      `json:"registry,omitempty"`
    Version string        `json:"version"`
    Api *Api              `json:"api"`
    Instances []*Instance `json:"instance,omitempty"`
    Volumes []*Volume     `json:"volumes,omitempty"`
}

func NewNode() Node {
    initEnv()
	node := Node {
        Environment: env["ENVIRONMENT"],
        Id: uuid.New().String(),
       	Role: env["NODE_ROLE"],
       	Registry: make([]*Node, 0),
        Version: "0.0.1",
        Api: new(Api),
        Instances: make([]*Instance, 0),
        Volumes: make([]*Volume, 0),
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

    registerUrl, _ := "http://"+env["ROOT_NODE_URL"]+"/register", "Content-Type: application/json"

    data, err := json.Marshal(this)
    if err != nil {
        panic(err)
    }

    var jsonStr = []byte(data)

	req, err := http.NewRequest("POST", registerUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}
