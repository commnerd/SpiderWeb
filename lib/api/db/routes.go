package db

import (
	"net/http"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/node_list": NodeListHandler,
}