package main

import (
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
)

func solve(input []string) string {

	connections := map[string][]string{}

	for _, line := range input {
		parts := strings.Split(line, "-")
		connections[parts[0]] = append(connections[parts[0]], parts[1])
		connections[parts[1]] = append(connections[parts[1]], parts[0])
	}

	// magic number reasonably guessed
	const k = 20

	groups := sets.Set[[k]string]{}
	for from, tos := range connections {
		for _, to := range tos {
			groups.Add([k]string{from, to})
		}
	}
	for {
		next := sets.Set[[k]string]{}
		for group := range groups.Seq {
			gs := group[:]
			gs = lists.Filter(gs, func(g string) bool { return g != "" })
			allHas := connections[gs[0]]
			for _, g := range gs {
				allHas = lists.Intersection(allHas, connections[g])
			}
			for _, n := range allHas {
				nn := append(slices.Clone(gs), n)
				slices.Sort(nn)
				next.Add([k]string(pad(k, nn)))
			}
		}
		if len(next) == 0 {
			out := groups.Slice()[0][:]
			out = lists.Filter(out, func(g string) bool { return g != "" })
			slices.Sort(out)
			return strings.Join(out, ",")
		}
		groups = next
	}
}

func pad(n int, s []string) []string {
	out := make([]string, n)
	copy(out, s)
	return out
}

func main() {
	h := harness.New(solve)
	h.Expect(1, "co,de,ka,ta")
	h.Run()
}
