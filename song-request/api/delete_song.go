package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/models"
	"github.com/davecusatis/song-request-backend/song-request/token"
)

// DeleteSong deletes a song
func (a *API) DeleteSong(w http.ResponseWriter, req *http.Request) {
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

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading data"))
		return
	}

	var songToDelete models.Song
	err = json.Unmarshal(body, &songToDelete)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading data"))
		return
	}

	a.Datasource.RemoveSongFromPlaylist(songToDelete)

	// hit db for playlist
	// a.db.GetPlaylist()
	a.Aggregator.MessageChan <- &models.SongRequestMessage{
		MessageType: "playlistUpdated",
		Data: models.MessageData{
			Playlist: a.Datasource.Playlist,
		},
		Token: token.CreateServerToken(tok),
	}
	w.Write([]byte("OK"))
}
