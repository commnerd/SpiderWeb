package main

import (
	"github.com/joho/godotenv"
	"os"
)

var env map[string]string

var defaultEnv = map[string]string {
	"NODE_ROLE": "node",
	"ROOT_ADDR": "spiderweb.com",
	"API_PORT": "80",
	"API_BASE_PATH": "/",
}

func InitEnv() map[string]string {
	_ = godotenv.Load(".env")
	env = defaultEnv
	overrides, _ := godotenv.Read()
	for k, _ := range(env) {
		if val, ok := overrides[k]; ok {
			env[k] = val
		}
		if val := os.Getenv(k); val != "" {
			env[k] = val
		}
	}

	return env
}
