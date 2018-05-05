package api

import "net/http"

// GetSetlist retrieves the current setlist
func (a *API) GetSetlist(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
}
