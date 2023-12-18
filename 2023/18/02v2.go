package main

import (
	"embed"
	"regexp"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 952408144115)
	h.Solve()
}

func hexToDec(hex string) int {
	dec := 0
	for _, char := range hex {
		dec *= 16
		if char >= '0' && char <= '9' {
			dec += int(char - '0')
		}
		if char >= 'a' && char <= 'f' {
			dec += int(char-'a') + 10
		}
	}
	return dec
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var coords [][2]int
	coord := [2]int{0, 0}
	sum := 0
	r, err := regexp.Compile("[0-9a-f]{6}")
	if err != nil {
		panic(err.Error())
	}
	for _, line := range s {
		hex := r.FindString(line)
		amount := hexToDec(hex[:5])
		switch hex[5] {
		case '0':
			coord[1] += amount
		case '1':
			coord[0] += amount
		case '2':
			coord[1] -= amount
		case '3':
			coord[0] -= amount
		}
		coords = append(coords, coord)
		sum += amount
	}

	last := coords[len(coords)-1]
	for _, coord := range coords {
		sum += coord[0]*last[1] - coord[1]*last[0]
		last = coord
	}
	return sum/2 + 1
}
