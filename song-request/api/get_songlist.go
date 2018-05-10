package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/models"
	"github.com/davecusatis/song-request-backend/song-request/token"
)

// GetSonglist retrieves the current song list
func (a *API) GetSonglist(w http.ResponseWriter, req *http.Request) {
	// validate token
	token, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error getting token %s", err)))
		return
	}

	// hit db for playlist
	// a.db.GetPlaylist()
	// update all clients with current state of the world (songlist + playlist)
	a.Aggregator.MessageChan <- &models.SongRequestMessage{
		MessageType: "songlistUpdated",
		Data: models.MessageData{
			Songlist: models.TestSonglist(),
		},
		Token: token,
	}

	log.Printf("Songlist Get")
	w.Write([]byte("OK"))
}
