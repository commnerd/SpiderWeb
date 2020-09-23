package id

import "github.com/google/uuid"

func New() Id {
	return Id(uuid.New())
}