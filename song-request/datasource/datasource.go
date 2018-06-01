package datasource

import (
	"log"

	"github.com/davecusatis/song-request-backend/song-request/models"
)

// Datasource is the datasource structs
type Datasource struct {
	Playlist map[int]models.Song
}

// NewDatasource returns a new datasource instance
func NewDatasource() *Datasource {

	return &Datasource{
		Playlist: make(map[int]models.Song),
	}
}

// GetPlaylist returns the playlist
func (d *Datasource) GetPlaylist() map[int]models.Song {
	return d.Playlist
}

// AddSongToPlaylist adds a song to the playlist
func (d *Datasource) AddSongToPlaylist(song models.Song) error {
	d.Playlist[len(d.Playlist)] = song
	return nil
}

// RemoveSongFromPlaylist removes a song from playlist
func (d *Datasource) RemoveSongFromPlaylist(song models.Song) {
	index := -1
	for i, s := range d.Playlist {
		if song.Artist == s.Artist && song.Title == s.Title {
			log.Printf("SOng: %#v, index: %d", s, i)
			index = i
		}
	}
	if index > -1 {
		delete(d.Playlist, index)
	}

	log.Printf("PLAYLIS: %#v", d.Playlist)
}
