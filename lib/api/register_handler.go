package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"fmt"

	_ "../config"
	_ "../port"
	"../id"
)


type RegisterResponseType int

const (
	Success = RegisterResponseType(iota)
	Forward
	Denied
	Failed
)

type RegisterRequest struct{
	Id string `json:"id"`
}

func (rr RegisterRequest) RegisterChild(child node) (*RegisterResponse, error) {
	return nil, nil
}

type RegisterResponse struct{
	Status RegisterResponseType `json:"status"`
	Version string              `json:"version"`
	AdjustedId string           `json:"address,omitempty"`
	Mask id.Mask                `json:"mask,omitempty"`
	Ip string                   `json:"ip"`
	Port int                    `json:"port"`
	PublicRsa string            `json:"pub_rsa,omitempty"`
}

var response RegisterResponse

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		RegisterOptions(w, r)
	case http.MethodPost:
		RegisterPost(w, r)
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
	for _, val := range []string{http.MethodOptions, http.MethodPost} {
		w.Header().Add("Access-Control-Allow-Methods", val)
	}
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err.Error())
	}

	registrantNode := RegisterRequest{}
	_ = json.Unmarshal(body, &registrantNode)
	if err != nil {
		panic(err.Error())
	}

	resp, err := RunningNode.RegisterChild(string(body))
	if err != nil {
		panic(err.Error())
	}

	respString, err := json.Marshal(resp)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "%s", respString)
}