package main

import (
	"context"
	"log"

	pb "github.com/eldersoon/go-grpc/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Elderson Gamma",
	})

	if err != nil {
		log.Fatalln("Coldn't greet: ", err)	
	}

	log.Println("Greeting: ", res.Result)
}