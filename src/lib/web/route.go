package web

import (
	"net/http"
	"reflect"
	"strings"
)

var routes = make([]Route, 0)

type Route struct {
	Path    string
	methods map[string]func(http.ResponseWriter, *http.Request)
}

func AddRoute(iRoute interface{}) {

	path := reflect.ValueOf(iRoute).Elem().FieldByName("Path").Interface().(string)

	route := Route{
		Path:    path,
		methods: make(map[string]func(http.ResponseWriter, *http.Request)),
	}

	for _, method := range []string{"Options", "Head", "Get", "Post", "Patch", "Delete", "Put"} {
		if reflect.ValueOf(iRoute).MethodByName(method).IsValid() {
			methodVal := reflect.ValueOf(iRoute).MethodByName(method).Interface().(func(http.ResponseWriter, *http.Request))
			route.methods[strings.ToUpper(method)] = methodVal
		}
	}

	routes = append(routes, route)
}

func (r Route) getMethods() []string {
	methods := make([]string, 0)

	for method := range r.methods {
		methods = append(methods, method)
	}

	return methods
}
