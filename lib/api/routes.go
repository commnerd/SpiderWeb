package api

import (
	"net/http"

	"./db"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/": HomeHandler,
	"/register": RegisterHandler,
	"/db": db.DbRootHandler,
}