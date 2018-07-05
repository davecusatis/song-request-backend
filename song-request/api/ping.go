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
	tok, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}
	a.Aggregator.MessageChan <- &models.SongRequestMessage{
		MessageType: "load",
		Data:        models.MessageData{},
		Token:       token.CreateServerToken(tok),
	}
	w.Write([]byte("OK"))
}
