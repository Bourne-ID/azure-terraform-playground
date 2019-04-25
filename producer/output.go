package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-service-bus-go"
)

type output interface {
	out()
}

type stdout struct {
	data        string
	destination string
}

type kafka struct {
	data        string
	destination string
}

type bus struct {
	data        string
	destination string
}

func (s stdout) out() {
	fmt.Println(s.data)
}

func (s kafka) out() {
	fmt.Println("Not Implemented")
}

func (s bus) out() {

	// Simple bus deliberate for now to prevent quick message generation - that comes later.
	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(s.destination))
	if err != nil {
		return
	}

	client, err := ns.NewQueue("test1")
	if err != nil {
		return
	}

	// Create a context to limit how long we will try to send, then push the message over the wire.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := client.Send(ctx, servicebus.NewMessageFromString(s.data)); err != nil {
		fmt.Println("FATAL: ", err)
	}
}
