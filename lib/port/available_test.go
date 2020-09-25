package port

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"strconv"
	"context"
	"../id"
	"time"
)

func init() {
	http.HandleFunc("/" + id.New().String(), func(w http.ResponseWriter, r *http.Request) {})
}

func TestAvailable(t *testing.T) {
	assert.True(t, Available(MIN))
}

func TestUnavailable(t *testing.T) {
	waitForPort(MIN)
	server := openPort(MIN)
	assert.True(t, !Available(MIN))
	closePort(server)
}

func openPort(port Port) *http.Server {
	server := &http.Server{Addr: ":" + strconv.Itoa(int(port))}
	go server.ListenAndServe()
	time.Sleep(time.Millisecond)
	return server
}

func closePort(server *http.Server) {
	server.Shutdown(context.Background())
	time.Sleep(time.Millisecond)
}

func waitForPort(port Port) {
	for !Available(port) {
		time.Sleep(time.Millisecond)
	}
}