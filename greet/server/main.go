package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/achmadAlli/go-grpc/greet/proto"
	pb "github.com/achmadAlli/go-grpc/greet/proto"
	"google.golang.org/grpc"
)

const port = 50001

type Server struct {
	proto.GreetingServiceServer
}

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("Failed to listen : %v \n", err)
	}
	defer listener.Close()

	log.Printf("Server is listening on port : %d", port)

	s := grpc.NewServer()
	defer s.Stop()

	proto.RegisterGreetingServiceServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to server : %v \n", err)
	}

}

func (*Server) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("Greet service was invoked : %v \n", req)

	res := &pb.GreetingResponse{
		Greeting: fmt.Sprintf("Hello %s, how are you?", req.Name),
	}

	return res, nil
}
