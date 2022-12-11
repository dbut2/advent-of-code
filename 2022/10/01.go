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
	t.Expect(1, 13140)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	x := 1
	cycle := 1

	checkCycles := []int{
		20, 60, 100, 140, 180, 220,
	}

	signal := 0

	for _, line := range s {

		split := strings.Split(line, " ")

		cmd := split[0]
		amt := 0

		if len(split) > 1 {
			amt = sti.Sti(split[1])
		}

		switch cmd {
		case "noop":
			cycle += 1
			if isin(checkCycles, cycle) {
				signal += cycle * x
			}

		case "addx":
			cycle++
			if isin(checkCycles, cycle) {
				signal += cycle * x
			}
			cycle++
			x += amt
			if isin(checkCycles, cycle) {
				signal += cycle * x
			}
		}

	}

	return signal
}

func isin[T comparable](s []T, i T) bool {
	for _, item := range s {
		if item == i {
			return true
		}
	}
	return false
}
