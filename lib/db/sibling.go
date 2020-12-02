package db

import (
	"../id"
)

type SiblingNode struct{
	Id id.Id
	RecordCount int
}

func (s SiblingNode) GetId() id.Id {
	return s.Id
}

func (s SiblingNode) GetRecordCount() int {
	return s.RecordCount
}