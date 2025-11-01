package main

// fitness scores a path according to:
// - reach goal fast
// - avoid loops
// - if not reached, get as close as possible
// - invalid move => huge penalty
func fitness(p Path) int {
	x, y := startPoint[0], startPoint[1]

	visited := map[[2]int]bool{
		{x, y}: true,
	}
	steps := 0
	loopPenalty := 0

	for _, mv := range p {
		d := moveDelta[mv]
		nx, ny := x+d[0], y+d[1]

		// invalid -> terrible score
		if !inBounds(nx, ny) {
			return -10000
		}
		if maze[nx][ny] == wallValue {
			return -10000
		}

		steps++
		x, y = nx, ny

		pos := [2]int{x, y}
		if visited[pos] {
			loopPenalty++
		} else {
			visited[pos] = true
		}

		// reached goal
		if x == endPoint[0] && y == endPoint[1] {
			return 1000000 - steps*1000 - loopPenalty*5
		}
	}

	// didn't reach goal
	dist := manhattan([2]int{x, y}, endPoint)
	score := 10000 - dist*500 - steps*10 - loopPenalty*50
	return score
}
