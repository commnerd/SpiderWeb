package db

func (db *DB) GetLabel() string {
	return "database"
}

func (db *DB) Send(msg interface{}) interface{} {
	return true
}