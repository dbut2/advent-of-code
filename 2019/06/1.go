package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
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

	count := len(orbits)
	for left := range orbits {
		count += distanceFromRoot(orbits, left) - 1
	}
	return count
}

var c = map[string]int{}

func distanceFromRoot(orbits map[string]string, object string) int {
	if v, ok := c[object]; ok {
		return v
	}

	parent, ok := orbits[object]
	if !ok {
		return 0 // is root
	}

	d := distanceFromRoot(orbits, parent) + 1
	c[object] = d
	return d
}
