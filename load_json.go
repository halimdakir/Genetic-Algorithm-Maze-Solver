package main

import (
	"encoding/json"
	"os"
)

type MazeData struct {
	Maze  [][]int `json:"maze"`
	Start [2]int  `json:"start"`
	End   [2]int  `json:"end"`
}

func LoadMaze(filename string) (MazeData, error) {
	var m MazeData
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return m, err
	}
	if err := json.Unmarshal(bytes, &m); err != nil {
		return m, err
	}
	return m, nil
}
