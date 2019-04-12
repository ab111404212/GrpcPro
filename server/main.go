//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// protoc -I ../pb-proto --go_out=plugins=grpc:../pb-proto ../pb-proto/*.proto

package main

import (
	"context"
	"log"
	"net"

	proto "github.com/ab111404212/pb-proto"
	"google.golang.org/grpc"
)

const (
	port               = ":50051"
	USER_LOGIN_SUCCEDD = iota
	USER_LOGIN_FAILED
)

// server is used to implement helloworld.GreeterServer.
type serverLogin struct{}

// SayHello implements helloworld.GreeterServer
func (s *serverLogin) LoginRequest(ctx context.Context, args *proto.UserLoginRequest) (*proto.UserLoginResponse, error) {
	log.Printf("Received: %v", args.GetId())
	return &proto.UserLoginResponse{Code: USER_LOGIN_SUCCEDD, Data: &proto.UserData{
		Name:   "Miko",
		Coin:   args.GetId() + 9999,
		Equips: []string{"knife", "sword"},
	}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterServerServer(s, &serverLogin{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
