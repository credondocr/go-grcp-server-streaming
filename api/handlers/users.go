package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/credondocr/go-grcp-server-streaming/client"
	pb "github.com/credondocr/go-grcp-server-streaming/proto"
)

type Users struct {
	l *log.Logger
	c *client.GRCPClient
}

func NewUsers(l *log.Logger, c *client.GRCPClient) *Users {
	return &Users{l, c}
}

func (u *Users) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	iu := &[]string{}
	err := json.NewDecoder(req.Body).Decode(iu)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		u.l.Fatalf("can not connect with server %v", err)
	}
	defer conn.Close()

	// create stream
	SSC := pb.NewStreamingServiceClient(conn)
	in := &pb.Request{Usernames: *iu}
	stream, err := SSC.FetchResponse(context.Background(), in)
	if err != nil {
		u.l.Fatalf("open stream error %v", err)
	}

	usr := make(chan *pb.Response)

	go func() {
		for {
			r, err := stream.Recv()

			if err == io.EOF {
				close(usr)
				return
			}
			if err != nil {
				u.l.Fatalf("cannot receive %v", err)
			}
			usr <- r
		}
	}()
	r := []*pb.Response{}
	for u := range usr {
		r = append(r, u)
	}
	res, err := json.Marshal(r)
	if err != nil {
		u.l.Fatalf("can not covert %v", err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(res)
}
