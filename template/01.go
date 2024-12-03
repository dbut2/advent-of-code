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
	h := harness.New(solve, inputs)
	h.Expect(1, 0)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
