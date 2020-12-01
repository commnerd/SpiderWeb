package db

func (db *DB) Insert(record *Record) interface{} {
	return record
}

func Insert(record *Record) interface{} {
	instance.Insert(record)
	return record
}