package datasource

import (
	"github.com/davecusatis/song-request-backend/song-request/models"
)

// Datasource is the datasource structs
type Datasource struct {
	Playlist map[string]models.Song
}

// NewDatasource returns a new datasource instance
func NewDatasource() *Datasource {

	return &Datasource{
		Playlist: make(map[string]models.Song),
	}
}

// GetPlaylist returns the playlist
func (d *Datasource) GetPlaylist() map[string]models.Song {
	return d.Playlist
}

// AddSongToPlaylist adds a song to the playlist
func (d *Datasource) AddSongToPlaylist(song models.Song) error {
	d.Playlist[song.Title+song.Artist] = song
	return nil
}

// RemoveSongFromPlaylist removes a song from playlist
func (d *Datasource) RemoveSongFromPlaylist(song models.Song) {
	delete(d.Playlist, song.Title+song.Artist)
}
