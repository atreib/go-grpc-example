package main

import (
	"context"
	"log"

	"github.com/atreib/go-grpc-example/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect with gRPC server: %v", err)
	}

	// The `defer` command observes the `connection` variable and, if it's not being used, closes it
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "André",
		Email: "andre@mail.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not send the request to gRPC: %v", err)
	}

	log.Println(res)
}
