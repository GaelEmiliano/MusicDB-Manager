package models

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/dhowden/tag"
)

const (
	// invalidArtist is used to represent an empty artist
	// string when metadata is missing.
	invalidArtist = ""
	// invalidTitle is used to represent an empty title
	// string when metadata is missing.
	invalidTitle = ""
	// invalidAlbum is used to represent an empty album
	// string when metadata is missing.
	invalidAlbum = ""
	// invalidGenre is used to represent an empty genre
	// string when metadata is missing.
	invalidGenre = ""
	// invalidTrack is used to represent the default track
	// number when metadata is missing.
	invalidTrack = 0
	// invalidYear is used to represent the default year
	// number when metadata is missing.
	invalidYear = 0
)

// Miner is the main structure that manages file paths, channels for
// extracted songs, and error handling.
type Miner struct {
	paths  []string   // List of file paths for the mp3 files
	ore    chan *Song // Channel to send extracted Song objects
	errors chan error // channel for error handling
}

// NewMiner initializes and returns a new Miner instance with all channels
// and paths prepared.
func NewMiner() *Miner {
	return &Miner{
		paths:  make([]string, 0),
		ore:    make(chan *Song),
		errors: make(chan error),
	}
}

// ScanMusicDirectory walks through the user's Music directory and identifies mp3
// files, adding their paths to the paths list in the Miner struct.
func (miner *Miner) ScanMusicDirectory() error {
	// Get the current user's information
	home, err := user.Current()
	if err != nil {
		// Return an error if the current user cannot be retrieved
		return fmt.Errorf("could not retrieve the current user: %w", err)
	}
	// Set the start directory to the user's "Music" folder
	start := filepath.Join(home.HomeDir, "Music")

	// Walk through the files and directories inside the music folder
	err = filepath.Walk(start, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			miner.errors <- fmt.Errorf("failure accesing path %s: %w", path, err)
			return nil // Continue walking the path despite the error
		}
		// If the file is not a directory and has an .mp3 extension, add it to
		// the path list
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
			miner.paths = append(miner.paths, path)
		}
		return nil
	})
	// If an error occurs during traversal, return it
	if err != nil {
		return fmt.Errorf("error walking the path: %w", err)
	}
	return nil // Return nil if there were no errors
}

// ExtractMetadata processes the mp3 files found during traversal and extracts
// metadata, sending Song objects to the ore channel.
func (miner *Miner) ExtractMetadata() {
	// Obtain a genre converter to handle genre normalization
	genreConverter := ObtainGenre()

	// Iterate over all the paths collected in the miner
	for _, path := range miner.paths {
		// Open the file located at the given path
		file, err := os.Open(path)
		if err != nil {
			// If the file cannot be opened, send an error message to the errors channel
			miner.errors <- fmt.Errorf("could not open file %s: %w", path, err)
			// Skip further processing of this file and move to the next one
			continue
		}

		// Attempt to read metadata from the opened file using the tag library
		metadata, err := tag.ReadFrom(file)
		if err != nil {
			// If metadata cannot be read, send an error message to the errors channel
			miner.errors <- fmt.Errorf("could not read the tag from file %s: %w", path, err)
			// Close the file and continue to the next one
			file.Close()
			continue
		}

		// Close the file as it's no longer needed after reading metadata
		file.Close()

		// Create a new Song object to store the extracted metadata
		song := NewSong()
		if metadata.Artist() != invalidArtist {
			song.SetArtist(metadata.Artist())
		}
		if metadata.Title() != invalidTitle {
			song.SetTitle(metadata.Title())
		}
		if metadata.Album() != invalidAlbum {
			song.SetAlbum(metadata.Album())
		}
		track, _ := metadata.Track()
		if track != invalidTrack {
			song.SetTrack(track)
		}
		if metadata.Year() != invalidYear {
			song.SetYear(metadata.Year())
		}
		if metadata.Genre() != invalidGenre {
			song.SetGenre(genreConverter.Get(metadata.Genre()))
		}
		song.SetPath(path)

		// Send the fully constructed Song object to the ore channel
		miner.ore <- song
	}

	// After processing all files, close the ore and errors channels to
	// signal that no more values will be send
	close(miner.ore)
	close(miner.errors)
}
