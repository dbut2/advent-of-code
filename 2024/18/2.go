package main

import (
	"time"

	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input [][]int) [2]int {
	grid := [71][71]bool{}
	i, j, k := 0, len(input)/2, len(input)
	lastJ := 0
	for i < k-1 {
		if lastJ < j {
			for _, cell := range input[lastJ:j] {
				grid[cell[0]][cell[1]] = true
			}
		} else {
			for _, cell := range input[j:lastJ] {
				grid[cell[0]][cell[1]] = false
			}
		}
		lastJ = j
		state := grid
		stack := [][2]int{{0, 0}}
		end := false
		for len(stack) > 0 {
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if state[cell[0]][cell[1]] {
				continue
			}
			state[cell[0]][cell[1]] = true
			{
				next := cell
				next[0]++
				if next[0] >= len(grid) {
					goto a
				}
				if grid[next[0]][next[1]] {
					goto a
				}
				if next == [2]int{70, 70} {
					end = true
					break
				}
				stack = append(stack, next)
			}
		a:
			{
				next := cell
				next[1]++
				if next[1] >= len(grid[0]) {
					goto c
				}
				if grid[next[0]][next[1]] {
					goto c
				}
				if next == [2]int{70, 70} {
					end = true
					break
				}
				stack = append(stack, next)
			}
		c:
			{
				next := cell
				next[0]--
				if next[0] < 0 {
					goto b
				}
				if grid[next[0]][next[1]] {
					goto b
				}
				if next == [2]int{70, 70} {
					end = true
					break
				}
				stack = append(stack, next)
			}
		b:
			{
				next := cell
				next[1]--
				if next[1] < 0 {
					goto d
				}
				if grid[next[0]][next[1]] {
					goto d
				}
				if next == [2]int{70, 70} {
					end = true
					break
				}
				stack = append(stack, next)
			}
		d:
		}
		if end {
			i = j
		} else {
			k = j
		}
		j = i + (k-i)/2
	}
	return [2]int(input[j])
}

func main() {
	h := harness.New(solve)
	h.Benchmark(benchmark.Time(time.Second))
	h.Run()
}
