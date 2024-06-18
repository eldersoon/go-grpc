package main

import (
	"context"

	pb "github.com/eldersoon/go-grpc/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest, ) (*pb.CalculatorResponse, error) {

	return &pb.CalculatorResponse{
		Sum: in.Number1 + in.Number2,
	}, nil
}