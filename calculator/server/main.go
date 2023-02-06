package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/achmadAlli/go-grpc/calculator/proto"
	"github.com/achmadAlli/go-grpc/calculator/server/service"
	"google.golang.org/grpc"
)

const port = 50002

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("Failed to listen port : %d \n", port)
	}

	log.Printf("Server listen on port : %d \n", port)

	s := grpc.NewServer()
	defer s.Stop()

	service := service.NewCalculatorServer()

	pb.RegisterCalculatorServiceServer(s, service)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve GRPC : %v \n", err)
	}

}
