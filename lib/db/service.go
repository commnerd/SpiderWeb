package db

import (
	"../config"
)

func (db *DB) UsesParody() bool {
	return true
}

func (db *DB) ParodySpan() int {
	return config.GetInt("parody_span")
}

func (db *DB) TriggerParodyRedistrobution() {

}