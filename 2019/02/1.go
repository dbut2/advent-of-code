package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
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
	program := sti.Stis(strings.Split(strings.TrimSpace(input), ","))
	program[1] = 12
	program[2] = 2
	i := 0
	for {
		switch program[i] {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
			i += 4
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
			i += 4
		case 99:
			return program[0]
		}
	}
}
