package models

import (
    "net/url"
    "testing"
    "reflect"
)

func TestNewConfig(t *testing.T) {
    master, _ := url.Parse("https://spiderweb.com")
    keyValMap := &config{
        "MASTER": master,
    }

    config := NewConfig()

    if !reflect.DeepEqual(keyValMap, config) {
        t.Errorf("Expected: \"%v\"\nGot: \"%v\"", keyValMap, config)
    }
}

func TestBasicStringStoreAndRetrieve(t *testing.T) {
    c := NewConfig()
    c.Set("biz", "baz")

    expected := "baz"
    got := c.Get("biz")

    if expected != got {
        t.Errorf("Expected: \"%v\"\nGot:\"%v\"", expected, got)
    }
}
