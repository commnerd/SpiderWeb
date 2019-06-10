package models

import (
    "net/url"
    "testing"
)

func TestNewConfig(t *testing.T) {
    master, _ := url.Parse("https://spiderweb.com/")
    keyValMap := map[string]interface{}{
        "MASTER": master,
    }

    config := NewConfig()

    for k, v := keyValMap {
        if config[k] != v {
            t.Errorf("Expected: \"%v\"\nGot: \"%v\"", v, config[k])
        }
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
