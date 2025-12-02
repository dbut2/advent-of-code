package main

import (
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input string) int {
	lines := strings.Split(input, ",")
	total := 0

	for _, line := range lines {
		parts := strings.Split(line, "-")
		start, end := atoi(parts[0]), atoi(parts[1])

		// split ranges into a list of ranges where each range has a fixed
		// number of digits
		for _, r := range splitRanges(start, end) {
			seen := map[int]bool{}
			start, end := r[0], r[1]

			// for each divisor, find all unique divisible numbers in the range
			for _, d := range divisors[log10(start)] {
				check := ((start + d - 1) / d) * d
				for check <= end {
					if !seen[check] {
						total += check
						seen[check] = true
					}
					check += d
				}
			}
		}
	}

	return total
}

var divisors = [][]int{
	2:  {11},
	3:  {111},
	4:  {101},
	5:  {11111},
	6:  {1001, 10101},
	7:  {1111111},
	8:  {10001},
	9:  {1001001},
	10: {100001, 101010101},
}

// splitRanges will return a list of ranges grouped by length of numbers
// eg [3,300] => [[3,9],[10,99],[100,300]]
func splitRanges(start, end int) [][]int {
	a, b := log10(start), log10(end)
	if a == b {
		return [][]int{{start, end}}
	}

	ranges := [][]int{{start, pow10(a) - 1}}
	for i := a; i < b-1; i++ {
		ranges = append(ranges, []int{pow10(i), pow10(i+1) - 1})
	}
	ranges = append(ranges, []int{pow10(b - 1), end})

	return ranges
}

// digits in number
func log10(n int) int {
	for i := range pows10 {
		if n < pows10[i] {
			return i
		}
	}
	panic("too big")
}

// 10 ** n
func pow10(n int) int {
	return pows10[n]
}

var pows10 = [11]int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000}

func atoi(s string) int {
	n := 0
	for i := range s {
		n *= 10
		n += int(s[i] - '0')
	}
	return n
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 4174379265)
	h.Run()
}
