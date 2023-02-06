package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/achmadAlli/go-grpc/greet/proto"
)

func (*Server) Greeta(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("Greet service was invoked : %v \n", req)

	res := &pb.GreetingResponse{
		Greeting: fmt.Sprintf("Hello %s, how are you?", req.Name),
	}

	return res, nil
}
