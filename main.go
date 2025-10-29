package main

import (
	"fmt"
	"math/rand"
	"time"
)

var maze [][]int
var startPoint [2]int
var endPoint [2]int

const (
	wallValue      = 1000
	populationSize = 750
	generations    = 1200
	mutationRate   = 0.1
	pathLength     = 100
)

var moveDelta = map[byte][2]int{
	'U': {-1, 0},
	'D': {1, 0},
	'L': {0, -1},
	'R': {0, 1},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//TODO Create LoadMaze() function ...
	data, err := LoadMaze("json/maze.json")
	if err != nil {
		panic(err)
	}

	maze = data.Maze
	startPoint = data.Start
	endPoint = data.End

	fmt.Printf("Maze loaded: %d rows x %d cols\n", len(maze), len(maze[0]))
	fmt.Printf("Start: %v  End: %v\n", startPoint, endPoint)

	//TODO Create geneticAlgorithm() function ...
	best := geneticAlgorithm()

	fmt.Println("Best path found:", string(best))

	//TODO Create savePathJSON() function ...
	if err := savePathJSON(best, "json/best_path.json"); err != nil {
		panic(err)
	}

	fmt.Println("Saved best_path.json")
}
