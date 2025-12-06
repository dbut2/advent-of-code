package main

import (
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

	// read ints from each line and process per operation
	// pray that correct number of ints on each line
	for _, line := range input[:len(input)-1] {
		for i, v := range Ints(line) {
			switch ops[i] {
			case '*':
				cols[i] *= v
			case '+':
				cols[i] += v
			}
		}
	}

	return Sum(cols...)
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 4277556)
	h.Run()
}
