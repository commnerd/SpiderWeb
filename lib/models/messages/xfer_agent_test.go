package messages

import (
    "github.com/commnerd/SpiderWeb/testing/fixtures"
    "github.com/commnerd/SpiderWeb/testing/tools"
    "testing"
)

func NewMsgXferAgentTest(t testing.T) {
    agent := NewMsgXferAgent()
    if _, ok := agent.(*MsgXferAgent); !ok {
        t.Errorf("Expecting *MsgXferAgent, Got %s", tools.GetInterfaceType(agent))
    }
}

func SendMessageTest(t testing.T) {
    agent := NewMsgXferAgent()
    msg := fixtures.NewLoginMessage()
    if agent.SendMessage(msg) != nil {
        t.Errorf("Failed to send message.")
    }
}
