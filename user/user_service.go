package user

import (
	"context"

	pb "github.com/ebrickdev/ebrick/example/user/proto"
	"github.com/ebrickdev/ebrick/logger"
)

type UserServiceServer struct {
	pb.UnimplementedUserServer
	log logger.Logger
}

func NewUserServiceServer(log logger.Logger) *UserServiceServer {
	return &UserServiceServer{log: log}
}

func (s *UserServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	// Dummy implementation (replace with actual logic)
	s.log.Info("Creating account for user", logger.String("username", req.Username))
	account := &pb.Account{
		Id:       "user-1234",
		Username: req.Username,
		Email:    req.Email,
		Created:  1675641600, // Dummy UNIX timestamp
		Updated:  1675641600,
		Verified: false,
		Profile:  req.Profile,
	}
	return &pb.CreateResponse{Account: account}, nil
}

func (s *UserServiceServer) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	account := &pb.Account{
		Id:       req.Id,
		Username: "exampleuser",
		Email:    "user@example.com",
		Verified: true,
		Created:  1675641600,
		Updated:  1675641600,
	}
	return &pb.ReadResponse{Account: account}, nil
}
