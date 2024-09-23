package models

import (
	"testing"
)

// Test for creating a new rola with default values
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

// Test for getters
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

// Auxiliar function to compare expected and received values
func assertEqual(t *testing.T, received, expected interface{}, fieldName string) {
	if received != expected {
		t.Errorf("expecting %v for field %s, received %v", expected, fieldName, received)
	}
}
