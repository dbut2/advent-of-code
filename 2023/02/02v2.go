package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(2, 2286)
	fmt.Println(solve(input))
	benchmark.Run(func() {
		solve(input)
	}, benchmark.Count(1000))
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
