package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	router := mux.NewRouter()
	for route, mapping := range Routes {
		for method, handler := range mapping {
			router.HandleFunc(route, handler).Methods(method)
		}
	}
	log.Fatal(http.ListenAndServe(":6444", router))
}
