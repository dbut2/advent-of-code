package main

import (
	"embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/buffers"
	"github.com/dbut2/advent-of-code/pkg/chars"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 71503)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	in := [2]int{}
	for i, line := range s {
		buffer := buffers.Number(0)
		for _, char := range line {
			if chars.IsNum(char) {
				buffer.Add(chars.NumVal(char))
			}
		}
		in[i] = buffer.Clear()
	}

	time := in[0]
	distance := in[1]

	perms := 0
	for j := 0; j <= time; j++ {
		movedDistance := j * (time - j)
		if movedDistance > distance {
			perms++
		}
	}

	return perms
}
