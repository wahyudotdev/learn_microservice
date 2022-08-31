package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"product_q_srv/config"
	"product_q_srv/repository/sqlite_repo"
	"product_q_srv/services/product"
)

func main() {

	srv := grpc.NewServer()
	// Init repository
	repo := sqlite_repo.New(config.Db)
	service := product.NewService(repo)

	product.RegisterProductQueryServiceServer(srv, service)
	log.Println("Starting device management service...")
	l, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on port %s", config.Port)
	log.Fatal(srv.Serve(l))
}
