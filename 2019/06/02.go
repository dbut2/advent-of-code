package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	orbits := map[string]string{}

	for j, line := range s {
		_, _ = j, line

		sides := strings.Split(line, ")")
		orbits[sides[1]] = sides[0]
	}

	pathA := lists.Stack[string]{"YOU"}
	for {
		curr := pathA.Peek()
		parent, ok := orbits[curr]
		if !ok {
			break
		}
		pathA.Push(parent)
	}

	pathB := lists.Stack[string]{"SAN"}
	for {
		curr := pathB.Peek()
		parent, ok := orbits[curr]
		if !ok {
			break
		}
		pathB.Push(parent)
	}

	for pathA.Pop() == pathB.Pop() {
	}

	return len(pathA) + len(pathB)
}
