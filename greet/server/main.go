package main

import (
	"log"
	"net"

	pb "github.com/eldersoon/go-grpc/greet/proto"
	"google.golang.org/grpc"
)

const address = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln("Failed to listen the server: ", err)
	}

	log.Println("Listening on ", address)

	serve := grpc.NewServer()
	pb.RegisterGreetServiceServer(serve, &Server{})

	if err = serve.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}