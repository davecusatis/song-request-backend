package api

import (
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/token"
)

// Ping is the health check endpoint
func (a *API) Ping(w http.ResponseWriter, req *http.Request) {
	// validate token
	token.ExtractTokenFromHeader(req.Header)

	// update all clients with current state of the world (songlist + playlist)
	w.Write([]byte("OK"))
}
