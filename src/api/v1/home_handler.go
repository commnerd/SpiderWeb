package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HomeResponse struct {
	Version string   `json:"version"`
	Routes  []string `json:"routes"`
}

var HomeEndpointMapper = map[string]func(http.ResponseWriter, *http.Request){
	http.MethodOptions: HomeOptionsHandler,
	http.MethodGet:     HomeGetHandler,
}

func HomeOptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS, GET")
}

func HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(&HomeResponse{
		Version: "0.4.0",
		Routes:  RouteList,
	})

	fmt.Fprintf(w, "%s", resp)

	if err != nil {
		panic(err.Error())
	}
}
