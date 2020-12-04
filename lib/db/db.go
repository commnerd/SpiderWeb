package db

import(
	"sort"

	"../config"
	"../id"
)

var instance *DB

type DB struct{
	ParentNode node
	MasterNode id.Id
	SiblingNodes []SiblingNode
	Collections map[string]*Collection
}

func New() *DB {
	return &DB{
		Collections: make(map[string]*Collection, 0),
	}
}

func Get() *DB {
	return instance
}

func (db *DB) GetSiblingNodes() []SiblingNode {
	sort.Slice(db.SiblingNodes, func(i, j int) bool {
		return db.SiblingNodes[i].RecordCount > db.SiblingNodes[j].RecordCount
	})
	return db.SiblingNodes
}

func (db *DB) HasCollection(label string) bool {
	return db.Collections[label] != nil
}

func (db *DB) GetRecordCount() int {
	count := 0
	for _, collection := range db.Collections {
		count += len(collection.Records)
	}
	return count
}

func init() {
	config.SetDefault("db_parody_span", 10)
	instance = New()
	config.Set("DB", instance)
}
