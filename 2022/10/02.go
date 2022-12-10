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
	t.Expected(1, 0)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	x := 1
	cycle := 1

	signal := 0

	chars := [240]bool{}

	for _, line := range s {
		split := strings.Split(line, " ")

		cmd := split[0]
		amt := 0

		if len(split) > 1 {
			amt = sti.Sti(split[1])
		}

		switch cmd {
		case "noop":
			if check(cycle, x) {
				chars[cycle] = true
			}
			cycle++
		case "addx":
			if check(cycle, x) {
				chars[cycle] = true
			}
			cycle++
			x += amt
			if check(cycle, x) {
				chars[cycle] = true
			}
			cycle++
		}

	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if chars[i*40+j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	return signal
}

func check(cycle, x int) bool {
	return (cycle%40)-1 == x || (cycle%40) == x || (cycle%40)+1 == x
}
