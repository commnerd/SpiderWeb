package initialization

import (
	"fmt"
	"net/http"
	"strconv"
)

type ReverseTunnelRequestBody struct {
	Hostname string `json:"hostname"`
}
type ReverseTunnelResponse struct {
	Port int `json:"port"`
}

var ReverseTunnelMethodMapper = map[string]func(http.ResponseWriter, *http.Request){
	http.MethodOptions: ReverseTunnelOptionsHandler,
	http.MethodPost:    ReverseTunnelPostHandler,
}

var ReverseTunnelStruct *ReverseTunnelResponse

func ReverseTunnelHandler(w http.ResponseWriter, r *http.Request) {
	if ReverseTunnelMethodMapper[r.Method] != nil {
		ReverseTunnelMethodMapper[r.Method](w, r)
		return
	}

	code := http.StatusMethodNotAllowed
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf("%v - %v method not allowed",
		strconv.Itoa(code),
		r.Method,
	)))
}

func ReverseTunnelOptionsHandler(w http.ResponseWriter, r *http.Request) {

}

func ReverseTunnelPostHandler(w http.ResponseWriter, r *http.Request) {

}
