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
	t.Expect(1, 436)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	lastSeen := make(map[int]int)
	currentNumber := 0
	numbers := sti.Stis(strings.Split(s[0], ","))
	i := 0
	for _, n := range numbers {
		lastNumber := currentNumber
		if _, ok := lastSeen[lastNumber]; !ok {
			lastSeen[lastNumber] = i
		}
		currentNumber = n
		lastSeen[lastNumber] = i
		i++
	}
	for ; i < 2020; i++ {
		lastNumber := currentNumber
		if _, ok := lastSeen[lastNumber]; !ok {
			lastSeen[lastNumber] = i
		}
		currentNumber = i - lastSeen[lastNumber]
		lastSeen[lastNumber] = i
	}
	return currentNumber
}
