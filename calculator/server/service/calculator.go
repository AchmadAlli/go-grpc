package service

import (
	"context"

	pb "github.com/achmadAlli/go-grpc/calculator/proto"
)

type calculatorServer struct {
	pb.UnimplementedCalculatorServiceServer
}

func NewCalculatorServer() *calculatorServer {
	return &calculatorServer{}
}

func (s *calculatorServer) Sum(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := float64(req.GetA() + req.GetB())
	return &pb.CalculationResponse{Result: result}, nil
}

func (s *calculatorServer) Minus(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := float64(req.GetA() - req.GetB())
	return &pb.CalculationResponse{Result: result}, nil
}

func (s *calculatorServer) Multiply(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := float64(req.GetA() * req.GetB())
	return &pb.CalculationResponse{Result: result}, nil
}

func (s *calculatorServer) Devide(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := float64(req.GetA()) / float64(req.GetB())
	return &pb.CalculationResponse{Result: result}, nil
}
