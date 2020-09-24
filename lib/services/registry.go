package services

import (
	"./term"
)

type registry struct {
	Services map[string]Service
}

var Registry *registry

func init() {
	Registry = &registry{
		Services: make(map[string]Service),
	}
	Registry.Services["term"] = term.Get()
}