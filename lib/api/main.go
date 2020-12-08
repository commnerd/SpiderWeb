package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"fmt"
	"log"
	"os"
)

var RunningNode node

func Serve() {
	r := mux.NewRouter()

	for path, handler := range(Routes) {
		r.HandleFunc("/api/v1" + path, handler)
	}

	r.NotFoundHandler = http.FileServer(http.Dir("./frontend"))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)  // for example /home/user

	log.Println("Serving SpiderWeb API.")

	log.Fatal(srv.ListenAndServe())
}