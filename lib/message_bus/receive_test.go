package message_bus

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGlobalReceivePackage(t *testing.T) {
	channelName := "test_global_receive_package"
	pkg := "channel callback tested..."

	MakeChannel(channelName)

	Send(channelName, pkg)

	payload := Receive(channelName)

	assert.Equal(t, pkg, payload)
}

func TestGlobalReceiveCallback(t *testing.T) {
	channelName := "test_global_receive_callback"
	pkg := "channel callback tested..."
	rs := ""

	MakeChannel(channelName)

	Receive(channelName, func(msg string) {
		rs = msg
	})
	time.Sleep(time.Millisecond)

	Send(channelName, pkg)
	time.Sleep(time.Millisecond)
	assert.Equal(t, pkg, rs)

	Send(channelName, "die")
}

func TestGlobalReceiveMultipleCallbacks(t *testing.T) {
	channelName := "test_global_receive_multiple_callbacks"
	pkg1 := "channel callback tested..."
	pkg2 := "channel callback tested... again..."
	rs := ""

	MakeChannel(channelName)
	Receive(channelName, func(msg string) {
		rs = msg
	})
	time.Sleep(time.Millisecond)

	Send(channelName, pkg1)
	time.Sleep(time.Millisecond)
	assert.Equal(t, pkg1, rs)

	Send(channelName, pkg2)
	time.Sleep(time.Millisecond)
	assert.Equal(t, pkg2, rs)

	Send(channelName, "die")

}