package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func serve() {
	r := mux.NewRouter()

	for path, handler := range(Routes) {
		r.HandleFunc(path, handler)
	}

    http.Handle("/", r)
}