package models

import (
	"testing"
)

// Test for creating a new rola with default values.
func TestNewRola(t *testing.T) {
	rola := NewRola()
	expecting := "Unknown"
	assertEqual(t, rola.GetArtist(), expecting, "artist")
	assertEqual(t, rola.GetTitle(), expecting, "title")
	assertEqual(t, rola.GetAlbum(), expecting, "album")
	assertEqual(t, rola.GetTrack(), 0, "track")
	assertEqual(t, rola.GetYear(), 2024, "year")
	assertEqual(t, rola.GetGenre(), expecting, "genre")
	assertEqual(t, rola.GetPath(), expecting, "path")
}

// Test for all getters methods in Rola.
func TestGetters(t *testing.T) {
	rola := NewRola()
	expecting := "Unknown"
	assertEqual(t, rola.GetArtist(), expecting, "artist")
	assertEqual(t, rola.GetTitle(), expecting, "title")
	assertEqual(t, rola.GetAlbum(), expecting, "album")
	assertEqual(t, rola.GetTrack(), 0, "track")
	assertEqual(t, rola.GetYear(), 2024, "year")
	assertEqual(t, rola.GetGenre(), expecting, "genre")
	assertEqual(t, rola.GetPath(), expecting, "path")
}

// Test for all setters methods in Rola.
func TestSetters(t *testing.T) {
	rola := NewRola()
	artist := "Frank Sinatra"
	title := "Born Free"
	album := "The World We Knew"
	track := 4
	year := 1967
	genre := "Swing"
	path := "Music/Frank Sinatra/The World We Knew/"

	rola.SetArtist(artist)
	rola.SetTitle(title)
	rola.SetAlbum(album)
	rola.SetTrack(track)
	rola.SetYear(year)
	rola.SetGenre(genre)
	rola.SetPath(path)

	assertEqual(t, rola.GetArtist(), artist, "artist")
	assertEqual(t, rola.GetTitle(), title, "title")
	assertEqual(t, rola.GetAlbum(), album, "album")
	assertEqual(t, rola.GetTrack(), track, "track")
	assertEqual(t, rola.GetYear(), year, "year")
	assertEqual(t, rola.GetGenre(), genre, "genre")
	assertEqual(t, rola.GetPath(), path, "path")
}

// Auxiliar function to compare expected and received values.
func assertEqual(t *testing.T, received, expected interface{}, fieldName string) {
	if received != expected {
		t.Errorf("expecting %v for field %s, received %v", expected, fieldName, received)
	}
}
