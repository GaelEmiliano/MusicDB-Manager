// Package models provides definitions and operations related
// to music tracks.
// It includes structures and methods to handle track information
// such as artis, title album, track number, year, genre, and the file path.
package models

// Rola represents a song with various attributes extracted from id3v2 tags.
type Rola struct {
	artist string
	title  string
	album  string
	track  int
	year   int
	genre  string
	path   string
	id     int64
}

// NewRola creates and returns a new Rola with default values.
// The text fields (artist, title, album, genre, path) are initialized
// to "Unknow", and numeric fields (track, year, id) are initialized
// to default values (track = 0, year = 2024, id = 0)
func NewRola() *Rola {
	initial := "Unknown"
	return &Rola{
		artist: initial,
		title:  initial,
		album:  initial,
		track:  0,
		year:   2024,
		genre:  initial,
		path:   initial,
		id:     0,
	}
}

func (rola *Rola) GetArtist() string {
	return rola.artist
}

func (rola *Rola) GetTitle() string {
	return rola.title
}

func (rola *Rola) GetAlbum() string {
	return rola.album
}

func (rola *Rola) GetTrack() int {
	return rola.track
}

func (rola *Rola) GetYear() int {
	return rola.year
}

func (rola *Rola) GetGenre() string {
	return rola.genre
}

func (rola *Rola) GetPath() string {
	return rola.path
}

func (rola *Rola) GetID() int64 {
	return rola.id
}
