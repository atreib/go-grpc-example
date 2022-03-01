package main

import (
	"context"
	"fmt"
	"io"
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

	AddUserBiStream(client)
}

func AddUserBiStream(client pb.UserServiceClient) {
	stream, err := client.AddUserBiStream(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Creating our input
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

	/*
		In order to send and receive something at the same time,
		we're going to use Go Routines, which is a feature from the Golang
		It works as a thread, but it is controlled by Go
		So, we're going to create 2 Go Routines
		  - One responsible for sending
		  - One responsible for receiving
	*/

	// This anounymous functions is ran async
	// and its responsibility is to send
	go func() {
		for idx, req := range users {
			fmt.Println("Sending ", idx, "# user...")
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	// We're creating a var of the "Channel" type
	// A channel is a place where you send a communication between two different Go Routines
	// Right now, our channel has no value and no Go Routines attached, we're only declaring it
	wait := make(chan int)

	// This anounymous functions is ran async
	// and its responsibility is to receive
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}

			fmt.Println("Received response: ", res.GetStatus(), " - ", res.GetUser().GetName())
		}

		// And, now, as soon as we get out of our eternal loop (on the io.EOF)
		// We're going to close our Channel
		close(wait)
	}()

	// We'll use our Channel var as a blocker to our program flow, sending its value to nowhere
	<-wait
}
