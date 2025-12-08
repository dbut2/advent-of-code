package main

import (
	"math"
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/strings"
)

func solve(input []string) int {
	// Process input as list of positions
	var poss []pos
	for _, line := range input {
		poss = append(poss, pos(strings.Ints(line)))
	}

	// Find all pairs of points and order a list on their distance
	pairs := [][2]pos{}
	distances := map[[2]pos]float64{}
	for i := 0; i < len(poss); i++ {
		for j := 0; j < i; j++ {
			pair := [2]pos{poss[i], poss[j]}
			distances[pair] = distance(poss[i], poss[j])
			pairs = append(pairs, pair)
		}
	}
	slices.SortFunc(pairs, func(a, b [2]pos) int {
		return int(distance(a[0], a[1]) - distance(b[0], b[1]))
	})

	// Continually connect the closest points until there is only a single group
	connections := make(map[pos][]pos)
	for _, pair := range pairs {
		connections[pair[0]] = append(connections[pair[0]], pair[1])
		connections[pair[1]] = append(connections[pair[1]], pair[0])

		if len(groups(poss, connections)) == 1 {
			return pair[0][0] * pair[1][0]
		}
	}

	panic("all processed pairs must result in a single group")
}

type pos [3]int

// cartesian distance of 2 3D points
func distance(a, b pos) float64 {
	return math.Sqrt(math.Pow(float64(a[0]-b[0]), 2) + math.Pow(float64(a[1]-b[1]), 2) + math.Pow(float64(a[2]-b[2]), 2))
}

// groups returns the group sizes for a disconnected graph of 3D points
func groups(poss []pos, connections map[pos][]pos) []int {
	posSet := sets.SetFrom(poss)
	var groupSizes []int
	seen := sets.Set[pos]{}

	// While we still have unprocessed positions, create a new groupSizes with
	// all connected points from some selected point
	for len(posSet) > 0 {
		groupSize := 0
		q := lists.NewQueue[pos]()
		q.Push(posSet.Slice()[0])
		// Queue and process all connected nodes starting from posSet[0]
		for pos := range q.Seq {
			if seen.Contains(pos) {
				continue
			}
			seen.Add(pos)

			groupSize++
			posSet.Remove(pos)
			for _, next := range connections[pos] {
				q.Push(next)
			}
		}
		groupSizes = append(groupSizes, groupSize)
	}
	return groupSizes
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 25272)
	h.Run()
}
