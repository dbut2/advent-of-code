package main

import (
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
)

func solve(input []string) int {
	connections := map[string][]string{}
	for _, line := range input {
		parts := strings.Split(line, "-")
		connections[parts[0]] = append(connections[parts[0]], parts[1])
		connections[parts[1]] = append(connections[parts[1]], parts[0])
	}

	thruples := sets.Set[[3]string]{}
	for from, tos := range connections {
		for _, to := range tos {
			thirds := lists.Intersection(connections[from], connections[to])
			for _, third := range thirds {
				key := []string{from, to, third}
				slices.Sort(key)
				thruples.Add([3]string(key))
			}
		}
	}

	total := 0
	for thruple := range thruples.Seq {
		for _, t := range thruple {
			if t[0] == 't' {
				total++
				break
			}
		}
	}
	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 7)
	h.Run()
}
