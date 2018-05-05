package api

import "net/http"

// GetSonglist retrieves the current song list
func (a *API) GetSonglist(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
}
