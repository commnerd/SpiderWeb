package web

import (
	"net/http"
	"runtime"
	"testing"

	"reflect"

	"fmt"

	"github.com/commnerd/SpiderWeb/src/lib/util"
	"github.com/stretchr/testify/assert"
)

type testRoute Route

func (r testRoute) Options(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "foo")
}

func (r testRoute) Get(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "foo")
}

func TestReflect(t *testing.T) {
	tRoute := reflect.ValueOf(testRoute{Path: "/foo"})
	f, ok := tRoute.Method(0).Interface().(func(w http.ResponseWriter, request *http.Request))
	if ok {
		panic(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	}

}

func TestGetMethods(t *testing.T) {
	routes = make([]Route, 0)
	AddRoute(testRoute{Path: "/foo"})
	methods := routes[0].getMethods()

	hasOptions, _ := util.InArray(http.MethodOptions, methods)
	assert.True(t, hasOptions)

	hasGet, _ := util.InArray(http.MethodGet, methods)
	assert.True(t, hasGet)

	hasPost, _ := util.InArray(http.MethodPost, methods)
	assert.True(t, !hasPost)

	routes = make([]Route, 0)
}

func TestAddRoute(t *testing.T) {
	routes = make([]Route, 0)

	route := testRoute{Path: "/foo"}

	AddRoute(route)

	assert.Equal(t, "/foo", routes[0].Path)
}
