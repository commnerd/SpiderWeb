package api

import (
	"net/http"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/": HomeHandler,
}