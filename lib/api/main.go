package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

var RunningNode node

func Serve() {
	r := mux.NewRouter()

	for path, handler := range(Routes) {
		r.HandleFunc(path, handler)
	}

    http.Handle("/", r)
}