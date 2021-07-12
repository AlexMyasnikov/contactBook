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

	request := args[1]
	switch request {
	case "add":
		{
			name := args[2]
			id, err := strconv.Atoi(args[3])
			if err != nil {
				panic(err)
			}
			request := &pb.AddRequest{
				Name: name, Id: int64(id),
			}
			response, err := client.AddContact(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			fmt.Println(response.Message)

		}
	case "get":
		{
			id, err := strconv.ParseInt(os.Args[2], 10, 64)
			if err != nil {
				panic(err)
			}
			request := &pb.GetRequest{Id: id}
			response, err := client.GetContact(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			fmt.Printf("Name of person with ID-%d: %s\n", response.Id, response.Name)

		}

	}

}
