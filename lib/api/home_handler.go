package api

import (
	"encoding/json"
	"../config"
	"net/http"
	"../port"
	"../id"
	"fmt"
)

type HomeResponse struct{
	Version string	`json:"version"`
	Address string	`json:"address"`
	Ip string		`json:"ip"`
	Port int		`json:"port"`
}

var HomeStruct *HomeResponse

func init() {
	HomeStruct = &HomeResponse{
		Version: config.GetString("project_version"),
		Address: id.New().String(),
		Ip: "192.168.10.10",
		Port: port.Next(),
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(HomeStruct)

	fmt.Fprintf(w, "%s", resp)

	if err != nil {
		panic(err.Error())
	}
}