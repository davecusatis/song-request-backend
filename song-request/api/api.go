package api

import (
	"github.com/davecusatis/song-request-backend/song-request/aggregator"
)

// API struct
type API struct {
	Aggregator *aggregator.Aggregator
}

// NewAPI creates a new instance of an API
func NewAPI() (*API, error) {
	a := aggregator.NewAggregator()
	a.Start()
	return &API{
		Aggregator: a,
	}, nil
}
