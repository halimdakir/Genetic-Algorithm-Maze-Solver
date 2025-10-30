package main

import "math/rand"

func crossover(a, b Path) (Path, Path) {
	if len(a) < 2 || len(b) < 2 {
		c1 := make(Path, len(a))
		copy(c1, a)
		c2 := make(Path, len(b))
		copy(c2, b)
		return c1, c2
	}

	maxCP := len(a)
	if len(b) < maxCP {
		maxCP = len(b)
	}
	maxCP--

	if maxCP <= 1 {
		c1 := make(Path, len(a))
		copy(c1, a)
		c2 := make(Path, len(b))
		copy(c2, b)
		return c1, c2
	}

	cp1 := rand.Intn(maxCP-1) + 1         // [1, maxCP-1]
	cp2 := rand.Intn(maxCP-cp1) + cp1 + 1 // (cp1+1 .. maxCP)

	child1 := append(append(append(Path{}, a[:cp1]...), b[cp1:cp2]...), a[cp2:]...)
	child2 := append(append(append(Path{}, b[:cp1]...), a[cp1:cp2]...), b[cp2:]...)
	return child1, child2
}
