package main

import (
	"embed"
	_ "embed"
	"regexp"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

func solve(input string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	total := 0
	for _, match := range matches {
		total += sti.Sti(match[1]) * sti.Sti(match[2])
	}

	return total
}

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 161)
	h.Run()
}

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS
