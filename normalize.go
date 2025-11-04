package main

// optimizePath removes loops:
// If we return to a cell we've already visited, we cut out that detour.
func optimizePath(p Path, start [2]int) Path {
	pos := start
	visitedIndex := map[[2]int]int{
		pos: 0,
	}
	var optimized Path
	positions := make([][2]int, 1)
	positions[0] = pos

	for _, mv := range p {
		d := moveDelta[mv]
		newPos := [2]int{pos[0] + d[0], pos[1] + d[1]}

		if idx, ok := visitedIndex[newPos]; ok {
			// loop: cut back
			optimized = optimized[:idx]
			positions = positions[:idx+1]
		} else {
			optimized = append(optimized, mv)
			positions = append(positions, newPos)
			visitedIndex[newPos] = len(optimized)
		}
		pos = newPos
	}

	return optimized
}

// normalizePath:
// - walk the path from the start
// - stop if we go out of bounds, hit a wall, or reach the goal
// - then remove useless loops
func normalizePath(p Path, start, end [2]int) Path {
	x, y := start[0], start[1]
	followed := make(Path, 0, len(p))

	for _, mv := range p {
		d := moveDelta[mv]
		nx, ny := x+d[0], y+d[1]

		// stop if invalid
		if !inBounds(nx, ny) {
			break
		}
		if maze[nx][ny] == wallValue {
			break
		}

		followed = append(followed, mv)
		x, y = nx, ny

		// stop if goal reached
		if x == end[0] && y == end[1] {
			break
		}
	}

	return optimizePath(followed, start)
}
