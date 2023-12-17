package main

import (
	"embed"
	"math"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 102)
	h.Solve()
}

type node struct {
	i, j       int
	value      int
	variations map[[3]*node]int
}

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := space.NewGrid[node](len(s[0]), len(s))

	for j, line := range s {
		_, _ = j, line

		for i, char := range line {
			grid[i][j] = node{
				i:          i,
				j:          j,
				value:      int(char - '0'),
				variations: map[[3]*node]int{},
			}
		}
	}

	grid[0][0].variations[[3]*node{}] = 0

	nodeQueue := lists.Queue[*node]{}
	nodeQueue.Push(&grid[1][0])
	nodeQueue.Push(&grid[0][1])

	for len(nodeQueue) > 0 {
		n := nodeQueue.Pop()

		seenNew := false

		for _, neighbour := range grid.Adjacent(n.i, n.j) {
			for last3Moves, minScore := range neighbour.variations {
				if last3Moves[0] == n {
					continue
				}
				if inARow(n, neighbour, last3Moves[0], last3Moves[1], last3Moves[2]) {
					continue
				}

				score := minScore + n.value

				key := [3]*node{neighbour, last3Moves[0], last3Moves[1]}

				if oldScore, ok := n.variations[key]; ok {
					if score < oldScore {
						seenNew = true
					}
					score = min(score, oldScore)
				} else {
					seenNew = true
				}
				n.variations[key] = score
			}
		}

		if seenNew {
			for _, neighbour := range grid.Adjacent(n.i, n.j) {
				nodeQueue.Push(neighbour)
			}
		}
	}

	vs := grid[len(grid)-1][len(grid[0])-1].variations
	m := math.MaxInt
	for _, v := range vs {
		m = min(m, v)
	}
	return m
}

func inARow(nodes ...*node) bool {
	is := sets.Set[int]{}
	js := sets.Set[int]{}

	for _, node := range nodes {
		if node == nil {
			return false
		}

		is.Add(node.i)
		js.Add(node.j)
	}

	return len(is) == 1 || len(js) == 1
}
