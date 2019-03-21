package node

import (
    "github.com/commnerd/SpiderWeb/node/services"
    "github.com/google/uuid"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "bytes"
    "time"
    "log"
    "fmt"
)

const VERSION = "0.0.1"

const (
    NODE_ROLE_ROOT = "root"
    NODE_ROLE_REGISTRY = "registry"
    NODE_ROLE_NODE = "node"
    NODE_ROLE_INIT = "init"
)

type Node struct {
    Id string                       `json:"id"`
    Env map[string]string
    Addr string                     `json:"address,omitempty"`
    PublicKey string
    PrivateKey string
    Role string                     `json:"role"`
    Version string                  `json:"version"`
    HostNode Node
    Services []*services.Service
    Registry []Node
}

func NewNode(env map[string]string) Node {
	node := Node {
        Id: uuid.New().String(),
        Env: env,
        HostNode: Node{Addr:env["ROOT_ADDR"]},
        Services: make([]*Service, 0),
       	Role: env["NODE_ROLE"],
       	Registry: make([]Node, 0),
        Version: VERSION,
    }
    pubBytes,privBytes := GenerateKeys()
    node.PublicKey = string(pubBytes)
    node.PrivateKey = string(privBytes)

    return node
}

func (this *Node) Run() {
    this.services = append(this.services, services.NewApi())
    if this.Role != NODE_ROLE_ROOT {
	   this.Hello()
    }
    for {
        time.Sleep(time.Minute)
        for _, service := range(this.Services) {
            if(!service.IsRunning()) {
                go service.Run()
            }
        }
    }
}

func (this *Node) Hello() {
    registerUrl := "http://"+this.HostNode.Addr+":"+this.HostNode.Api.HostPort+"/hello"

    data, err := json.Marshal(this)
    fmt.Println("Me: "+string(data))
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println("Sending 'hello' to "+registerUrl)
	req, err := http.NewRequest("POST", registerUrl, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        log.Fatalln(err)
	}

    respData, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
        log.Fatalln(err)
	}

    // Unmarshal
    var node Node
    err = json.Unmarshal(respData, &node)
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

    data, err = json.Marshal(this)
    fmt.Println("NewMe: "+string(data))
    if err != nil {
        log.Fatalln(err)
    }

    this.Services = append(this.Services, services.NewTunnel(this))
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

	req, err := http.NewRequest("POST", registerUrl, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        log.Fatalln(err)
	}
	resp.Body.Close()
}
