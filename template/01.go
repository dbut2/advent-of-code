package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input []string) int {
	
}

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 0)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
