package models

import (
    "net/url"
    "log"
)

type config map[string]interface{}

type Config interface{
    Set(string, interface{})
    Get(string) interface{}
}

func NewConfig() *config {
    master, err := url.Parse("https://spiderweb.com")
    if err != nil {
        log.Fatal(err)
    }
    return &config{
        "MASTER": master,
    }
}

func (c *config) Get(k string) interface{} {
    m := *c
    return m[k]
}

func (c *config) Set(k string, v interface{}) {
    m := *c
    m[k] = v
}
