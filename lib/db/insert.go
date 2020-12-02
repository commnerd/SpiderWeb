package db

import (
	"errors"
)

func (db *DB) Insert(record *Record) (*Record, error) {
	if !db.HasCollection(record.CollectionLabel) {
		db.AddCollection(record.CollectionLabel)
	}
	if db.Collections[record.CollectionLabel].HasRecord(record.Id) {
		return nil, errors.New("Record already exists in DB.")
	}
	db.Collections[record.CollectionLabel].Records[record.Id] = record
	return record, nil
}

func Insert(record *Record) (*Record, error) {
	return instance.Insert(record)
}