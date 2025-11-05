package main

import "math/rand"

// Path is a sequence of moves like []byte{'R','R','D','L',...}
type Path []byte

// inBounds checks if (x,y) is inside the maze grid.
func inBounds(x, y int) bool {
	return x >= 0 && x < len(maze) && y >= 0 && y < len(maze[0])
}

// manhattan distance between two grid cells
func manhattan(a, b [2]int) int {
	dx := a[0] - b[0]
	if dx < 0 {
		dx = -dx
	}
	dy := a[1] - b[1]
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

// generateValidStep(x,y):
// returns (move, nx, ny, ok)
// where move is 'U','D','L','R'
func generateValidStep(x, y int) (byte, int, int, bool) {
	type candidate struct {
		mv byte
		nx int
		ny int
	}
	opts := []candidate{}

	// up
	if x > 0 && maze[x-1][y] == 1 {
		opts = append(opts, candidate{'U', x - 1, y})
	}
	// down
	if x < len(maze)-1 && maze[x+1][y] == 1 {
		opts = append(opts, candidate{'D', x + 1, y})
	}
	// left
	if y > 0 && maze[x][y-1] == 1 {
		opts = append(opts, candidate{'L', x, y - 1})
	}
	// right
	if y < len(maze[0])-1 && maze[x][y+1] == 1 {
		opts = append(opts, candidate{'R', x, y + 1})
	}

	if len(opts) == 0 {
		return 0, 0, 0, false
	}
	choice := opts[rand.Intn(len(opts))]
	return choice.mv, choice.nx, choice.ny, true
}
