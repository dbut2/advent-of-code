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
	t.Expected(1, 1068781)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	i := 0
	period := 1
	for {
		for j, busid := range strings.Split("x,"+s[1], ",") {
			if busid == "x" {
				continue
			}
			bus := sti.Sti(busid)

			for (i+j)%bus != 0 {
				i += period
			}
			period *= bus
		}
		return i + 1
	}
}
