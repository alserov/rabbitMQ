package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Getting started with Rabbit 🥕")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Connected to Rabbit! 😉")

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish("", "TestQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello!"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully published a message 🤝")
}
