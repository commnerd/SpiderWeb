package api

import (
	"fmt"
	"net/http"
)

var Routes = map[string]map[string]func(w http.ResponseWriter, r *http.Request){
	"/":     HomeEndpointMapper,
	"/auth": auth.AuthEndpointMapper,
	// "/init": initialization.RootHandler,
}

var RouteList = make([]string, 0)

func init() {
	for route, mapper := range Routes {
		for method := range mapper {
			RouteList = append(RouteList, fmt.Sprintf("%s %s", method, route))
		}
	}
	// for route := range auth.Routes {
	// 	RouteList[route] = []string{http.MethodOptions}
	// }
}
