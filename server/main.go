package main

import (
	"context"
	Db "github.com/ChuvashPeople/contactBook/fakeDB"
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

	db := Db.Db{}
	v1API := NewContactBookServer(&db)
	//opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer()

	pb.RegisterContactBookServer(grpcServer, v1API)
	err = grpcServer.Serve(listener)
	if err != nil {
		grpclog.Fatalf("%v", err)
	}

}

type server struct {
	db *Db.Db
}

func NewContactBookServer(db *Db.Db) pb.ContactBookServer {
	return &server{db: db}
}

func (s *server) AddContact(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	s.db.AddContact(request)
	return &pb.AddResponse{
		Message: "Person " + request.Name + " added",
	}, nil
}

func (s *server) GetContact(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	contact, err := s.db.GetContact(request)
	if err != nil {
		grpclog.Fatalf("%v", err)
	}
	return &pb.GetResponse{
		Name: contact.Name,
		Id:   int64(contact.Id),
	}, nil
}
