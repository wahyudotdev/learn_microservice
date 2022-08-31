package config

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
	"log"
	"os"
)

var Port = func() string {
	return "0.0.0.0:" + os.Getenv("PORT")
}()

var Cors = cors.New(cors.Config{
	AllowMethods: "GET, POST, OPTIONS, PUT, DELETE",
	AllowOrigins: "*",
})

var ProductCService = func() *grpc.ClientConn {
	host := os.Getenv("PRODUCT_C_SERVICE")
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to ", host, err)
	}
	return conn
}()

var ProductQService = func() *grpc.ClientConn {
	host := os.Getenv("PRODUCT_Q_SERVICE")
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to ", host, err)
	}
	return conn
}()
