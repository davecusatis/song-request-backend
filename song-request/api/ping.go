package api

import (
	"fmt"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/models"
	"github.com/davecusatis/song-request-backend/song-request/token"
)

// Ping is the health check endpoint
func (a *API) Ping(w http.ResponseWriter, req *http.Request) {
	// validate token
	token, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}

	// get state

	// update all clients with current state of the world (songlist + playlist)
	// a.Aggregator.MessageChan <- &models.SongRequestMessage{
	// 	MessageType: "load",
	// 	Data:        nil,
	// 	Token:       token,
	// }

	a.Aggregator.MessageChan <- &models.SongRequestMessage{
		MessageType: "load",
		Data: models.MessageData{
			Playlist: models.TestPlaylist(),
			Songlist: models.TestSonglist(),
		},
		Token: token,
	}
	w.Write([]byte("OK"))
}
