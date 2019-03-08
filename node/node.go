package node

import (
    "github.com/google/uuid"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "bytes"
)

type Node struct {
    Id string             `json:"id"`
    Ip string             `json:"ip,omitempty"`
    PublicKey string      `json:"id_rsa.pub"`
    PrivateKey string     `json:"id_rsa"`
    Environment string    `json:"environment"`
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
    pubBytes,privBytes := GenerateKeys()
    node.PublicKey = string(pubBytes)
    node.PrivateKey = string(privBytes)
    node.Api = InitApi(&node)

    return node
}

func (this *Node) Execute() {
	this.Hello()
	this.Api.Listen()
}

func (this *Node) Hello() {
    if this.Role == "root" {
        respJson := this.SendHello()
        this.ProcessHelloResponse(respJson)
    }
}

func (this *Node) SendHello() string {
    registerUrl := "http://"+env["ROOT_NODE_URL"]+"/hello"

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

    b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
        panic(err)
	}

    return string(b)

}

func (this *Node) ProcessHelloResponse(respJson string) {
    // Unmarshal
    var node Node
    err := json.Unmarshal([]byte(respJson), &node)
    if err != nil {
        panic(err)
    }

    this.Ip = node.Ip
}

func (this *Node) Register() {
	if this.Role == "root" {
        registry := this.Registry
        this.Registry = append(registry, this)
        return
    }

    registerUrl := "http://"+env["ROOT_NODE_URL"]+"/register"

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
