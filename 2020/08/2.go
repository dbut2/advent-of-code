package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 8)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	for j := range s {
		done := sets.Set[int]{}
		v := 0

		i := 0
		for {
			if i == len(s) {
				return v
			}

			line := strings.Split(s[i], " ")
			cmd := line[0]
			amt := sti.Sti(line[1])

			if i == j {
				switch cmd {
				case "nop":
					cmd = "jmp"
				case "jmp":
					cmd = "nop"
				}
			}

			if done.Has(i) {
				break
			}
			done.Add(i)

			switch cmd {
			case "nop":
				i++
			case "acc":
				v += amt
				i++
			case "jmp":
				i += amt
			}
		}
	}

	return -1
}
