package main

import (
	"embed"
	"regexp"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

func solve(input string) int {
	r := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	matches := r.FindAllStringSubmatch(input, -1)

	enabled := true
	total := 0
	for _, match := range matches {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				total += sti.Sti(match[2]) * sti.Sti(match[3])
			}
		}
	}

	return total
}

func main() {
	h := harness.New(solve, input, tests, harness.Nothing())
	h.Tester.Expect(2, 48)
	h.Run()
}

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS
