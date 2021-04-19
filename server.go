package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/yqlbu/grpc-go-demo/chat"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

const (
	TARGET_IP = "localhost:9000"
)

// define the service
func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}

// define the connection
func main() {

	fmt.Printf("Server is listening to %s\n", TARGET_IP)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
