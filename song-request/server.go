package songrequest

import (
	"log"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/api"
	cors "github.com/heppu/simple-cors"
	"goji.io"
	"goji.io/pat"
)

const (
	apiBase = "/api/v0"
)

// Server is the struct representing the doorman Server
type Server struct {
	Port string
	Mux  *goji.Mux
}

// NewServer returns a new instance of the doorman Server
func NewServer(api *api.API) (*Server, error) {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Post(apiBase+"/ping"), api.Ping)

	// playlist handlers
	mux.HandleFunc(pat.Get(apiBase+"/playlist"), api.GetPlaylist)
	mux.HandleFunc(pat.Delete(apiBase+"/playlist"), api.DeleteSong)
	mux.HandleFunc(pat.Put(apiBase+"/playlist"), api.AddSong)
	mux.HandleFunc(pat.Post(apiBase+"/playlist/skip"), api.SkipSong)

	// songlist handlers
	mux.HandleFunc(pat.Get(apiBase+"/songlist"), api.GetSonglist)
	mux.HandleFunc(pat.Post(apiBase+"/songlist"), api.PostSonglist)

	return &Server{
		Port: "3030",
		Mux:  mux,
	}, nil
}

// Start starts the webserver
func (s *Server) Start() {
	log.Printf("Starting server on port %s", s.Port)
	log.Fatal(http.ListenAndServe(":"+s.Port, cors.CORS(s.Mux)))
}
