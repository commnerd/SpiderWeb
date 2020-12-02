package db

import (
	"../id"
)

type node interface{
	GetId() id.Id
}