package main

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func failOnError(err error, msg string)()  {
	if err != nil {
		log.Fatal("%s: %s",msg,err)
	}
}

func main() {
	//connect to rabbitmq
	conn,err := amqp.Dial("amqp://guest:guest@192.168.8.131:5678/")
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

	//publish message
	payload := `{"code":0,"msg":"succeed","data":[1,2,3]}`
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:     "json/application",
			Timestamp:       time.Time{},
			Body:            []byte(payload),
		})
	log.Println("[x] send ",payload)
	failOnError(err,"Failed to publish a message")
}