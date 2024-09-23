package model

import (
	"testing"
)

// Test for creating a new rola with default values
func TestNewRola(t *testing.T) {
	rola := NewRola()
	expecting := "Unknown"
	assertEqual(t, rola.Artist(), expecting, "artist")
	assertEqual(t, rola.Title(), expecting, "title")
	assertEqual(t, rola.Album(), expecting, "album")
	assertEqual(t, rola.Track(), 0, "track")
	assertEqual(t, rola.Year(), 2024, "year")
	assertEqual(t, rola.Genre(), expecting, "genre")
	assertEqual(t, rola.Path(), expecting, "path")
}

// Auxiliar function to compare expected and received values
func assertEqual(t *testing.T, received, expected interface{}, fieldName string) {
	if received != expected {
		t.Errorf("expecting %v for field %s, received %v", expected, fieldName, received)
	}
}
