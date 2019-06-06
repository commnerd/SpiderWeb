package models

import . "./messages"

// Message structure
type msg struct {
    type MessageType
    payload interface{}
}

// Craft a new message
func NewMessage(mt MessageType) Message
{
    var message msg = msg{
        type: mt
    }
}

// Get the message type
func (this msg) GetType() MessageType {
    return this.type
}

// Get the message payload
func (this msg) GetPayload() interface{} {

}

// Send the message
func (this msg) Send() {
    agent := messages.NewMsgXferAgent()
    agent.SendMessage(this)
}
