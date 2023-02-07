package services

import (
	"io"
	"log"

	pb "github.com/achmadAlli/go-grpc/statistic/proto"
)

type averageService struct {
	pb.UnimplementedStatisticProcessorServer
}

func NewAverageService() *averageService {
	return &averageService{}
}

func (s *averageService) CalculateAverage(stream pb.StatisticProcessor_CalculateAverageServer) error {
	log.Println("Calculation average was invoked")

	var total float64 = 0
	counter := 0

	done := make(chan bool)

	go func() {
		for {
			streamItem, err := stream.Recv()

			if err == io.EOF {
				log.Println("Client streaming was ended")

				stream.SendAndClose(&pb.AverageResponse{
					Number: total / float64(counter),
				})

				done <- true
				break
			}

			if err != nil {
				log.Fatalf("Error while receive client stream data : %v \n", err)
			}

			total += streamItem.GetNumber()
			counter++
		}
	}()

	<-done

	return nil
}
