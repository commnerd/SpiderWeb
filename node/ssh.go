package node

import (
	"os/exec"
	"syscall"
	"os"
)

type Tunnel struct{

}

func NewTunnel(node *Node) {
	binary, lookErr := exec.LookPath("ssh")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}