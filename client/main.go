package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/finalproject/generator"
)

const (
	address     = "localhost:50052"
	defaultName = 1
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	response, err := c1.GenerateUUIDs(context.Background(), &pb.RequireUuidCount{Count: 3})

	log.Printf("Greeting: %+v", response.Message)

	response, err = c1.GenerateUUIDs(context.Background(), &pb.RequireUuidCount{Count: 2})

	log.Printf("Greeting: %+v", response.Message)

}
