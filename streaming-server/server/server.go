package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"

	pb "github.com/credondocr/go-grcp-server-streaming/proto"
)

type Server struct {
	pb.UnimplementedStreamingServiceServer
}

func (s Server) FetchResponse(in *pb.Request, srv pb.StreamingService_FetchResponseServer) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Printf("fetch response for : %s", in.Usernames)
	//use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	for i := 0; i < len(in.Usernames); i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()
			client := http.Client{}
			req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s", in.Usernames[count]), nil)
			if err != nil {
				log.Fatalln(err)
			}

			req.Header = http.Header{
				"Authorization": []string{fmt.Sprintf("token %s", os.Getenv("GITHUB_TOKEN"))},
			}

			res, err := client.Do(req)

			if err != nil {
				log.Fatalln(err)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatalln(err)
			}

			user := pb.User{}
			err = json.Unmarshal(body, &user)
			if err != nil {
				fmt.Println(err)
				return
			}

			resp := pb.Response{User: &user}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("Response #%v", &user)
			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}

	wg.Wait()
	return nil
}
