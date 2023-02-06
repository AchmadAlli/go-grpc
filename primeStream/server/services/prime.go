package services

import (
	"log"

	pb "github.com/achmadAlli/go-grpc/primeStream/proto"
)

type primeService struct {
	pb.UnimplementedGeneratePrimeServiceServer
}

func NewPrimeService() *primeService {
	return &primeService{}
}

func (s *primeService) GetPrimeNumber(req *pb.PrimeRequest, stream pb.GeneratePrimeService_GetPrimeNumberServer) error {
	log.Printf("GetPrimeNumber is invoked with : %v \n", req.Number)

	num := req.Number
	mod := int32(2)

	for num > 1 {
		log.Printf("GetPrimeNumber is processed with : %v \n", num)

		if num%mod == 0 {
			num /= mod

			stream.Send(&pb.PrimeResponse{
				Prime: mod,
			})

			continue
		}

		mod++
	}

	return nil
}
