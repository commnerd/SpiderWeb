package initialization

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TokenResponse struct {
	Version string              `json:"version"`
	Routes  map[string][]string `json:"routes"`
}

var TokenStruct *TokenResponse = &TokenResponse{}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(TokenResponse{})

	fmt.Fprintf(w, "%s", resp)

	if err != nil {
		panic(err.Error())
	}
}
