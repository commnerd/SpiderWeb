package node

import (
    "github.com/google/uuid"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "os/exec"
    "bytes"
    "log"
)

const (
    NODE_ROLE_ROOT = "root"
    NODE_ROLE_REGISTRY = "registry"
    NODE_ROLE_VOLUME = "volume"
    NODE_ROLE_INSTANCE = "instance"
    NODE_ROLE_NODE = "node"
)

type Node struct {
    Id string             `json:"id"`
    Ip string             `json:"ip,omitempty"`
    Services []*struct{}  `json:"services, omitempty"`
    PublicKey string      `json:"id_rsa_pub"`
    PrivateKey string     `json:"id_rsa"`
    Environment string    `json:"environment"`
    Role string           `json:"role"`
    Registry []*Node      `json:"registry,omitempty"`
    Version string        `json:"version"`
    Api *Api              `json:"api"`
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
    if this.Role != NODE_ROLE_ROOT {
	   this.Hello()
    }
	this.Api.Listen()
}

func (this *Node) Hello() {    
    respJson := this.SendHello()
    this.ProcessHelloResponse(respJson)
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
        log.Fatal(err)
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

    this.Id = node.Id
    this.Ip = node.Ip

    exec.Start("ssh -o ServerAliveInterval=300 -NR 2222:localhost:22 root@localhost")

}

func (this *Node) PromotePublic() {
    this.Role = "registry"
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
