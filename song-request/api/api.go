package api

import "net/http"

// API struct
type API struct {
	twitchClient *http.Client
}

// NewAPI creates a new instance of an API
func NewAPI(db *struct{}) (*API, error) {
	return &API{}, nil
}
