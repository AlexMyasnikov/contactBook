package main

import (
	"context"
	"fmt"
	pb "github.com/ChuvashPeople/contactBook/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"os"
	"strconv"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	args := os.Args

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		grpclog.Fatalf("%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	client := pb.NewContactBookClient(conn)

	i, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	request := &pb.AddRequest{
		Name: args[1],
		Id:   int64(i),
	}

	response, err := client.AddContact(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("%v", err)
	}

	fmt.Printf(response.Message)
}
