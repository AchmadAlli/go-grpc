package main

import (
	"context"
	"log"
	"time"

	pb "github.com/achmadAlli/go-grpc/statistic/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddress = "0.0.0.0:50004"

func main() {
	cc, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server: %v \n", err)
	}

	statisticClient := pb.NewStatisticProcessorClient(cc)

	req := []*pb.CalculationItem{
		{Number: 2},
		{Number: 13},
		{Number: 25},
		{Number: 1},
		{Number: 16},
	}

	stream, err := statisticClient.CalculateAverage(context.Background())

	if err != nil {
		log.Fatalf("Failed to stream average calculation :  %v \n", err)
	}

	for _, item := range req {
		log.Printf("Sending calculation item : %v", item.Number)

		err := stream.Send(item)
		time.Sleep(2 * time.Second)

		if err != nil {
			log.Fatalf("Failed while sending stream data : %v", err)
		}
	}

	result, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error occurred while receiving average result : %v", err)
	}

	log.Printf("Average result is : %v", result.Number)

}
