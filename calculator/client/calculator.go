package main

import (
	"context"
	"log"

	pb "github.com/eldersoon/go-grpc/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {
	res, err := client.Sum(context.Background(), &pb.SumRequest{
		Number1: 10,
		Number2: 30,
	})

	if err != nil {
		log.Fatalln("Coldn't calculator: ", err)	
	}

	log.Println("Sum is: ", res.Sum)
}