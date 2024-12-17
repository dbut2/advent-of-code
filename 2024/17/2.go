package main

import (
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input [][]int) int {
	b, c := input[1][0], input[2][0]
	prog := input[4]

	a := 1
	for {
		out := run(a, b, c, prog)
		if slices.Equal(out, prog) {
			break
		}
		if slices.Equal(out, prog[len(prog)-len(out):]) {
			a *= 8
		} else {
			if a%8 == 7 {
				a /= 8
			}
			a++
		}
	}

	return a
}

func run(a int, b int, c int, prog []int) []int {
	output := []int{}

	combo := func(i int) int {
		switch {
		case i < 4:
			return i
		case i == 4:
			return a
		case i == 5:
			return b
		case i == 6:
			return c
		case i == 7:
			fallthrough
		default:
			panic("invalid combo operand")
		}
	}

	i := 0
	for {
		if i >= len(prog) {
			break
		}
		switch prog[i] {
		case 0:
			a >>= combo(prog[i+1])
		case 1:
			b ^= prog[i+1]
		case 2:
			b = combo(prog[i+1]) & 7
		case 3:
			if a != 0 {
				i = prog[i+1]
				continue
			}
		case 4:
			b ^= c
		case 5:
			output = append(output, combo(prog[i+1])&7)
		case 6:
			b = a >> combo(prog[i+1])
		case 7:
			c = a >> combo(prog[i+1])
		}

		i += 2
	}
	return output
}

func main() {
	h := harness.New(solve)
	h.Expect(2, 117440)
	h.Run()
}
