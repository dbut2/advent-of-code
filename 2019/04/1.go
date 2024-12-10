package main

import (
	"embed"
	"strconv"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	ranges := sti.Stis(strings.Split(s[0], "-"))

	count := 0
	for i := ranges[0]; i <= ranges[1]; i++ {

		n := strconv.Itoa(i)

		hasDouble := false
		for j := 1; j < len(n); j++ {
			if n[j] == n[j-1] {
				hasDouble = true
				break
			}
		}

		decreasing := true
		for j := 1; j < len(n); j++ {
			if n[j] < n[j-1] {
				decreasing = false
				break
			}
		}

		if hasDouble && decreasing {
			count++
		}
	}

	return count
}
