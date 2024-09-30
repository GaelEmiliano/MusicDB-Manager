package models

import "testing"

// pathsDiscovered has the number of total mp3
const pathsDiscovered = 1

// TestScanMusicDirectory tests if the ScanMusicDirectory function correctly
// scans the user's Music directory and identifies mp3 files.
func TestScanMusicDirectory(t *testing.T) {
	// Initializes a new Miner instance
	miner := NewMiner()

	// Call Traverse to populate paths
	err := miner.ScanMusicDirectory()
	if err != nil {
		t.Fatalf("ScanMusicDirectory returned an error: %v", err)
	}

	// Verify that the first path is correct
	expectedFirstPath := "/home/gael/Music/TestProject2/full.mp3"
	if miner.paths[0] != expectedFirstPath {
		t.Errorf("expecting %v, received %v", expectedFirstPath, miner.paths[0])
	}

	// Verify the number of paths discovered
	expectedPathCount := pathsDiscovered
	if len(miner.paths) != expectedPathCount {
		t.Errorf("expecting %v paths, received %v paths", expectedPathCount, len(miner.paths))
	}
}

// TestExtractMetadata tests if the ExtractMetadata function processes
// the mp3 files found in the paths and extracts their metadata correctly.
func TestExtractMetadata(t *testing.T) {
	miner := NewMiner()

	// Populate paths
	err := miner.ScanMusicDirectory()
	if err != nil {
		t.Fatalf("ScanMusicDirectory returned an error: %v", err)
	}

	// Call ExtractMetadata to process metadata
	go miner.ExtractMetadata()

	// Collect paths from the ore channel
	extractedPaths := make([]string, 0)
	for song := range miner.ore {
		extractedPaths = append(extractedPaths, song.GetPath())
	}

	// Ensure all paths were extracted correctly
	if len(extractedPaths) != len(miner.paths) {
		t.Errorf("expected %d paths, but got %d", len(miner.paths), len(extractedPaths))
	}

	// Compare the extracted paths with the original paths
	for i, path := range extractedPaths {
		if path != miner.paths[i] {
			t.Errorf("expecting %v, received %v", miner.paths[i], path)
		}
	}
}
