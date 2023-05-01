package main

import (
	"context"
	"fmt"
	"log"
	"rpc-demo/protorpc/server/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8087", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewArithSeriveClient(conn)
	response, err := client.Multiply(context.Background(), &pb.ArithRequst{
		A: 100,
		B: 20,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response.Pro)
}
