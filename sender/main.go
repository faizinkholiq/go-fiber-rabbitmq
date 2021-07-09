package main

import (
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// Define RabbitMQ server URL
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	connectionRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil{
		panic(err)
	}
	defer connectionRabbitMQ.Close()

	// Let's try by opening channel to our RabbitMQ
	// Instance over the connection we have already
	// established
	channelRabbitMQ, err := connectionRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// With the instance and declare Queues that we can 
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		"QueueService1", //Queue name
		true,			 //durable
		false,			 //
		false,
		false,
		nil,
	)
}