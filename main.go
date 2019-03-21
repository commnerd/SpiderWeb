package main

import (
	"github.com/commnerd/SpiderWeb/node"
)

func main() {
	env := InitEnv()
	n := node.NewNode(env)
	n.Run()
}
