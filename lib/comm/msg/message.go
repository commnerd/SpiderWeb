package msg

import "github.com/google/uuid"

type Message struct{
	Id uuid.UUID
	Body interface{}
}