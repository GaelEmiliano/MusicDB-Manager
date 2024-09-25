package models

import (
	"testing"
)

// Test for creating a new song with default values.
func TestNewSong(t *testing.T) {
	song := NewSong()
	expecting := "Unknown"
	assertEqual(t, song.GetArtist(), expecting, "artist")
	assertEqual(t, song.GetTitle(), expecting, "title")
	assertEqual(t, song.GetAlbum(), expecting, "album")
	assertEqual(t, song.GetTrack(), 0, "track")
	assertEqual(t, song.GetYear(), 2024, "year")
	assertEqual(t, song.GetGenre(), expecting, "genre")
	assertEqual(t, song.GetPath(), expecting, "path")
}

// Test for all getters methods in Song.
func TestGetters(t *testing.T) {
	song := NewSong()
	expecting := "Unknown"
	assertEqual(t, song.GetArtist(), expecting, "artist")
	assertEqual(t, song.GetTitle(), expecting, "title")
	assertEqual(t, song.GetAlbum(), expecting, "album")
	assertEqual(t, song.GetTrack(), 0, "track")
	assertEqual(t, song.GetYear(), 2024, "year")
	assertEqual(t, song.GetGenre(), expecting, "genre")
	assertEqual(t, song.GetPath(), expecting, "path")
}

// Test for all setters methods in Song.
func TestSetters(t *testing.T) {
	song := NewSong()
	artist := "Frank Sinatra"
	title := "Born Free"
	album := "The World We Knew"
	track := 4
	year := 1967
	genre := "Swing"
	path := "Music/Frank Sinatra/The World We Knew/"

	song.SetArtist(artist)
	song.SetTitle(title)
	song.SetAlbum(album)
	song.SetTrack(track)
	song.SetYear(year)
	song.SetGenre(genre)
	song.SetPath(path)

	assertEqual(t, song.GetArtist(), artist, "artist")
	assertEqual(t, song.GetTitle(), title, "title")
	assertEqual(t, song.GetAlbum(), album, "album")
	assertEqual(t, song.GetTrack(), track, "track")
	assertEqual(t, song.GetYear(), year, "year")
	assertEqual(t, song.GetGenre(), genre, "genre")
	assertEqual(t, song.GetPath(), path, "path")
}

// Auxiliar function to compare expected and received values.
func assertEqual(t *testing.T, received, expected interface{}, fieldName string) {
	if received != expected {
		t.Errorf("expecting %v for field %s, received %v", expected, fieldName, received)
	}
}
