package main

import (
	"log"

	pb "github.com/eldersoon/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Failed to connect: ", err)
	}

	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	// doGreet(client)
	doGreetManyTimes(client)
}