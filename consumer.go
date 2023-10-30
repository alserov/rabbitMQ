package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer has started 🔊")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume("TestQueue", "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received: ", string(d.Body), "📩")
		}
	}()

	fmt.Println("Successfully connected to Rabbit 😉")
	fmt.Println("Waiting for messages")
	<-forever
}
