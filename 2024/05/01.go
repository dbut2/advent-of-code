package main

import (
	"embed"
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

func solve(input [2][]string) int {
	orderingRules, updates := input[0], input[1]

	// Create a dependancy map of Key depends on []Values
	dep := make(map[string][]string)

	for _, line := range orderingRules {
		parts := strings.Split(line, "|")
		depender, dependant := parts[0], parts[1]
		dep[dependant] = append(dep[dependant], depender)
	}

	total := 0
	for _, line := range updates {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")

		valid := true
		for i, part := range parts {
			for j := i + 1; j < len(parts); j++ {
				// If there exists a later value that the current value depends
				// on, this is invalid as it needs to be to the left.
				if slices.Contains(dep[part], parts[j]) {
					valid = false
				}
			}
		}

		if valid {
			total += sti.Int(parts[len(parts)/2])
		}
	}

	return total
}

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 143)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
