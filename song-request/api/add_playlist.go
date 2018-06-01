package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/models"
	"github.com/davecusatis/song-request-backend/song-request/token"
)

// AddSong deletes a song
func (a *API) AddSong(w http.ResponseWriter, req *http.Request) {
	// validate token
	token, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error getting token %s", err)))
		return
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading data"))
		return
	}

	var songToAdd models.Song
	err = json.Unmarshal(body, &songToAdd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading data"))
		return
	}

	songToAdd.RequestedBy = token.ChannelID
	log.Printf("Song to add %#v", songToAdd)
	a.Datasource.AddSongToPlaylist(songToAdd)
	log.Printf("Playlist: %#v", a.Datasource.Playlist)
	// hit db for playlist
	// a.db.GetPlaylist()
	a.Aggregator.MessageChan <- &models.SongRequestMessage{
		MessageType: "playlistUpdated",
		Data: models.MessageData{
			Playlist: parsePlaylistSongs(a.Datasource.Playlist),
		},
		Token: token,
	}

	w.Write([]byte("OK"))
}

func parsePlaylistSongs(playlist map[int]models.Song) []models.Song {
	ret := make([]models.Song, len(playlist))
	for i := 0; i < len(playlist); i++ {
		ret[i] = playlist[i]
	}
	return ret
}
