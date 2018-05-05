package api

import "net/http"

// Ping is the health check endpoint
func (a *API) Ping(w http.ResponseWriter, req *http.Request) {
	// validate secret
	w.Write([]byte("OK"))
}
