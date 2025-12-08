package main

import (
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input [][]int) int {
	type Pos [3]int

	// Calculate squared Euclidean distance
	// https://en.wikipedia.org/wiki/Euclidean_distance
	distance := func(a, b Pos) int {
		dx := a[0] - b[0]
		dy := a[1] - b[1]
		dz := a[2] - b[2]
		return dx*dx + dy*dy + dz*dz
	}

	// Convert input to positions
	var positions []Pos
	for _, line := range input {
		positions = append(positions, Pos(line))
	}

	// Generate all pairs of positions and sort by distance
	var pairs [][2]Pos
	for i := 1; i < len(positions); i++ {
		for j := 0; j < i; j++ {
			pairs = append(pairs, [2]Pos{positions[i], positions[j]})
		}
	}
	slices.SortFunc(pairs, func(a, b [2]Pos) int {
		return distance(a[0], a[1]) - distance(b[0], b[1])
	})

	// Initialize union-find structure
	// graphs holds the nodes in each connected component
	// graphIndex maps each position to its graph index
	graphs := make([][]Pos, len(positions))
	graphIndex := map[Pos]int{}
	for i, pos := range positions {
		graphs[i] = []Pos{pos}
		graphIndex[pos] = i
	}

	// Connect closest points until all positions form a single graph
	for _, pair := range pairs {
		indexA, indexB := graphIndex[pair[0]], graphIndex[pair[1]]

		if indexA == indexB {
			continue
		}

		// Merge graph B into graph A
		for _, node := range graphs[indexB] {
			graphIndex[node] = indexA
		}
		graphs[indexA] = append(graphs[indexA], graphs[indexB]...)

		if len(graphs[indexA]) == len(positions) {
			// Return product of x value of last connecting pair
			return pair[0][0] * pair[1][0]
		}
	}

	panic("all processed pairs must result in a single graph")
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 25272)
	h.Run()
}
