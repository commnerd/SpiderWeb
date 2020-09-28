package port

import (
	"strconv"
	"time"
	"net"
	"fmt"
)

const MIN = 49152
const MAX = 65535

func Available(port int) bool {
	ret := true
	timeout := time.Millisecond

	// Attempt to connect
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(port)), timeout)

	// If there was an error connecting, the port is available
	if err != nil {
		ret = true
	}

	if conn != nil {
		err = conn.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		ret = false
	}

	return ret
}