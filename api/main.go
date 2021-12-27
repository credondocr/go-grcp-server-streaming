package main

import (
	"log"
	"net/http"
	"os"

	"github.com/credondocr/go-grcp-server-streaming/api/handlers"
	"github.com/credondocr/go-grcp-server-streaming/client"
)

func main() {
	l := log.New(os.Stdout, "users-api", log.LstdFlags)

	c := client.NewGRCPClient(l)
	uh := handlers.NewUsers(l, c)
	sm := http.NewServeMux()
	sm.Handle("/users", uh)

	http.ListenAndServe(":9000", sm)
}
