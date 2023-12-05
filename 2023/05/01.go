package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 35)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var seeds []int
	var mappings []mapping

	currentMapping := mapping{}
	for _, line := range s {
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds: ") {
			seedsString := strings.Split(line, "seeds: ")
			seedList := strings.Split(seedsString[1], " ")
			for _, seedItem := range seedList {
				seeds = append(seeds, sti.Sti(seedItem))
			}
			continue
		}

		if strings.Contains(line, "-") {
			if len(currentMapping) > 0 {
				mappings = append(mappings, currentMapping)
			}
			currentMapping = mapping{}
			continue
		}

		values := strings.Split(line, " ")

		currentMapping = append(currentMapping, submapping{
			source: sti.Sti(values[1]),
			size:   sti.Sti(values[2]),
			offset: sti.Sti(values[0]) - sti.Sti(values[1]),
		})
	}
	if len(currentMapping) > 0 {
		mappings = append(mappings, currentMapping)
	}

	lowest := -1
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

		if lowest == -1 {
			lowest = val
		}
		lowest = min(lowest, val)
	}

	return lowest
}

type mapping []submapping

type submapping struct {
	source, size, offset int
}
