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
	Label string
	Records map[id.Id]*Record
}

func (db *DB) AddCollection(label string) error {
	if _, present := db.Collections[label]; present {
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