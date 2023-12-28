package main

import (
	"embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 288)
	fmt.Println(solve(input))
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
