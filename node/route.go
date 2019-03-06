package node

import (
	"net/http"
)

type Route struct {
	path string
	methods []string
	handler func(w http.ResponseWriter, request *http.Request)
}