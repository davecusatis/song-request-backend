package api

import (
	"fmt"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/models"
	"github.com/davecusatis/song-request-backend/song-request/token"
)

// SkipSong is how broadcasters skips the current song
func (a *API) SkipSong(w http.ResponseWriter, req *http.Request) {
	// validate token
	token, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error getting token %s", err)))
		return
	}

	// hit db for playlist
	// a.db.GetPlaylist()
	a.Aggregator.MessageChan <- &models.SongRequestMessage{
		MessageType: "playlistUpdated",
		Data: models.MessageData{
			Playlist: models.TestSkipPlaylist(),
		},
		Token: token,
	}
	w.Write([]byte("OK"))
}
