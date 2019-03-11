package node

import (
    "github.com/google/uuid"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "bytes"
    "log"
    "fmt"
)

const (
    NODE_ROLE_ROOT = "root"
    NODE_ROLE_REGISTRY = "registry"
    NODE_ROLE_VOLUME = "volume"
    NODE_ROLE_INSTANCE = "instance"
    NODE_ROLE_NODE = "node"
)

type Node struct {
    Id string                       `json:"id"`
    Addr string                     `json:"address,omitempty"`
    HostNode *Node                  `json:"host, omitempty"`
    Api *Api                        `json:"api"`
    Services map[string][]Service   `json:"services, omitempty"`
    PublicKey string                `json:"id_rsa_pub"`
    PrivateKey string               `json:"id_rsa"`
    Role string                     `json:"role"`
    Registry []*Node                `json:"registry,omitempty"`
    Version string                  `json:"version"`
}

func NewNode() Node {
    initEnv()
	node := Node {
        Id: uuid.New().String(),
        HostNode: &Node{Addr:env["ROOT_ADDR"],Api: &Api{HostPort: "80"}},
        Services: make(map[string][]Service, 0),
       	Role: env["NODE_ROLE"],
       	Registry: make([]*Node, 0),
        Version: "0.0.1",
    }
    pubBytes,privBytes := GenerateKeys()
    node.PublicKey = string(pubBytes)
    node.PrivateKey = string(privBytes)
    node.Api = InitApi(&node)

    if(env["NODE_ROLE"] == "root") {
        node.Addr = "root"
    }

    return node
}

func (this *Node) Execute() {
    if this.Role != NODE_ROLE_ROOT {
	   this.Hello()
    }
	this.Api.Run()
}

func (this *Node) Hello() {
    respJson := this.SendHello()
    this.ProcessHelloResponse(respJson)
}

func (this *Node) SendHello() string {
    registerUrl := "http://"+this.HostNode.Addr+":"+this.HostNode.Api.HostPort+"/hello"

    data, err := json.Marshal(this)
    fmt.Println("Me: "+string(data))
    if err != nil {
        log.Fatalln(err)
    }

    var jsonStr = []byte(data)

    fmt.Println("Sending 'hello' to "+registerUrl)
	req, err := http.NewRequest("POST", registerUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        log.Fatalln(err)
	}

    b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
        log.Fatalln(err)
	}

    return string(b)

}

func (this *Node) ProcessHelloResponse(respJson string) {
    // Unmarshal
    var node Node
    err := json.Unmarshal([]byte(respJson), &node)
    if err != nil {
        log.Fatalln(err)
    }

    this.Addr = node.Addr

    if this.HostNode.Addr != node.HostNode.Addr {
        fmt.Println("Trying again... "+this.HostNode.Addr+" is not "+node.HostNode.Addr)
        this.HostNode = node.HostNode
        this.Hello()
        return
    }

    this.HostNode = node.HostNode
    this.Role = node.Role

    data, err := json.Marshal(this)
    fmt.Println("NewMe: "+string(data))
    if err != nil {
        log.Fatalln(err)
    }


    tunnel := NewTunnel(this)
    this.Services["tunnels"] = append(this.Services["tunnels"], tunnel)
    fmt.Println("Starting tunnel.")
    tunnel.Run()
}

func (this *Node) Register() {
	if this.Role == NODE_ROLE_ROOT {
        registry := this.Registry
        this.Registry = append(registry, this)
        return
    }

    registerUrl := "http://"+env["ROOT_NODE_URL"]+"/register"

    data, err := json.Marshal(this)
    if err != nil {
        log.Fatalln(err)
    }

    var jsonStr = []byte(data)

	req, err := http.NewRequest("POST", registerUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        log.Fatalln(err)
	}
	resp.Body.Close()
}
