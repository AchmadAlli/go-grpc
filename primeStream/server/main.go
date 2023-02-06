package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/achmadAlli/go-grpc/primeStream/proto"
	"github.com/achmadAlli/go-grpc/primeStream/server/services"
	"google.golang.org/grpc"
)

const port = 50003

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("Failed to listen port : %d with error : %v \n", port, err)
	}

	log.Printf("Server is listening on port : %d \n", port)

	s := grpc.NewServer()
	primeService := services.NewPrimeService()

	pb.RegisterGeneratePrimeServiceServer(s, primeService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC with err : %v \n", err)
	}

}
