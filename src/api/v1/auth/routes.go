package auth

import (
	"net/http"
)

var Routes = map[string]map[string]func(w http.ResponseWriter, r *http.Request){
	"": RootEndpointMapper
	"/login":  LoginEndpointMapper,
	"/logout": LogoutEndpointMapper,
}
