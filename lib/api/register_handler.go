package api

import (
	_ "encoding/json"
	"net/http"
	"strconv"
	"fmt"

	_ "../config"
	_ "../port"
	"../id"
)


type ResponseType int

const (
	Success = ResponseType(iota)
	Forward
	Denied
	Failed
)

type RegisterRequest struct{
	Id string `json:"id"`
	Type string `json:"type"`
}

type RegisterResponse struct{
	Status ResponseType `json:"status"`
	AdjustedId string	`json:"adjusted_id"`
	Mask id.Mask        `json:"mask"`
}

var RegisterStruct *RegisterResponse

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		RegisterOptions(w, r)
	case http.MethodPost:
		RegisterPost(w, r)
	case http.MethodPut:
		RegisterPut(w, r)
	default:
		code := http.StatusMethodNotAllowed
		w.WriteHeader(code)
		w.Write([]byte(fmt.Sprintf("%v - %v method not allowed",
			strconv.Itoa(code),
			r.Method,
		)))
	}
}

func RegisterOptions(w http.ResponseWriter, r *http.Request) {
	for _, val := range []string{http.MethodOptions, http.MethodPost, http.MethodPut} {
		w.Header().Add("Access-Control-Allow-Methods", val)
	}
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {

}

func RegisterPut(w http.ResponseWriter, r *http.Request) {

}