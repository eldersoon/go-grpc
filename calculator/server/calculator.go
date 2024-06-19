package main

import (
	"context"

	pb "github.com/eldersoon/go-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest, ) (*pb.SumResponse, error) {

	return &pb.SumResponse{
		Sum: in.Number1 + in.Number2,
	}, nil
}