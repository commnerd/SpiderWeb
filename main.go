package main

import (
	"context"

	"github.com/containerd/containerd"
)

func main() {
	client, _ := containerd.New("/run/containerd/containerd.sock")
	defer client.Close()
	eventService := client.EventService()
	ctx := context.TODO()
	channel, _ := eventService.Subscribe(ctx)

	for {
		envelope := <-channel
		println(envelope.Event.TypeUrl)
	}
}
