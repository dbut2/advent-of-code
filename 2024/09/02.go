package main

import (
	"embed"
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input string) int {
	var line [][2]int
	for i, char := range input {
		n := -1
		if i%2 == 0 {
			n = i / 2
		}
		line = append(line, [2]int{n, int(char - '0')})
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i][0] == -1 {
			continue
		}
		for j := 0; j < i; j++ {
			if line[j][0] != -1 {
				continue
			}
			if line[i][1] > line[j][1] {
				continue
			}

			line[i], line[j] = line[j], line[i]
			line = slices.Insert(line, j+1, [2]int{-1, line[i][1] - line[j][1]})
			line[i+1][1] = line[j][1]
			i++
			break
		}
	}

	total := 0
	c := -1
	for i := range line {
		for range line[i][1] {
			c++
			if line[i][0] == -1 {
				continue
			}
			total += line[i][0] * c
		}
	}
	return total
}

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 2858)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
