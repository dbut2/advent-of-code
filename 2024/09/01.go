package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input string) int {
	var line []int
	for i, char := range input {
		n := -1
		if i%2 == 0 {
			n = i / 2
		}
		for range char - '0' {
			line = append(line, n)
		}
	}

	for i, j := 0, len(line)-1; i < j; {
		if line[i] != -1 {
			i++
			continue
		}

		if line[j] == -1 {
			j--
			continue
		}

		line[i], line[j] = line[j], line[i]
	}

	total := 0
	for i, c := range line {
		if c == -1 {
			continue
		}
		total += c * i
	}

	return total
}

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 1928)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
