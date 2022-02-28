package services

import (
	"context"
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
