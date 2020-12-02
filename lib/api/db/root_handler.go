package db

import (
	"net/http"
	"strconv"
	"fmt"
)

func DbRootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/db" {
		handler := Routes[r.URL.Path[3:]]
		handler(w, r)
		return
	}
	switch r.Method {
	case http.MethodOptions:
		DbRootOptions(w, r)
	default:
		code := http.StatusMethodNotAllowed
		w.WriteHeader(code)
		w.Write([]byte(fmt.Sprintf("%v - %v method not allowed",
			strconv.Itoa(code),
			r.Method,
		)))
	}
}

func DbRootOptions(w http.ResponseWriter, r *http.Request) {
	for _, val := range []string{http.MethodOptions} {
		w.Header().Add("Access-Control-Allow-Methods", val)
	}
}