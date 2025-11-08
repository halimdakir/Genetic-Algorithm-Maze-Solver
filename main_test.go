package main

import "testing"

// smoke test: load maze and make sure GA returns something non-empty
func TestGAProducesPath(t *testing.T) {
	data, err := LoadMaze("json/maze.json")
	if err != nil {
		t.Fatalf("failed to load maze: %v", err)
	}
	maze = data.Maze
	startPoint = data.Start
	endPoint = data.End

	best := geneticAlgorithm()
	if len(best) == 0 {
		t.Fatal("GA returned empty path")
	}
}
