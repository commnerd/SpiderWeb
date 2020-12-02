package db

import (
	"errors"
	"fmt"

	"../id"
)

const (
	CollectionPresentError = "Collection '%s' already exists."
)

type Collection struct{
	DB *DB
	Label string
	Records map[id.Id]*Record
}

func (collection *Collection) HasRecord(id id.Id) bool {
	return collection.Records[id] != nil
}

func (collection *Collection) Redistribute() {
	for _, record := range collection.Records {
		record.Redistribute()
	}
}

func (db *DB) CollectionExists(label string) bool {
	if _, present := db.Collections[label]; present {
		return true
	}
	return false
}

func (db *DB) AddCollection(label string) error {
	if db.CollectionExists(label) {
		return errors.New(fmt.Sprintf(CollectionPresentError, label));
	}

	db.Collections[label] = &Collection{
		Label: label,
		Records: make(map[id.Id]*Record, 0),
	}

	return nil
}

func AddCollection(label string) error {
	return instance.AddCollection(label)
}