package main

import (
	"embed"
	_ "embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var test embed.FS

func main() {
	h := harness.New(solve, input, test)
	h.Expect(1, 24000)
	h.Solve()
}

func solve(input string) int {
	return lists.Reduce(strings.Split(input, "\n\n"), func(m int, s string) int { return max(m, math.Sum(sti.Stis(strings.Split(s, "\n"))...)) })
}
