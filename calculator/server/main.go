package main

import (
	"log"
	"net"

	pb "github.com/eldersoon/go-grpc/calculator/proto"
	"google.golang.org/grpc"
)

const address = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln("Failet to listen: ", err)
	}

	log.Println("Listen on ", address)

	serve := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(serve, &Server{})

	if err = serve.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}