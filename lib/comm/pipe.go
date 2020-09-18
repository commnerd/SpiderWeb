package comm

import (
	"github.com/google/uuid"
	"./msg"
)

type pipe struct{
	Id uuid.UUID
	Channel chan msg.Message
}