package main

import (
	"math"
	"math/bits"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input []string) int {
	total := 0
	for _, line := range input {
		parts := strings.Split(line, " ")
		targetStr := parts[0]
		targetStr = targetStr[1 : len(targetStr)-1]
		buttonsStrs := parts[1 : len(parts)-1]

		var buttons []int
		for _, buttonStr := range buttonsStrs {
			button := 0
			for _, b := range Ints(buttonStr) {
				button |= 1 << b
			}
			buttons = append(buttons, button)
		}

		target := 0
		for i, c := range targetStr {
			if c == '#' {
				target |= 1 << i
			}
		}

		m := math.MaxInt
		for i := range 1 << len(buttons) {
			state := 0
			for j, mask := range buttons {
				if i>>j&1 == 1 {
					state ^= mask
				}
			}
			if state == target {
				m = min(m, bits.OnesCount(uint(i)))
			}
		}
		total += m
	}
	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 7)
	h.Run()
}
