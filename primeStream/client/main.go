package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/achmadAlli/go-grpc/primeStream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddress = "0.0.0.0:50003"

func main() {

	dial, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server: %v \n", err)
	}

	defer dial.Close()

	client := proto.NewGeneratePrimeServiceClient(dial)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	payload := 120
	stream, err := client.GetPrimeNumber(ctx, &proto.PrimeRequest{Number: int32(payload)})

	if err != nil {
		log.Fatalf("Failed to call getPrime service with err : %v \n", err)
	}

	values := []int32{}
	done := make(chan bool, 1)

	go func() {
		for {
			response, err := stream.Recv()

			if err == io.EOF {
				done <- true
				break
			}

			if err != nil {
				log.Fatalf("error occurred when receiving stream data: %v \n", err)
			}

			values = append(values, response.Prime)
		}
	}()

	<-done

	fmt.Printf("Prime values of %d is : %v \n", payload, values)
}
