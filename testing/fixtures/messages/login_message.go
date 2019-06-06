package fixtures

import (
    "github.com/commnerd/SpiderWeb/models/messages"
    "github.com/commnerd/SpiderWeb/models"
)

func NewLoginRequestMessage() models.Message {
    messages.NewMessage(messages.LoginRequest)
}
