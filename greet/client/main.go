package main

import (
	"log"
	"strconv"

	"github.com/achmadAlli/go-grpc/greet/client/clientservice"
	pb "github.com/achmadAlli/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddress = "0.0.0.0:50001"

func main() {

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server : %v \n", err)
	}

	defer conn.Close()
	gc := pb.NewGreetingServiceClient(conn)

	greetClient := clientservice.NewGreetServiceClient(gc)

	for i := 0; i < 10; i++ {
		greetClient.Greet("Ali " + strconv.Itoa(i))

	}

}
