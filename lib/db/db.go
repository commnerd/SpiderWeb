package db

import(
	"../message_bus"
	"../config"
)

var instance *DB

func init() {
	config.SetDefault("parody_span", 10)
	instance = New()
	message_bus.Register(instance)
}

type DB struct{
	Collections map[string]*Collection
}

func New() *DB {
	return &DB{
		Collections: make(map[string]*Collection, 0),
	}
}