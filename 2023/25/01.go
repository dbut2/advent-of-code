package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/graphs"
	"github.com/dbut2/advent-of-code/pkg/harness"
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

	graph := graphs.New[string]()

	for _, line := range s {
		line = strings.ReplaceAll(line, ":", "")
		splits := strings.Split(line, " ")
		in := splits[0]
		out := splits[1:]
		for _, o := range out {
			graph.Connect(in, o, 1)
		}
	}

	for {
		cutEdges, nodePools := graph.MinCut()
		if len(cutEdges) == 3 {
			return len(nodePools[0]) * len(nodePools[1])
		}
	}
}
