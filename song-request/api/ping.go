package api

import (
	"fmt"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/models"
	"github.com/davecusatis/song-request-backend/song-request/token"
	"github.com/davecusatis/song-request-backend/song-request/util"
)

// Ping is the health check endpoint
func (a *API) Ping(w http.ResponseWriter, req *http.Request) {
	// validate token
	token, err := token.ExtractTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
	}

	util.SendPubsubBroadcastMessage(&models.SongRequestMessage{
		MessageType: "playlistUpdated",
		Data: []models.Song{{
			Title:  "ttfaf",
			Artist: "dragonforce",
			Genre:  "bad",
			Game:   "gh3",
		}},
	}, token)

	// update all clients with current state of the world (songlist + playlist)
	w.Write([]byte("OK"))
}
