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

			// split number into n pieces for every factor of the length
			for j := 2; j <= l; j++ {
				if l%j != 0 {
					continue
				}

				// if all pieces other than the first match with the first,
				// the whole number must be a repeating sequence
				all := true
				for k := 1; k < j; k++ {
					if line[:l/j] != line[k*(l/j):(k+1)*(l/j)] {
						all = false
					}
				}
				if all {
					total += i
					break
				}
			}
		}
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Benchmark(benchmark.Count(30))
	h.Expect(1, 4174379265)
	h.Run()
}
