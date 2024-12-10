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
	t.Expect(1, 301)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	m := make(map[string]string)

	for _, str := range s {
		l := strings.Split(str, ": ")
		if l[0] == "root" {
			l[1] = strings.ReplaceAll(l[1], "+", "-")
		}
		m[l[0]] = l[1]
	}

	return solvefor(m, "root", 0, "humn")
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

func treehaskey(m map[string]string, key string, find string) bool {
	v := m[key]
	l := strings.Split(v, " ")
	if len(l) == 1 {
		return key == find
	}
	return treehaskey(m, l[0], find) || treehaskey(m, l[2], find)
}

func solvefor(m map[string]string, key string, known int, find string) int {
	if key == "humn" {
		return known
	}
	v := m[key]
	l := strings.Split(v, " ")

	if treehaskey(m, l[0], find) {
		switch l[1] {
		case "+":
			newKnown := known - getnumber(m, l[2])
			return solvefor(m, l[0], newKnown, find)
		case "-":
			newKnown := known + getnumber(m, l[2])
			return solvefor(m, l[0], newKnown, find)
		case "*":
			newKnown := known / getnumber(m, l[2])
			return solvefor(m, l[0], newKnown, find)
		case "/":
			newKnown := known * getnumber(m, l[2])
			return solvefor(m, l[0], newKnown, find)
		}
	}

	if treehaskey(m, l[2], find) {
		switch l[1] {
		case "+":
			newKnown := known - getnumber(m, l[0])
			return solvefor(m, l[2], newKnown, find)
		case "-":
			newKnown := getnumber(m, l[0]) - known
			return solvefor(m, l[2], newKnown, find)
		case "*":
			newKnown := known / getnumber(m, l[0])
			return solvefor(m, l[2], newKnown, find)
		case "/":
			newKnown := getnumber(m, l[0]) / known
			return solvefor(m, l[2], newKnown, find)
		}
	}

	panic("something went wrong")
}
