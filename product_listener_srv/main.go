package main

import (
	"github.com/streadway/amqp"
	"log"
	"product_listener_srv/config"
	"product_listener_srv/consumer"
	"product_listener_srv/repository/sql_repo"
)

func main() {
	// Init message broker
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(config.MessageBrokerUrl)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Open channel to rabbitMQ
	channel, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")
	forever := make(chan bool)
	repo := sql_repo.New(config.Db)
	consume := consumer.New(repo, channel)
	consume.AddProduct()
	<-forever
}
