package main

import (
	"time"

	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input string) int {
	total := 0

	var a, b int
	ranges := [2][2]int{}
	for _, c := range input {
		if c == '-' {
			a, b = b, a
			continue
		}
		if c == ',' {
			x, y := log10(a), log10(b)
			if x == y {
				ranges[0][0] = a
				ranges[0][1] = b
				ranges[1][0] = 0
				ranges[1][1] = 0
			} else {
				ranges[0][0] = a
				ranges[0][1] = pow10[x] - 1
				ranges[1][0] = pow10[x]
				ranges[1][1] = b
			}

			for _, d := range divisors[x] {
				check := ((ranges[0][0] + d - 1) / d) * d
				for check <= ranges[0][1] {
					total += check
					check += d
				}
			}

			if x == 6 {
				const d = 111111
				check := ((ranges[0][0] + d - 1) / d) * d
				for check <= ranges[0][1] {
					total -= check
					check += d
				}
			}

			if x == 10 {
				const d = 1111111111
				check := ((ranges[0][0] + d - 1) / d) * d
				for check <= ranges[0][1] {
					total -= check
					check += d
				}
			}

			if ranges[1][1] != 0 {
				z := log10(ranges[1][1])
				for _, d := range divisors[z] {
					check := ((ranges[1][0] + d - 1) / d) * d
					for check <= ranges[1][1] {
						total += check
						check += d
					}
				}

				if z == 6 {
					const d = 111111
					check := ((ranges[1][0] + d - 1) / d) * d
					for check <= ranges[1][1] {
						total -= check
						check += d
					}
				}

				if z == 10 {
					const d = 1111111111
					check := ((ranges[1][0] + d - 1) / d) * d
					for check <= ranges[1][1] {
						total -= check
						check += d
					}
				}
			}

			a, b = 0, 0
			continue
		}
		b *= 10
		b += int(c) - '0'
	}

	x, y := log10(a), log10(b)
	if x == y {
		ranges[0][0] = a
		ranges[0][1] = b
		ranges[1][0] = 0
		ranges[1][1] = 0
	} else {
		ranges[0][0] = a
		ranges[0][1] = pow10[x] - 1
		ranges[1][0] = pow10[x]
		ranges[1][1] = b
	}

	for _, d := range divisors[x] {
		check := ((ranges[0][0] + d - 1) / d) * d
		for check <= ranges[0][1] {
			total += check
			check += d
		}
	}

	if x == 6 {
		const d = 111111
		check := ((ranges[0][0] + d - 1) / d) * d
		for check <= ranges[0][1] {
			total -= check
			check += d
		}
	}

	if x == 10 {
		const d = 1111111111
		check := ((ranges[0][0] + d - 1) / d) * d
		for check <= ranges[0][1] {
			total -= check
			check += d
		}
	}

	if ranges[1][1] != 0 {
		z := log10(ranges[1][1])
		for _, d := range divisors[z] {
			check := ((ranges[1][0] + d - 1) / d) * d
			for check <= ranges[1][1] {
				total += check
				check += d
			}
		}

		if z == 6 {
			const d = 111111
			check := ((ranges[1][0] + d - 1) / d) * d
			for check <= ranges[1][1] {
				total -= check
				check += d
			}
		}

		if z == 10 {
			const d = 1111111111
			check := ((ranges[1][0] + d - 1) / d) * d
			for check <= ranges[1][1] {
				total -= check
				check += d
			}
		}
	}

	return total
}

var divisors = [11][]int{
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

// digits in number
// binary search log
func log10(n int) int {
	if n >= 100000 {
		if n >= 10000000 {
			if n >= 1000000000 {
				return 10
			}
			if n >= 100000000 {
				return 9
			}
			return 8
		}
		if n >= 1000000 {
			return 7
		}
		return 6
	}
	if n >= 1000 {
		if n >= 10000 {
			return 5
		}
		return 4
	}
	if n >= 100 {
		return 3
	}
	if n >= 10 {
		return 2
	}
	return 1
}

var pow10 = [11]int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000}

func main() {
	h := harness.New(solve, harness.WithNoSubmit[string, int]())
	h.Benchmark(benchmark.Time(time.Second))
	h.Expect(1, 4174379265)
	h.Run()
}
