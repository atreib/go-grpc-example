package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/atreib/go-grpc-example/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func InitUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// Helper wrapper to send a value as address
	f := func(s int32) *int32 {
		return &s
	}

	return &pb.User{
		Id:    "1",
		Name:  req.GetName(),
		Email: req.GetEmail(),
		Age:   f(req.GetAge()),
	}, nil
}

func (*UserService) AddUserServerStream(req *pb.User, stream pb.UserService_AddUserServerStreamServer) error {
	// Helper wrapper to send a value as address
	f := func(s int32) *int32 {
		return &s
	}

	stream.Send(&pb.UserStream{
		Status: "Initializing",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserStream{
		Status: "Inserting on DB",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserStream{
		Status: "Success",
		User: &pb.User{
			Id:    "2",
			Name:  req.GetName(),
			Email: req.GetEmail(),
			Age:   f(req.GetAge()),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUserClientStream(stream pb.UserService_AddUserClientStreamServer) error {
	// Helper wrapper to send a value as address
	f := func(s int32) *int32 {
		return &s
	}

	users := []*pb.User{} // Creating a empty list of Users

	for {
		req, err := stream.Recv() // Receiving the stream

		// If is the end of the stream
		if err == io.EOF {
			// Respond and close
			return stream.SendAndClose(&pb.Users{
				Users: users,
			})
		}

		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		// While we're receiving the stream, append to the list
		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
			Age:   f(req.GetAge()),
		})

		fmt.Println("Received user: ", req.GetName())
	}
}

func (*UserService) AddUserBiStream(stream pb.UserService_AddUserBiStreamServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error receiving client's stream: %v", err)
		}

		err = stream.Send(&pb.UserStream{
			Status: "Added",
			User:   req,
		})

		if err != nil {
			log.Fatalf("Error sending response to the client: %v", err)
		}
	}
}
