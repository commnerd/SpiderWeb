package db

import (
	"../../id"
)

type sibling interface{
	GetId() id.Id
	GetRecordCount() int
}

type db interface{
	GetSiblingNodes() []sibling
}