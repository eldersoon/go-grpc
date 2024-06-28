package main

import (
	"fmt"
	"log"

	pb "github.com/eldersoon/go-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Println("GreetManyTimes was invoked with: ", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %v this is the %v message!", in.FirstName, i+1)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}