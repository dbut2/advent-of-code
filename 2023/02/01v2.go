package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
	"time"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 8)
	fmt.Println(solve(input))
	benchmark.Run(func() {
		solve(input)
	}, benchmark.Time(time.Second))
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
