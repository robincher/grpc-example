package main

import (
	"log"
	"os"

	pb "github.com/robincher/grpc-example/grpc-proxy/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "robincher"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Echo(context.Background(), &pb.Message{Id: "1", Msg: name})
	if err != nil {
		log.Fatalf("Unable to greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Msg)
}
