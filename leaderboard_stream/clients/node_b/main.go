package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/achmadAlli/go-grpc/leaderboard_stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddress = "0.0.0.0:50005"

func main() {
	cc, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server : %v \n", err)
	}

	s := pb.NewLeaderboardServiceClient(cc)

	stream, err := s.AddNewPlayer(context.Background())

	if err != nil {
		log.Fatalf("Failed to start streaming process : %v \n", err)
		return
	}

	requests := []*pb.Player{
		{Name: "user b1", Point: 49},
		{Name: "user b2", Point: 15},
		{Name: "user b3", Point: 33},
		{Name: "user b4", Point: 21},
		{Name: "user b5", Point: 88},
		{Name: "user b6", Point: 37},
	}

	done := make(chan bool)

	// send data go routint
	go func() {
		for _, req := range requests {
			log.Printf("add new player : %v", req)

			stream.Send(&pb.StoreNewLeaderboardRequest{
				NewPlayer: req,
			})
			time.Sleep(2 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				log.Printf("Streaming was ended")
				close(done)
				break
			}

			if err != nil {
				log.Fatalf("Failed to receive streaming response :  %v \n", err)
			}

			log.Printf("Current leader board is : %v \n", res.Leaderboard)
		}

	}()

	<-done

}
