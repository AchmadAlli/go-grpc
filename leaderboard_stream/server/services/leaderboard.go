package services

import (
	"context"
	"io"
	"log"
	"sort"

	pb "github.com/achmadAlli/go-grpc/leaderboard_stream/proto"
)

type leaderBoardService struct {
	pb.UnimplementedLeaderboardServiceServer
	leaderBoard []*pb.Player
}

func NewLeaderBoardService() *leaderBoardService {
	return &leaderBoardService{
		leaderBoard: []*pb.Player{},
	}
}

func (s *leaderBoardService) AddNewPlayer(stream pb.LeaderboardService_AddNewPlayerServer) error {
	log.Printf("Add new player was invoked")

	done := make(chan bool)
	go func() {
		for {
			req, err := stream.Recv()

			if err == io.EOF {
				done <- true
				break
			}

			if err != nil {
				log.Fatalf("Failed to do streaming data : %v \n", err)
			}

			s.UpdateLeaderBoard(req.NewPlayer)

			stream.Send(&pb.LeaderboardResponse{
				Leaderboard: s.leaderBoard,
			})
		}
	}()

	<-done

	return nil
}

func (s *leaderBoardService) GetLeaderboard(ctx context.Context, req *pb.EmptyParam) (*pb.LeaderboardResponse, error) {
	return &pb.LeaderboardResponse{Leaderboard: s.leaderBoard}, nil
}

func (s *leaderBoardService) UpdateLeaderBoard(newPlayer *pb.Player) {
	if newPlayer.GetName() == "" || newPlayer.GetPoint() == 0 {
		log.Fatalf("Failed to add new player due empty value : %v \n", newPlayer)
	}

	if len(s.leaderBoard) == 0 {
		newPlayer.Rank = 1
		s.leaderBoard = append(s.leaderBoard, newPlayer)
		log.Println("Success to add 1st player on leaderBoard")
	} else {
		newPlayer.Rank = int32(len(s.leaderBoard) + 1)
		s.leaderBoard = append(s.leaderBoard, newPlayer)
	}

	sort.Slice(s.leaderBoard, func(i, j int) bool {
		isGreater := s.leaderBoard[i].GetPoint() > s.leaderBoard[j].GetPoint()

		if isGreater {
			tempRank := s.leaderBoard[j].Rank
			s.leaderBoard[j].Rank = s.leaderBoard[i].Rank
			s.leaderBoard[i].Rank = tempRank
		}

		return isGreater
	})
}
