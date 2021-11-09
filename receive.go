package main

import (
	"fmt"
	"log"

	c "github.com/rabbitmq/amqp091-go"
)

func main() {
	var queue string
	fmt.Print("Enter the sender queuename: ")
	fmt.Scanf("%s", &queue)

	conn, _ := c.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	q, _ := ch.QueueDeclare(queue, false, false, false, false, nil)

	message, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	loop := make(chan bool)

	go func() {
		for d := range message {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" message is arraving ....... ")
	<-loop
}
