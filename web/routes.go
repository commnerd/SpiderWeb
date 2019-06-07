package main

import (
    "net/http"
    "fmt"
)

var Routes []route = []route{
    route{"/", HomeHandler},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    // fmt.Fprintf(w, "Category: %v\n", vars["category"])
    fmt.Fprintf(w, "Hello World")
}
