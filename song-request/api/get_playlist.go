package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/models"
	"github.com/davecusatis/song-request-backend/song-request/token"
)

// GetPlaylist retrieves the current setlist
func (a *API) GetPlaylist(w http.ResponseWriter, req *http.Request) {
	// validate token
	tok, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error getting token %s", err)))
		return
	}

	// hit db for playlist
	// a.db.GetPlaylist()

	// update all clients with current state of the world (songlist + playlist)
	a.Aggregator.MessageChan <- &models.SongRequestMessage{
		MessageType: "playlistUpdated",
		Data: models.MessageData{
			Playlist: a.Datasource.Playlist,
		},
		Token: token.CreateServerToken(tok),
	}
	log.Printf("Playlist Get")
	w.Write([]byte("OK"))
}
