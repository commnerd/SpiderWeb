package port

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"net"
)

func TestMinNext(t *testing.T) {
	waitForPort(MIN)

	assert.Equal(t, MIN, Next())
}

func TestNextNext(t *testing.T) {
	waitForPort(MIN)

	server := openPort(MIN)

	assert.Equal(t, MIN + 1, Next())

	closePort(server)
}

func TestNextTen(t *testing.T) {
	servers := make([]*http.Server, 0)
	ports := make([]net.Listener, 0)
	done := make(chan bool)

	go func() {
		for port := MIN; port <= MIN + 10; port++ {
			waitForPort(port)
			servers = append(servers, openPort(port))
		}
	}()

	go func() {
		<-done

		for _, server := range servers {
			closePort(server)
		}
	}()

	assert.Equal(t, Port(int(MIN) + len(ports)), Next())
	done <- true
}