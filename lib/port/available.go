package port

import (
	"strconv"
	"time"
	"net"
	"fmt"
)

type Port int

const MIN = Port(49152)
const MAX = Port(65535)

func Available(port Port) bool {
	ret := true
	timeout := time.Millisecond

	// Attempt to connect
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(int(port))), timeout)

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