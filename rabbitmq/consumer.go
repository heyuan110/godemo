package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string)()  {
	if err != nil {
		log.Fatalf("%s: %s ",msg,err)
	}
}

func main() {
	//connect to rabbitmq
	conn,err := amqp.Dial("amqp://bruce1:patpat@192.168.8.131:5678/")
	failOnError(err,"Failed to connect to RabbitMQ")
	defer conn.Close()

	//open channel
	ch,err := conn.Channel()
	failOnError(err,"Failed to open a channel")
	defer ch.Close()

	//declare a queue
	q, err := ch.QueueDeclare(
		"hellogolang",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err,"Failed to declare a queue")

	//Consume message
	msgs,err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err,"Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d:= range msgs {
			log.Printf("Received a message: %s",d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}