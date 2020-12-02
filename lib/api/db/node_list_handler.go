package db

import (
	_ "encoding/json"
	_ "io/ioutil"
	"net/http"
	"strconv"
	"fmt"

	_ "../../message_bus"
	"../../config"
	_ "../../id"
)

func NodeListHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		NodeListOptions(w, r)
	case http.MethodGet:
		NodeListGet(w, r)
	default:
		code := http.StatusMethodNotAllowed
		w.WriteHeader(code)
		w.Write([]byte(fmt.Sprintf("%v - %v method not allowed",
			strconv.Itoa(code),
			r.Method,
		)))
	}
}

func NodeListOptions(w http.ResponseWriter, r *http.Request) {
	for _, val := range []string{http.MethodOptions, http.MethodGet} {
		w.Header().Add("Access-Control-Allow-Methods", val)
	}
}

func NodeListGet(w http.ResponseWriter, r *http.Request) {

	database := config.Get("DB").(db)
	resp := "["

	for index, sib := range database.GetSiblingNodes() {
		if index > 0 {
			resp += ","
		}
		resp += "{"
		resp += "\"id\":\"" + sib.GetId().String() + "\","
		resp += "\"recordCount\":\"" + strconv.Itoa(sib.GetRecordCount()) + "\""
		resp += "}"
	}

	resp += "]"
	// resp, err := RunningNode.NodeListChild(string(body))
	// if err != nil {
	// 	panic(err.Error())
	// }

	// respString, err := json.Marshal(resp)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Fprintf(w, "%s", respString)
	fmt.Fprintf(w, "%s", "woot")

}