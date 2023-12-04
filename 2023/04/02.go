package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(2, 30)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	//Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

	total := 0
	multis := map[int]int{}

	for i := range s {
		multis[i] = 1
	}

	for i, line := range s {
		game := strings.Split(line, ": ")

		games := strings.Split(game[1], "|")

		left := strings.Split(games[0], " ")
		right := strings.Split(games[1], " ")

		leftNumbers := []int{}
		for _, n := range left {
			if n == "" {
				continue
			}
			n = strings.TrimSpace(n)
			leftNumbers = append(leftNumbers, sti.Sti(n))
		}

		rightNumbers := []int{}
		for _, n := range right {
			if n == "" {
				continue
			}
			n = strings.TrimSpace(n)
			rightNumbers = append(rightNumbers, sti.Sti(n))
		}

		matches := lists.Intersection(leftNumbers, rightNumbers)

		multiplier := multis[i]

		total += multiplier

		for j, _ := range matches {
			multis[i+j+1] += multiplier
		}
	}

	return total
}
