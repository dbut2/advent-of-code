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
	h.Expect(1, 8)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0
	var bufferCount int
	var i, j int
	var line string

	for i, line = range s {
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
				if bufferCount > 12 {
					goto exit
				}
			case 'g':
				if bufferCount > 13 {
					goto exit
				}
			case 'b':
				if bufferCount > 14 {
					goto exit
				}
			}

			bufferCount = 0
		}
		total += i + 1
	exit:
	}

	return total
}
