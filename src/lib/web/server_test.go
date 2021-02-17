package web

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServe(t *testing.T) {
	routes = make([]Route, 0)

	AddRoute(&testRoute{Path: "/foo"})

	go func() {
		Serve(":12345")
	}()

	client := &http.Client{}

	time.Sleep(time.Millisecond)

	resp, err := client.Get("http://127.0.0.1:12345/foo")

	out := make([]byte, 3)
	fmt.Println(resp.Body.Read(out))

	assert.Equal(t, error(nil), err)
	assert.Equal(t, "foo", string(out))
}
