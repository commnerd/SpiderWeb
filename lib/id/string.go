package id

import "github.com/google/uuid"

func (i Id) String() string {
	return uuid.UUID(i).String()
}