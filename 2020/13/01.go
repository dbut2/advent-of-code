package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expected(1, 295)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	current := sti.Sti(s[0])
	earliest := -1
	var earliestBus int
	for _, busid := range strings.Split(s[1], ",") {
		if busid == "x" {
			continue
		}
		bus := sti.Sti(busid)
		until := bus - (current % bus)
		if earliest == -1 || until < earliest {
			earliest = until
			earliestBus = bus
		}
	}

	return earliest * earliestBus
}
