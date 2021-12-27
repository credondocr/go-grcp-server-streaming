package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/credondocr/go-grcp-server-streaming/proto"
	"github.com/credondocr/go-grcp-server-streaming/streaming-server/server"
)

func main() {
	// create listiner
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterStreamingServiceServer(s, server.Server{})

	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
