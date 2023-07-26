package main

import (
	"fmt"

	"github.com/gsouza97/go-expert-utils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	for i := 0; i < 100000; i++ {
		msg := fmt.Sprintf("Hello World! Iteration: %d", i+1)
		rabbitmq.Publish(ch, msg, "amq.direct")
	}
}
