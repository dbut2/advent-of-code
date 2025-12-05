package main

import (
	"slices"
	"strings"
	"time"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

func solve(input [2][]string) int {
	var ranges [][]int
	for _, line := range input[0] {
		ranges = append(ranges, sti.Ints(strings.Split(line, "-")))
	}

	// store the amount of ranges we enter and exit for an index
	stepMap := map[int]int{}
	for _, r := range ranges {
		stepMap[r[0]]++
		stepMap[r[1]+1]-- // decrement after the index
	}

	// take our entry and exit indexes and create an ordered list of indexes
	steps := [][2]int{}
	for index, increase := range stepMap {
		if increase == 0 { // 0 step is a no-op
			continue
		}
		steps = append(steps, [2]int{index, increase})
	}
	slices.SortFunc(steps, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	// iterate each range to find overlapping sets of ranges, and their lengths
	count := 0
	lastStart := 0
	rangeCounter := 0
	for _, step := range steps {
		// if we step up from 0, this is the start of set of overlapping ranges
		if step[1] >= 1 && rangeCounter == 0 {
			lastStart = step[0]
		}

		// make step
		rangeCounter += step[1]

		// if we step down to 0, we're at the tail end of a set of ranges
		// increment our counter by the length
		if step[1] <= -1 && rangeCounter == 0 {
			count += step[0] - lastStart
		}
	}

	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 14)
	h.Benchmark(time.Second)
	h.Run()
}
