package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 62)
	h.Solve()
}

var deltas = map[rune][2]int{
	'U': {0, -1},
	'D': {0, 1},
	'L': {-1, 0},
	'R': {1, 0},
}

func solve(input string) int {
	length := 0
	sum := 0

	coord := [2]int{0, 0}
	var newCoord [2]int
	var direction [2]int
	var amount int

	for _, char := range input {
		if char == ' ' {
			if amount != 0 {
				newCoord = coord

				newCoord[0] += direction[0] * amount
				newCoord[1] += direction[1] * amount

				sum += newCoord[0]*coord[1] - newCoord[1]*coord[0]
				coord = newCoord
				length += amount
			}
		}

		if char >= 'A' && char <= 'Z' {
			amount = 0
			direction = deltas[char]
		}

		if char >= '0' && char <= '9' {
			amount = amount*10 + int(char-'0')
		}
	}

	if sum < 0 {
		sum = -sum
	}

	return sum/2 + length/2 + 1
}
