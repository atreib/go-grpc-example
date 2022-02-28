package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	AddUserClientStream(client)
}

func AddUserClientStream(client pb.UserServiceClient) {
	users := []*pb.User{
		{
			Id:    "1",
			Name:  "Andre",
			Email: "andre@mail.com",
		},
		{
			Id:    "2",
			Name:  "Andre2",
			Email: "andre2@mail.com",
		},
		{
			Id:    "3",
			Name:  "Andre3",
			Email: "andre3@mail.com",
		},
		{
			Id:    "4",
			Name:  "Andre4",
			Email: "andre@mail.com",
		},
	}

	stream, err := client.AddUserClientStream(context.Background())

	if err != nil {
		log.Fatalf("Error creating the request: %v", err)
	}

	for index, req := range users {
		fmt.Println("Sending the ", index, "# user")
		stream.Send(req)
		time.Sleep(time.Second * 1)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving the response: %v", err)
	}

	fmt.Println(res)
}
