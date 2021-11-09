package main

import (
	"fmt"
	"log"

	c "github.com/rabbitmq/amqp091-go"
)

func main() {
	var queue, data string
	fmt.Print("Enter your queuename: ")
	fmt.Scanf("%s", &queue)

	fmt.Print("Enter your data to be send: ")
	fmt.Scanf("%s", &data)

	conn, _ := c.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	q, _ := ch.QueueDeclare(queue, false, false, false, false, nil)

	body := data
	ch.Publish(
		"",
		q.Name,
		false,
		false,
		c.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf("messages Sent %s", body)
}

i completed this task .Created two api sender and reciver, sender send the and the message  and this meassage is stored in the queue 
of rabbitmq and when reciver want to recive that mesaage they easily received it any time .  
