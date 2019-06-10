package models

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "log"
)

type SchemaDefinition struct {
    Id string                 `json:"@id,omitempty"`
    Context map[string]string `json:"@context,omitempty"`
    Type string               `json:"@type,omitempty"`
    Graph []SchemaDefinition  `json:"@graph,omitempty"`
}

func NewFromSchema(url string) SchemaDefinition {
    res, err := http.Get(url + ".jsonld")
	if err != nil {
		log.Fatal(err)
	}
	schema, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

    def := SchemaDefinition{}
    if err := json.Unmarshal(schema, &def); err != nil {
		log.Fatal(err)
	}

    return def
}
