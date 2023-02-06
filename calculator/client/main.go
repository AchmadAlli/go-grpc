package main

import (
	"log"

	"github.com/achmadAlli/go-grpc/calculator/client/services"
	"github.com/achmadAlli/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddress = "0.0.0.0:50002"

func main() {

	client, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial the server: %v \n", err)
	}

	cc := proto.NewCalculatorServiceClient(client)

	serviceClient := services.NewCalculatorServiceClient(cc)
	serviceClient.TestCalculatorServer()
}
