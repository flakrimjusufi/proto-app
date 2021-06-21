package main

import (
	"context"
	pb "github.com/flakrimjusufi/proto-app/person"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "Flakrim"
	defaultAge  = 27
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
	var age int32 = defaultAge
	if len(os.Args) > 2 {
		convertAge, _ := strconv.ParseInt(os.Args[2], 10, 0)
		age = int32(convertAge)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.PersonData(ctx, &pb.Person{Name: name, Age: age})
	if err != nil {
		log.Fatalf("could not get person data: %v", err)
	}
	log.Printf("Sent Person data to the server: Name %s, Age %d", r.GetName(), r.GetAge())

}
