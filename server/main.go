package main

import (
	"encoding/json"
	pb "finalprojectServer/generator"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	//"net"

	"net"
	"net/http"
	"strings"
	"sync"
)

const (
	port = ":9090"
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
	fmt.Println("Inside GenerateUUIDs")
	genaretedUUIds := make([]string, int(in.GetCount()))

	for i := 0; i < int(in.GetCount()); i++ {

		out := uuid.NewV4()
		fmt.Printf("%s", out)
		genaretedUUIds[i] = out.String()
		mu.Lock()
		sum = sum + 1

		mu.Unlock()
	}

	m := Message{genaretedUUIds, int32(sum)}

	b, _ := json.Marshal(m)

	return &pb.UuidString{Message: string(b)}, nil

}

func main() {
	fmt.Println("Starting Server")
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

	//http.HandleFunc("/", sayhelloName)       // set router
	//err := http.ListenAndServe(":9090", nil) // set listen port
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}
