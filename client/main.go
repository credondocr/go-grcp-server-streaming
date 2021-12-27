package client

import (
	"log"

	"google.golang.org/grpc"

	pb "github.com/credondocr/go-grcp-server-streaming/proto"
)

type GRCPClient struct {
	l   *log.Logger
	SSC *pb.StreamingServiceClient
}

func NewGRCPClient(l *log.Logger) *GRCPClient {
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())

	if err != nil {
		l.Fatal("can not connect with server %v", err)
	}

	// create stream
	SSC := pb.NewStreamingServiceClient(conn)
	return &GRCPClient{l, &SSC}
}
