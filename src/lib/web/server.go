package web

import (
	"log"
	"net/http"
)

func Serve(basePath string) {
	for _, route := range routes {
		for method, handler := range route.methods {
			router.HandleFunc(route.Path, handler).Methods(method)
		}
	}
	log.Fatal(http.ListenAndServe(basePath, router))
}
