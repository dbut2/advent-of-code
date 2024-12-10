package main

import (
	"embed"
	_ "embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(2, 2286)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0
	var bufferCount int
	var j int
	var line string

	var redCount, greenCount, blueCount int

	for _, line = range s {
		redCount, greenCount, blueCount = 0, 0, 0

		for j = range line {
			if line[j] >= '0' && line[j] <= '9' {
				bufferCount *= 10
				bufferCount += int(line[j] - '0')
				continue
			}
			if line[j] == ' ' {
				continue
			}
			switch line[j] {
			case 'r':
				redCount = max(redCount, bufferCount)
			case 'g':
				greenCount = max(greenCount, bufferCount)
			case 'b':
				blueCount = max(blueCount, bufferCount)
			}

			bufferCount = 0
		}

		total += redCount * greenCount * blueCount
	}

	return total
}
