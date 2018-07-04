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
	tok, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error getting token %s", err)))
		return
	}

	if tok.Role != "broadcaster" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized: not broadcaster"))
		return
	}

	if len(a.Datasource.Playlist) > 0 {
		a.Datasource.RemoveSongFromPlaylist(a.Datasource.Playlist[0])
	}

	// hit db for playlist
	// a.db.GetPlaylist()
	a.Aggregator.MessageChan <- &models.SongRequestMessage{
		MessageType: "playlistUpdated",
		Data: models.MessageData{
			Playlist: parsePlaylistSongs(a.Datasource.Playlist),
		},
		Token: token.CreateServerToken(tok),
	}
	w.Write([]byte("OK"))
}
