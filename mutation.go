package main

import "math/rand"

func mutate(ind Path) Path {
	// maybe no mutation
	if rand.Float64() > mutationRate || len(ind) == 0 {
		out := make(Path, len(ind))
		copy(out, ind)
		return out
	}

	// pick cut index
	cut := rand.Intn(len(ind))

	// simulate first part to get position at cut
	x, y := startPoint[0], startPoint[1]
	for i := 0; i < cut; i++ {
		mv := ind[i]
		d := moveDelta[mv]
		x += d[0]
		y += d[1]
	}

	newMoves := make(Path, cut)
	copy(newMoves, ind[:cut])

	var last byte
	if len(newMoves) > 0 {
		last = newMoves[len(newMoves)-1]
	}

	for i := cut; i < pathLength; i++ {
		mv, nx, ny, ok := generateValidStep(x, y)
		if !ok {
			break
		}

		// avoid instant undo like U then D
		if last != 0 && ((last == 'U' && mv == 'D') ||
			(last == 'D' && mv == 'U') ||
			(last == 'L' && mv == 'R') ||
			(last == 'R' && mv == 'L')) {
			continue
		}

		newMoves = append(newMoves, mv)
		x, y = nx, ny
		last = mv

		if x == endPoint[0] && y == endPoint[1] {
			break
		}
	}

	return newMoves
}
