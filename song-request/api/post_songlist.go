package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/token"
)

// PostSonglist is how broadcasters send us songlists
func (a *API) PostSonglist(w http.ResponseWriter, req *http.Request) {
	// validate token
	token, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}

	if token.Role != "broadcaster" {
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

	log.Printf("body: %v", string(body))

	// verify songlist
	// save songlist

	// blast message to clients
	// a.Aggregator.MessageChan <- &models.SongRequestMessage{
	// 	MessageType: "songlistUpdated",
	// 	Data: models.MessageData{
	// 		Playlist: models.TestPlaylist(),
	// 		Songlist: models.TestSonglist(),
	// 	},
	// 	Token: token,
	// }
	w.Write([]byte("OK"))
}
