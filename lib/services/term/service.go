package term

import (
	"os/exec"
)

type service struct{
	Command *exec.Cmd
	listener chan string
}

var s *service

func init() {
	s = &service{
		Command: exec.Command("/usr/bin/docker", "run", "--rm", "-it", "ubuntu", "bash"),
		listener: make(chan string),
	}

	s.Listen()
}

func Get() *service {
	return s
}

func (svc *service) GetLabel() string {
	return "Ubuntu Bash"
}

func (svc *service) Receiver() chan string {
	return svc.listener
}

func (svc *service) Listen() {
	svc.Command.Run()
}

func (svc *service) Send(msg interface{}) bool {
	s, ok := msg.(string)
	if ok {
		svc.listener <- s
		return true
	}
	return false
}