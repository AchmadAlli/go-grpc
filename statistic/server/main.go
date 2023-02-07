package main

import (
	"fmt"
	"log"
	"net"

	"github.com/achmadAlli/go-grpc/statistic/proto"
	"github.com/achmadAlli/go-grpc/statistic/server/services"
	"google.golang.org/grpc"
)

const port = 50004

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("Failed to listen on : %d \n", port)
	}

	log.Printf("Server is listening on port : %d \n ", port)

	s := grpc.NewServer()

	averageService := services.NewAverageService()

	proto.RegisterStatisticProcessorServer(s, averageService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC of : %v \n", err)
	}
}
