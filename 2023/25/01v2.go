package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/graphs"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 54)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	connections := map[string][]string{}
	pairs := [][2]string{}

	g := graphs.New[string]()

	for j, line := range s {
		_, _ = j, line

		line = strings.ReplaceAll(line, ":", "")
		splits := strings.Split(line, " ")

		in := splits[0]
		out := splits[1:]

		connections[in] = append(connections[in], out...)

		for _, o := range out {
			g.Connect(in, o, 1)
			g.Connect(o, in, 1)
			pairs = append(pairs, [2]string{in, o})
		}
	}

	// Initialization:
	//Start with the original graph.
	//Initialize an empty set A to keep track of the nodes in the first partition.
	A := sets.Set[string]{}
	a := graphs.New[string]()

	//Main Loop:
	//Repeat the following steps until only two nodes remain in the graph:
	//
	//Node Selection:
	//Choose any node v that has not been contracted before. This node will be merged into the other partition.
	//
	//Contract Node:
	//Merge node v with the partition represented by set A. Add the edges of node v to the cut.
	//Update the edge weights: For each edge (v, w) where w is in partition A, update the weight to be the sum of the weights of edges (v, w) and (w, v).
	//Remove node v and its incident edges from the graph.

	for {
		var n string
		for node := range g.Nodes {
			if !A.Contains(node) {
				n = node
				break
			}
		}

		A.Add(n)
		edges := g.Remove(n)
		for edge, distance := range edges {
			if A.Contains(edge) {
				a.Connect(edge, n, distance)
			}
		}
	}

	//Calculate Min-Cut Value:
	//After all nodes are contracted, the minimum cut value is the sum of the weights of the edges that are part of the cut.
	//
	//Repeat:
	//Repeat the main loop for different choices of starting nodes.
	//Keep track of the minimum cut value found across all iterations.
	//
	//Return Minimum Cut:
	//Once all iterations are completed, return the minimum cut value found. The nodes that belong to the smaller partition will constitute one side of the cut.
}

func group(connections map[string]map[string]bool) [][]string {
	seen := sets.Set[string]{}

	var groups []sets.Set[string]

	stack := lists.Stack[string]{}
	for in := range connections {
		stack.Push(in)
	}
	for len(stack) > 0 {
		in := stack.Pop()

		if seen.Contains(in) {
			continue
		}
		seen.Add(in)

		var g sets.Set[string]
		for _, gn := range groups {
			if gn.Contains(in) {
				g = gn
			}
		}

		if g == nil {
			g = sets.Set[string]{}
			groups = append(groups, g)
		}

		for o, active := range connections[in] {
			if active {
				g.Add(o)
				stack.Push(o)
			}
		}

		for in2, out := range connections {
			for o, active := range out {
				if active && o == in {
					g.Add(in2)
					stack.Push(in2)
				}
			}
		}
	}

	pools := [][]string{}
	for _, s := range groups {
		pools = append(pools, s.Slice())
	}
	return pools
}
