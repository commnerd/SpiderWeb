package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/events"
)

func main() {
	client, _ := containerd.New("/run/containerd/containerd.sock")
	defer client.Close()
	eventService := client.EventService()
	ctx := context.TODO()
	eventChannel, errorChannel := eventService.Subscribe(ctx)

	go func(channel <-chan *events.Envelope) {
		for {
			handleEvent(<-channel)
		}
	}(eventChannel)

	go func(channel <-chan error) {
		for {
			handleError(<-errorChannel)
		}
	}(errorChannel)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for _ = range c {
		os.Exit(0)
	}
}

func handleEvent(envelope *events.Envelope) {
	println(envelope.Event.TypeUrl)
}

func handleError(err error) {
	println(err)
}
