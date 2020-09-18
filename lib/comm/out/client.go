package out

import (
	"../msg"
	"fmt"
)
type client struct{}

func (c *client) Send(m msg.Message) {
	fmt.Println(m)
}