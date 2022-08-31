package main

import (
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"log"
	"net"
	"product_c_srv/config"
	"product_c_srv/repository/sqlite_repo"
	"product_c_srv/services/product"
)

func main() {
	srv := grpc.NewServer()

	// Init message broker
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(config.MessageBrokerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer connectRabbitMQ.Close()
	// Open channel to rabbitMQ
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Init repository
	repo := sqlite_repo.New(config.Db, channelRabbitMQ)
	service := product.NewService(repo)

	product.RegisterProductCommandServiceServer(srv, service)
	l, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on port %s", config.Port)
	log.Fatal(srv.Serve(l))
}
