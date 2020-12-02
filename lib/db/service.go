package db

import (
	"../config"
)

func (db *DB) UsesParody() bool {
	return db.ParodySpan() > 1
}

func (db *DB) ParodySpan() int {
	return config.GetInt("db_parody_span")
}

func (db *DB) TriggerParodyRedistrobution() {
	for _, collection := range db.Collections {
		collection.Redistribute()
	}
}