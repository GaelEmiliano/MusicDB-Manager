package models

const (
	// defaultYear represents the default year for a Song.
	defaultYear = 2024
	// defaultTrack represents the default track for a Song.
	defaultTrack = 0
	// defaultID represents the default database ID for a Song.
	defaultID = 0
)

// Song represents a song with various attributes extracted from ID3v2.3 tags.
type Song struct {
	artist string // Artist name (performer of the song)
	title  string // Tittle of the song
	album  string // Album name where the song belongs
	track  int    // Track number in the album
	year   int    // Year when the song was released
	genre  string // Genre of the song (e.g. Pop, Rock)
	path   string // File path of the mp3 file
	id     int64  // ID in the database (used for stored the song)
}

// NewSong creates and returns a new Song with default values.
// The text fields (artist, title, album, genre, path) are initialized
// to "Unknow", and numeric fields (track, year, id) are initialized
// to default values (track = 0, year = 2024, id = 0).
func NewSong() *Song {
	defaultValue := "Unknown"
	return &Song{
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

// GetArtist returns the performer of the song.
func (song *Song) GetArtist() string {
	return song.artist
}

// SetArtist sets the Song performer.
func (song *Song) SetArtist(artist string) {
	song.artist = artist
}

// GetTitle returns the title of the Song.
func (song *Song) GetTitle() string {
	return song.title
}

// SetTittle sets the title of the Song
func (song *Song) SetTitle(title string) {
	song.title = title
}

// GetAlbum returns the album where the Song is included.
func (song *Song) GetAlbum() string {
	return song.album
}

// SetAlbum sets the album title of the Song.
func (song *Song) SetAlbum(album string) {
	song.album = album
}

// GetTrack returns the track number of the Song.
func (song *Song) GetTrack() int {
	return song.track
}

// SetTrack sets the track number of the Song.
func (song *Song) SetTrack(track int) {
	song.track = track
}

// GetYear returns the year of the Song.
func (song *Song) GetYear() int {
	return song.year
}

// SetYear sets the year of the Song.
func (song *Song) SetYear(year int) {
	song.year = year
}

// GetGenre returns the genre of the Song.
func (song *Song) GetGenre() string {
	return song.genre
}

// SetGenre sets the genre of the Song.
func (song *Song) SetGenre(genre string) {
	song.genre = genre
}

// GetPath returns the file path of the song where the Song is located.
func (song *Song) GetPath() string {
	return song.path
}

// SetPath sets the file path of the Song.
func (song *Song) SetPath(path string) {
	song.path = path
}

// GetID returns the database ID assigned to the Song.
func (song *Song) GetID() int64 {
	return song.id
}

// SetID sets the database ID of the Song.
func (song *Song) SetID(id int64) {
	song.id = id
}
