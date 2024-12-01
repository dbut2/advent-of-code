package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []string) int {
	for i, line := range input {
		_, _ = i, line

	}
}

func main() {
	h := harness.New(solve, input, tests, harness.SplitNewlines())
	h.Tester.Expect(1, 0)
	h.Run()
}

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS
