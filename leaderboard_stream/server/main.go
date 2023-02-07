package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/achmadAlli/go-grpc/leaderboard_stream/proto"
	"github.com/achmadAlli/go-grpc/leaderboard_stream/server/services"
	"google.golang.org/grpc"
)

const port = 50005

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Printf("Failed to listen on port : %d, err : %v", port, err)
	}

	log.Printf("Server is listening on port : %d", port)

	s := grpc.NewServer()
	leaderBoardService := services.NewLeaderBoardService()

	pb.RegisterLeaderboardServiceServer(s, leaderBoardService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server server : %v", err)
	}
}
