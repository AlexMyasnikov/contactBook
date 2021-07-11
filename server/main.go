package main

import (
	"context"
	pb "github.com/ChuvashPeople/contactBook/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		grpclog.Fatalf("%v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterContactBookServer(grpcServer, &server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		grpclog.Fatalf("%v", err)
	}

}

type server struct {
}

func (s server) AddContact(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	panic("implement me")
}

func (s server) GetContact(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	panic("implement me")
}
