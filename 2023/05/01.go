package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 35)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	type submapping struct {
		source, size, offset int
	}
	type mapping []submapping
	var seeds []int
	var mappings []mapping

	currentMapping := mapping{}
	for _, line := range s {
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds: ") {
			line = strings.ReplaceAll(line, "seeds: ", "")
			seeds = sti.Ints(strings.Split(line, " "))
			continue
		}

		if strings.Contains(line, "-to-") {
			if len(currentMapping) > 0 {
				mappings = append(mappings, currentMapping)
			}
			currentMapping = mapping{}
			continue
		}

		values := sti.Ints(strings.Split(line, " "))
		currentMapping = append(currentMapping, submapping{
			source: values[1],
			size:   values[2],
			offset: values[0] - values[1],
		})
	}
	if len(currentMapping) > 0 {
		mappings = append(mappings, currentMapping)
	}

	lowest := math.MaxInt
	for _, seed := range seeds {
		val := seed

		for _, mapping := range mappings {
			for _, submapping := range mapping {
				if val >= submapping.source && val <= submapping.source+submapping.size {
					val += submapping.offset
					break
				}
			}
		}

		lowest = min(lowest, val)
	}

	return lowest
}
