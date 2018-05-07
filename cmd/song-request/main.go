package main

import (
	"log"

	songrequest "github.com/davecusatis/song-request-backend/song-request"
	"github.com/davecusatis/song-request-backend/song-request/api"
)

func main() {
	var s struct{}
	api, err := api.NewAPI(&s)
	server, err := songrequest.NewServer(api)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}

	server.Start()
}
