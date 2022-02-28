package main

import (
	"context"
	"fmt"
	"io"
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

	AddUserServerStream(client)
}

func AddUserServerStream(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "André",
		Email: "andre@mail.com",
	}

	resStream, err := client.AddUserServerStream(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not send the request to gRPC: %v", err)
	}

	// Eternal loop to read the stream
	for {
		stream, err := resStream.Recv() // Receiving messages

		// If we get END_OF_FILE, we break the loop
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not read the stream: %v", err)
		}

		fmt.Println("Status: ", stream.Status)
	}
}
