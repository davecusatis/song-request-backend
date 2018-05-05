package api

// API struct
type API struct {
	db *struct{}
}

// NewAPI creates a new instance of an API
func NewAPI(db *struct{}) (*API, error) {
	return &API{}, nil
}
