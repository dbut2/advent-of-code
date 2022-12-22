package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

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
	t.Expect(1, 152)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	m := make(map[string]string)

	for _, str := range s {
		l := strings.Split(str, ": ")
		m[l[0]] = l[1]
	}

	return getnumber(m, "root")
}

func getnumber(m map[string]string, key string) int {
	v := m[key]
	l := strings.Split(v, " ")
	if len(l) == 1 {
		return sti.Sti(l[0])
	}
	switch l[1] {
	case "+":
		return getnumber(m, l[0]) + getnumber(m, l[2])
	case "-":
		return getnumber(m, l[0]) - getnumber(m, l[2])
	case "*":
		return getnumber(m, l[0]) * getnumber(m, l[2])
	case "/":
		return getnumber(m, l[0]) / getnumber(m, l[2])
	default:
		panic("unknown operand")
	}
}
