package main

import (
	"embed"
	"slices"
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
	for i := 0; i < len(program); i++ {
		for j := 0; j < len(program); j++ {
			if compute(slices.Clone(program), i, j) == 19690720 {
				return i*100 + j
			}
		}
	}
	return 0
}

func compute(program []int, noun, verb int) int {
	program[1] = noun
	program[2] = verb
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
