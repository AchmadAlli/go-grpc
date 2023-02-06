package services

import (
	"context"
	"log"
	"time"

	"github.com/achmadAlli/go-grpc/calculator/proto"
)

type calculatorServiceClient struct {
	client proto.CalculatorServiceClient
}

func NewCalculatorServiceClient(client proto.CalculatorServiceClient) *calculatorServiceClient {
	return &calculatorServiceClient{
		client: client,
	}
}

func (c *calculatorServiceClient) TestCalculatorServer() {
	var res *proto.CalculationResponse
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	cases := [][]int32{
		{1, 3}, {10, 2}, {4, 6},
	}

	for i := 0; i <= 3; i++ {
		for _, v := range cases {
			req := proto.CalculationRequest{
				A: v[0],
				B: v[1],
			}

			switch i {
			case 0:
				res, err = c.client.Sum(ctx, &req)
			case 1:
				res, err = c.client.Minus(ctx, &req)
			case 2:
				res, err = c.client.Multiply(ctx, &req)
			case 3:
				res, err = c.client.Devide(ctx, &req)
			}

			if err != nil {
				log.Printf("Error occurred : %v", err)
			}

			log.Printf("Calculation type %d result : %.2f \n", i, res.Result)
		}
	}
}
