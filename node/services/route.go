package services

import (
	"net/http"
)

type Route struct {
	Path string
	Methods []string
	Handler func(w http.ResponseWriter, request *http.Request)
}
