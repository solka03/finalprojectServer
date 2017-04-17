package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/finalproject/generator"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os/exec"
	"sync"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

type Message struct {
	UUIDs          []string
	TotalGenerated int32
}

func (c *Message) increment() {
	mu.Lock()
	c.TotalGenerated += 1
	mu.Unlock()
}

var mu sync.Mutex
var sum = 0

// SayHello implements helloworld.GreeterServer
func (s *server) GenerateUUIDs(ctx context.Context, in *pb.RequireUuidCount) (*pb.UuidString, error) {

	genaretedUUIds := make([]string, int(in.GetCount()))

	for i := 0; i < int(in.GetCount()); i++ {
		out, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", out)
		genaretedUUIds[i] = string(out)
		mu.Lock()
		sum = sum + 1

		mu.Unlock()
	}

	m := Message{genaretedUUIds, int32(sum)}

	b, _ := json.Marshal(m)

	return &pb.UuidString{Message: string(b)}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
