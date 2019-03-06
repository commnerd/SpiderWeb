package node

import (
	"github.com/joho/godotenv"
)

var env map[string]string

var defaultEnv = map[string]string {
	"ENVIRONMENT": "production",
	"NODE_ROLE": "node",
	"ROOT_NODE_URL": "spiderweb.com:80",
	"API_PORT": "80",
	"SERVICE_PORT": "22",
	"API_BASE_PATH": "/",
}

func initEnv() {
	_ = godotenv.Load(".env")
	var env map[string]string = defaultEnv
	overrides, _ := godotenv.Read()
	for k, v := range(env) {
		if val, ok := overrides[k]; ok {
			env[k] = val
		}
	}
}
