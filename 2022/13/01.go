package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 13)
	fmt.Println(solve(input))
}

func solve(input string) int {
	pairs := strings.Split(input, "\n\n")

	indices := 0

	for i, pair := range pairs {
		p := strings.Split(pair, "\n")

		left := parse(p[0])
		right := parse(p[1])

		if isInOrder(left, right) > -1 {
			indices += i + 1
		}
	}

	return indices
}

// return inorder, ok
func isInOrder(a, b intorlist) int {
	switch a.(type) {
	case Int:
		switch b.(type) {
		case Int:
			// a int b int
			a2 := int(a.(Int))
			b2 := int(b.(Int))

			diff := math.Sign(b2 - a2)

			return diff
		default:
			// a int b list

			return isInOrder(List([]intorlist{a}), b)
		}
	default:
		switch b.(type) {
		case Int:
			// a list b int

			return isInOrder(a, List([]intorlist{b}))
		default:
			//a list b list

			a2 := a.(List)
			b2 := b.(List)

			smaller := math.Min(len(a2), len(b2))
			asmaller := math.Sign(len(b2) - len(a2))

			for i := 0; i < smaller; i++ {
				ord := isInOrder(a2[i], b2[i])
				if ord != 0 {
					return ord
				}
			}

			// -1 for a bigger, 0 for same, 1 for smaller
			return asmaller
		}
	}
}

type intorlist interface {
	getVal()
}

type Int int

func (i Int) getVal() {}

type List []intorlist

func (l List) getVal() {}

func parse(s string) intorlist {

	level := -1
	buffer := ""
	is := make(map[int]List)

	for _, char := range strings.Split(s, "") {
		switch char {
		case "[":
			level++
			is[level] = List{}
		case "]":
			if buffer != "" {
				num := sti.Sti(buffer)
				buffer = ""
				is[level] = append(is[level], Int(num))
			}
			level--
			is[level] = append(is[level], is[level+1])
		case ",":
			// buffer nil if last item was list, already added
			if buffer != "" {
				num := sti.Sti(buffer)
				buffer = ""
				is[level] = append(is[level], Int(num))
			}
		default:
			buffer += char
		}
	}
	return is[0]
}
