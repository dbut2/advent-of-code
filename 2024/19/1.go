package main

import (
	"regexp"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []string) int {
	supply := strings.Split(input[0], ", ")
	r := regexp.MustCompile(`^(` + strings.Join(supply, "|") + `)*$`)
	total := 0
	for _, towel := range input[2:] {
		if r.FindString(towel) != "" {
			total++
		}
	}
	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 6)
	h.Run()
}
