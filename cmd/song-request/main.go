package main

import (
	"log"

	"github.com/davecusatis/doorman/doorman/api"
	songrequest "github.com/davecusatis/song-request-backend/song-request"
)

func main() {
	api := api.NewAPI()
	server, err := songrequest.NewServer(api)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}

	server.Start()
}
