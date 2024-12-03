package main

import (
	"embed"
	_ "embed"

	"github.com/dbut2/advent-of-code/pkg/chars"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 142)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0
	for _, line := range s {
		firstDigit := -1
		lastDigit := 0

		for _, char := range line {
			if !chars.IsNum(char) {
				continue
			}

			if firstDigit == -1 {
				firstDigit = chars.NumVal(char)
			}
			lastDigit = chars.NumVal(char)
		}
		total += (firstDigit * 10) + lastDigit
	}
	return total
}
