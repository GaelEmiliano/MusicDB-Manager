// Package models provides definitions and operations related
// to music tracks.
// It includes structures and methods to handle track information
// such as artis, title album, track number, year, genre, and the file path.
package models

const (
	// defaultYear represents the default year for a Rola.
	defaultYear = 2024
	// defaultTrack represents the default track for a Rola.
	defaultTrack = 0
	// defaultID represents the default database ID for a Rola.
	defaultID = 0
)

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
// to default values (track = 0, year = 2024, id = 0).
func NewRola() *Rola {
	defaultValue := "Unknown"
	return &Rola{
		artist: defaultValue,
		title:  defaultValue,
		album:  defaultValue,
		track:  defaultTrack,
		year:   defaultYear,
		genre:  defaultValue,
		path:   defaultValue,
		id:     defaultID,
	}
}

// GetArtist returns the performer of the rola.
func (rola *Rola) GetArtist() string {
	return rola.artist
}

// SetArtist sets the Rola performer.
func (rola *Rola) SetArtist(artist string) {
	rola.artist = artist
}

// GetTitle returns the title of the Rola.
func (rola *Rola) GetTitle() string {
	return rola.title
}

// SetTittle sets the title of the Rola
func (rola *Rola) SetTitle(title string) {
	rola.title = title
}

// GetAlbum returns the album where the Rola is included.
func (rola *Rola) GetAlbum() string {
	return rola.album
}

// SetAlbum sets the album title of the Rola.
func (rola *Rola) SetAlbum(album string) {
	rola.album = album
}

// GetTrack returns the track number of the Rola.
func (rola *Rola) GetTrack() int {
	return rola.track
}

// SetTrack sets the track number of the Rola.
func (rola *Rola) SetTrack(track int) {
	rola.track = track
}

// GetYear returns the year of the Rola.
func (rola *Rola) GetYear() int {
	return rola.year
}

// SetYear sets the year of the Rola.
func (rola *Rola) SetYear(year int) {
	rola.year = year
}

// GetGenre returns the genre of the Rola.
func (rola *Rola) GetGenre() string {
	return rola.genre
}

// SetGenre sets the genre of the Rola.
func (rola *Rola) SetGenre(genre string) {
	rola.genre = genre
}

// GetPath returns the file path of the song where the Rola is located.
func (rola *Rola) GetPath() string {
	return rola.path
}

// SetPath sets the file path of the Rola.
func (rola *Rola) SetPath(path string) {
	rola.path = path
}

// GetID returns the database ID assigned to the Rola.
func (rola *Rola) GetID() int64 {
	return rola.id
}

// SetID sets the database ID of the Rola.
func (rola *Rola) SetID(id int64) {
	rola.id = id
}
