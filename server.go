package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/yqlbu/grpc-go-demo/chat"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

const (
	TARGET_IP = "localhost:9000"
)

func main() {

	fmt.Printf("Server is listening to %s\n", TARGET_IP)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
