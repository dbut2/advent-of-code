package main

import (
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input []string) int {
	var cols []int
	var ops []rune

	for _, op := range input[len(input)-1] {
		if op == ' ' {
			continue
		}
		// add each operation to ops
		ops = append(ops, op)
		// add appropriate started to cols
		switch op {
		case '*':
			cols = append(cols, 1)
		case '+':
			cols = append(cols, 0)
		}
	}
	// lets ops match rotated blocks index
	slices.Reverse(ops)

	// rotated input, reading top to bottom, right to left
	rotated := rotate(input[:len(input)-1])

	i := 0
	for _, line := range rotated {
		// empty line starts next block
		if strings.TrimSpace(line) == "" {
			i++
			continue
		}

		// should be exactly 1 number if not an empty line
		num := Ints(line)[0]
		switch ops[i] {
		case '*':
			cols[i] *= num
		case '+':
			cols[i] += num
		}
	}

	return Sum(cols...)
}

func rotate(in []string) []string {
	longest := 0
	for _, line := range in {
		longest = max(longest, len(line))
	}

	out := make([]string, longest)
	for _, line := range in {
		for j, char := range line {
			out[longest-j-1] += string(char)
		}
	}
	return out
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 3263827)
	h.Run()
}
