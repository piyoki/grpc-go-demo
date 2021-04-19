package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/yqlbu/grpc-go-demo/chat"
)

const (
	TARGET_IP = "localhost:9000"
)

// define the connection to the server
func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(TARGET_IP, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	// Contact the server and print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.SayHello(ctx, &pb.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
