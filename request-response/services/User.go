package services

import (
	"context"

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
