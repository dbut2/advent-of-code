package main

import (
	"embed"
	_ "embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/ll"
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
	t.Expect(1, 1623178306)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var nodes []*ll.Double[int]

	var prev *ll.Double[int]
	for _, str := range s {
		node := &ll.Double[int]{Val: sti.Sti(str) * 811589153}
		nodes = append(nodes, node)
		ll.Link(prev, node)
		prev = node
	}

	ll.Link(nodes[len(nodes)-1], nodes[0])

	for i := 0; i < 10; i++ {
		for _, node := range nodes {
			if node.Val < 0 {
				for i := 0; i > node.Val%(len(nodes)-1); i-- {
					node.MoveLeft()
				}
			} else {
				for i := 0; i < node.Val%(len(nodes)-1); i++ {
					node.MoveRight()
				}
			}

			a := ""
			_ = a
		}
	}

	val := 0

	cur := ll.Find(nodes[0], 0)
	for i := 0; i <= 3000; i++ {
		if i == 1000 || i == 2000 || i == 3000 {
			fmt.Println(cur.Val)
			val += cur.Val
		}

		cur = cur.Next
	}

	return val
}
