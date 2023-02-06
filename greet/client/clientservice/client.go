package clientservice

import (
	"context"
	"log"

	pb "github.com/achmadAlli/go-grpc/greet/proto"
)

type GreetClient struct {
	client pb.GreetingServiceClient
}

func NewGreetServiceClient(client pb.GreetingServiceClient) *GreetClient {
	return &GreetClient{client: client}
}

func (c *GreetClient) Greet(name string) {
	log.Printf("greeting function invoked")

	ctx := context.Background()

	req, err := c.client.Greet(ctx, &pb.GreetingRequest{Name: name})

	if err != nil {
		log.Fatalf("Failed to do greet the server : %v \n", err)
	}

	log.Printf("Server response the greeting with : %v \n", req.Greeting)
}
