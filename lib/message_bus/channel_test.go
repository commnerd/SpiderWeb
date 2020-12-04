package message_bus

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"fmt"
)

func TestSuccessfulMakeChannel(t *testing.T) {
	err := MakeChannel("test_successful_make_channel")
	assert.Equal(t, err, nil)
	assert.IsType(t, make(Channel), instance.channels["test_successful_make_channel"])
}

func TestFailMakeChannel(t *testing.T) {
	err := MakeChannel("test_fail_make_channel")
	assert.Equal(t, err, nil)
	assert.IsType(t, make(Channel), instance.channels["test_fail_make_channel"])
	err = MakeChannel("test_fail_make_channel")
	assert.Equal(t, err.Error(), fmt.Sprintf(channel_exists_error, "test_fail_make_channel"))
}

func TestSuccessfulGetChannel(t *testing.T) {
	MakeChannel("test_successful_get_channel")
	channel, err := GetChannel("test_successful_make_channel")
	assert.Equal(t, err, nil)
	assert.Equal(t, instance.channels["test_successful_make_channel"], channel)
}

func TestFailGetChannel(t *testing.T) {
	_, err := GetChannel("test_fail_get_channel")
	assert.Equal(t, err.Error(), fmt.Sprintf(channel_not_exists_error, "test_fail_get_channel"))
}

func TestChannelSend(t *testing.T) {
	channelName := "test_channel_successful_send"
	pkg := "channel send tested"

	MakeChannel(channelName)
	channel := instance.channels[channelName]
	channel.Send(pkg)
	msg := <- instance.channels[channelName]
	msgString, ok := msg.(string)
	assert.True(t, ok)
	assert.Equal(t, pkg, msgString)
}

func TestChannelReceiveCallback(t *testing.T) {
	channelName := "test_channel_receive_callback"
	pkg := "channel callback tested..."
	rs := ""

	MakeChannel(channelName)

	channel, _ := GetChannel(channelName)
	channel.Receive(func(msg string) {
		rs = msg
	})
	time.Sleep(time.Millisecond)
	channel.Send(pkg)
	time.Sleep(time.Millisecond)

	assert.Equal(t, pkg, rs)

	channel.Send("die")
}

func TestChannelReceiveMultipleCallbacks(t *testing.T) {
	channelName := "test_channel_receive_multiple_callbacks"
	pkg1 := "channel callback tested..."
	pkg2 := "channel callback tested... again..."
	rs := ""

	MakeChannel(channelName)

	channel, _ := GetChannel(channelName)
	channel.Receive(func(msg string) {
		rs = msg
	})
	time.Sleep(time.Millisecond)

	channel.Send(pkg1)
	time.Sleep(time.Millisecond)
	assert.Equal(t, pkg1, rs)

	channel.Send(pkg2)
	time.Sleep(time.Millisecond)
	assert.Equal(t, pkg2, rs)

	channel.Send("die")
}

func TestChannelReceivePackage(t *testing.T) {
	channelName := "test_channel_receive_callback"
	pkg := "channel callback tested..."

	MakeChannel(channelName)

	channel, _ := GetChannel(channelName)
	channel.Send(pkg)

	payload := channel.Receive()

	assert.Equal(t, pkg, payload)
}