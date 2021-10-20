package main

import (
	"context"
	"log"
	"net"

	"github.com/microgate-io/microgate-demo/user"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, new(UserServiceImpl))
	log.Println("serving gRPC", "addr", ":8000")
	grpcServer.Serve(lis)
}

type UserServiceImpl struct {
	user.UnimplementedUserServiceServer
}

func (s *UserServiceImpl) CheckUser(ctx context.Context, r *user.CheckUserRequest) (*user.CheckUserResponse, error) {
	log.Println("CheckUser", r.Username)
	return &user.CheckUserResponse{IsValid: true}, nil
}
