package datasource

import (
	"fmt"

	"github.com/davecusatis/song-request-backend/song-request/models"
)

// Datasource is the datasource structs
type Datasource struct {
	Playlist []models.Song
}

// NewDatasource returns a new datasource instance
func NewDatasource() *Datasource {
	return &Datasource{
		Playlist: []models.Song{},
	}
}

// GetPlaylist returns the playlist
func (d *Datasource) GetPlaylist() []models.Song {
	return d.Playlist
}

// AddSongToPlaylist adds a song to the playlist
func (d *Datasource) AddSongToPlaylist(song models.Song) error {
	d.Playlist = append(d.Playlist, song)
	return nil
}

// RemoveSongFromPlaylist removes a song from playlist
func (d *Datasource) RemoveSongFromPlaylist(song models.Song) error {
	for index, s := range d.Playlist {
		if (s.Title == song.Title) && (s.Artist == song.Artist) {
			d.Playlist = append(d.Playlist[:index], d.Playlist[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Song not found in playlist")
}
