package db

import (
	"../config"
	"../id"
)

type Record struct{
	Collection *Collection
	CollectionLabel string
	Id id.Id
	Nodes []id.Id
	Body interface{}
	SyncService interface{}
}

func (record *Record) Redistribute() {
	node := config.Get("NODE").(node)
	targets := record.Collection.DB.GetSiblingNodes()[0:3]
	inTargets := false

	for _, target := range targets {
		if target.Id.String() == node.GetId().String() {
			inTargets = true
		}
	}

	if !inTargets && targets[2].RecordCount < record.Collection.DB.GetRecordCount() {
		record.xfer(targets[0].GetId())
	}
}

func (record *Record) xfer(id id.Id) {
	// Send to node id above
}