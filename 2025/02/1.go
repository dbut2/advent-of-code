package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input string) int {
	lines := strings.Split(input, ",")
	total := 0

	for _, line := range lines {
		parts := strings.Split(line, "-")

		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		// test every number in range inclusively
		for i := start; i <= end; i++ {
			num := fmt.Sprint(i)

			l := len(num)
			// skip odd length numbers
			if l%2 != 0 {
				continue
			}

			// add as invalid if first and second half are the same
			if num[:l/2] == num[l/2:] {
				total += i
			}
		}
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Benchmark(benchmark.Count(100))
	h.Expect(1, 1227775554)
	h.Run()
}
