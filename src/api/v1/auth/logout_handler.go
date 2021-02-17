package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LogoutResponse struct {
	Version string              `json:"version"`
	Routes  map[string][]string `json:"routes"`
}

var LogoutStruct *LogoutResponse = &LogoutResponse{}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(LoginStruct)

	fmt.Fprintf(w, "%s", resp)

	if err != nil {
		panic(err.Error())
	}
}
