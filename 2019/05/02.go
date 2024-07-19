package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(3, 3)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	for _, line := range s {
		p := sti.Stis(strings.Split(line, ","))
		return compute(p)
	}

	return -1
}

func compute(program []int) int {
	in := 5
	i := 0
	for {
		ins := program[i] % 100
		modes := program[i] / 100
		switch ins {
		case 1:
			a := program[i+1]
			if modes/1&1 == 0 {
				a = program[a]
			}
			b := program[i+2]
			if modes/10&1 == 0 {
				b = program[b]
			}
			c := program[i+3]

			program[c] = a + b
			i += 4
		case 2:
			a := program[i+1]
			if modes/1&1 == 0 {
				a = program[a]
			}
			b := program[i+2]
			if modes/10&1 == 0 {
				b = program[b]
			}
			c := program[i+3]

			program[c] = a * b
			i += 4
		case 3:
			program[program[i+1]] = in
			i += 2
		case 4:
			fmt.Println(program[program[i+1]])
			i += 2
		case 5:
			a := program[i+1]
			if modes/1&1 == 0 {
				a = program[a]
			}
			b := program[i+2]
			if modes/10&1 == 0 {
				b = program[b]
			}

			i += 3
			if a != 0 {
				i = b
			}
		case 6:
			a := program[i+1]
			if modes/1&1 == 0 {
				a = program[a]
			}
			b := program[i+2]
			if modes/10&1 == 0 {
				b = program[b]
			}

			i += 3
			if a == 0 {
				i = b
			}
		case 7:
			a := program[i+1]
			if modes/1&1 == 0 {
				a = program[a]
			}
			b := program[i+2]
			if modes/10&1 == 0 {
				b = program[b]
			}
			c := program[i+3]

			if a < b {
				program[c] = 1
			} else {
				program[c] = 0
			}

			i += 4
		case 8:
			a := program[i+1]
			if modes/1&1 == 0 {
				a = program[a]
			}
			b := program[i+2]
			if modes/10&1 == 0 {
				b = program[b]
			}
			c := program[i+3]

			if a == b {
				program[c] = 1
			} else {
				program[c] = 0
			}

			i += 4
		case 99:
			return program[0]
		}
	}
}
