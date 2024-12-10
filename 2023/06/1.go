package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 288)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	times := strings.Ints(s[0])
	distances := strings.Ints(s[1])

	total := 1
	for i := range times {

		perms := 0
		for j := 0; j <= times[i]; j++ {
			movedDistance := j * (times[i] - j)
			if movedDistance > distances[i] {
				perms++
			}
		}
		total *= perms
	}

	return total
}
