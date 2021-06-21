package main

import (
	"context"
	pb "github.com/flakrimjusufi/proto-app/person"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "Flakrim"
	age         = 27
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGetPersonClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.PersonData(ctx, &pb.Person{Name: name, Age: age})
	if err != nil {
		log.Fatalf("could not get person data: %v", err)
	}
	log.Printf("Message: %s", r.GetMessage())

}
