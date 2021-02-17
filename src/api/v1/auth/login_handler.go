package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginResponse struct {
	Version string              `json:"version"`
	Routes  map[string][]string `json:"routes"`
}

var LoginStruct *LoginResponse = &LoginResponse{}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(LoginStruct)

	fmt.Fprintf(w, "%s", resp)

	if err != nil {
		panic(err.Error())
	}
}
