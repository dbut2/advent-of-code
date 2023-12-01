package main

import (
	"embed"
	_ "embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 0)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0

	for _, line := range s {
		firstDigit := 0
		firstSet := false

		lastDigit := 0

		for _, char := range line {
			if char < '0' {
				continue
			}

			if char > '9' {
				continue
			}

			dig := char - '0'

			if !firstSet {
				firstDigit = int(dig)
				firstSet = true
			}

			lastDigit = int(dig)
		}

		total += (firstDigit * 10) + lastDigit
	}

	return total
}
