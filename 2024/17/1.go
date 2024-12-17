package main

import (
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input [][]int) string {
	a, b, c := input[0][0], input[1][0], input[2][0]
	prog := input[4]

	output := []string{}

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
			output = append(output, fmt.Sprint(combo(prog[i+1])&7))
		case 6:
			b = a >> combo(prog[i+1])
		case 7:
			c = a >> combo(prog[i+1])
		}

		i += 2
	}

	return strings.Join(output, ",")
}

func main() {
	h := harness.New(solve)
	h.Expect(1, "4,6,3,5,6,3,5,2,1,0")
	h.Run()
}
