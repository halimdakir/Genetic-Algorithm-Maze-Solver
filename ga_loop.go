package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// Scored wraps (fitness, path)
type Scored struct {
	fit  int
	path Path
}

// tournamentSelection: pick best from a random mini-pool
func tournamentSelection(pop []Scored, k int) Path {
	bestFit := math.MinInt
	var best Path

	for i := 0; i < k; i++ {
		cand := pop[rand.Intn(len(pop))]
		if cand.fit > bestFit {
			bestFit = cand.fit
			best = cand.path
		}
	}
	return best
}

// generateIndividual = Python generate_individual()
func generateIndividual() Path {
	x, y := startPoint[0], startPoint[1]
	var moves Path
	var last byte

	for i := 0; i < pathLength; i++ {
		mv, nx, ny, ok := generateValidStep(x, y)
		if !ok {
			break
		}

		// avoid immediate backtracking like U->D
		if last != 0 && ((last == 'U' && mv == 'D') ||
			(last == 'D' && mv == 'U') ||
			(last == 'L' && mv == 'R') ||
			(last == 'R' && mv == 'L')) {
			continue
		}

		moves = append(moves, mv)
		x, y = nx, ny
		last = mv

		if x == endPoint[0] && y == endPoint[1] {
			break
		}
	}

	// clean it
	return normalizePath(moves, startPoint, endPoint)
}

// geneticAlgorithm = core evolution loop
func geneticAlgorithm() Path {
	// 1. initial population
	pop := make([]Path, populationSize)
	for i := 0; i < populationSize; i++ {
		pop[i] = generateIndividual()
	}

	bestFit := math.MinInt
	var best Path

	for gen := 1; gen <= generations; gen++ {
		// score all
		scored := make([]Scored, len(pop))
		for i, ind := range pop {
			f := fitness(ind)
			scored[i] = Scored{fit: f, path: ind}

			if f > bestFit {
				bestFit = f
				best = append(Path{}, ind...)
			}
		}

		fmt.Printf("Generation %d: best fitness so far = %d\n", gen, bestFit)

		// sort desc by fitness
		sort.Slice(scored, func(i, j int) bool {
			return scored[i].fit > scored[j].fit
		})

		// elitism: top 10%
		eliteCount := int(0.1 * float64(populationSize))
		if eliteCount < 1 {
			eliteCount = 1
		}

		nextGen := make([]Path, 0, populationSize)
		for i := 0; i < eliteCount; i++ {
			elite := scored[i].path
			elite = normalizePath(elite, startPoint, endPoint)
			nextGen = append(nextGen, elite)
		}

		// fill the rest
		for len(nextGen) < populationSize {
			p1 := tournamentSelection(scored, 5)
			p2 := tournamentSelection(scored, 5)

			// crossover
			c1, c2 := crossover(p1, p2)

			// normalize (legalize them)
			c1 = normalizePath(c1, startPoint, endPoint)
			c2 = normalizePath(c2, startPoint, endPoint)

			// mutate
			c1 = mutate(c1)
			c2 = mutate(c2)

			// normalize again after mutation
			c1 = normalizePath(c1, startPoint, endPoint)
			c2 = normalizePath(c2, startPoint, endPoint)

			nextGen = append(nextGen, c1)
			if len(nextGen) < populationSize {
				nextGen = append(nextGen, c2)
			}
		}

		pop = nextGen
	}

	return best
}
