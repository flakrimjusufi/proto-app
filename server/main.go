package main

import (
	"context"
	pb "github.com/flakrimjusufi/proto-app/person"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement person.GetPerson.
type server struct {
	pb.UnimplementedGetPersonServer
}

// GetPerson implements person.GetPerson
func (s *server) GetPerson(ctx context.Context, in *pb.Person) (*pb.Person, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Person{Name: "Hello " + in.GetName(), Age: in.GetAge()}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetPersonServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
