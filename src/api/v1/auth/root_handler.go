package auth

import (
	"net/http"
)

func AuthOptionsHandler(w http.ResponseWriter, r *http.Request) {
	for _, val := range []string{http.MethodOptions} {
		w.Header().Add("Access-Control-Allow-Methods", val)
	}
}
