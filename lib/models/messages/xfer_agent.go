// The messages package provides message constructs to communicate between nodes
package messages

import (
    "fmt"
)

// The transfer agent that sends messages
type MsgXferAgent struct {}

func NewMsgXferAgent() *MsgXferAgent {
    return &MsgXferAgent{}
}

func (this MsgXferAgent) SendMessage(m Message) error {
    return nil
}
