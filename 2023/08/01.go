package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 2)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var instruction string
	type node struct {
		left, right string
	}
	nodes := map[string]node{}

	for i, line := range s {
		if i == 0 {
			instruction = line
			continue
		}
		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, "=", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ")", "")

		parts := strings.Split(line, " ")
		nodes[parts[0]] = node{
			left:  parts[2],
			right: parts[3],
		}
	}

	count := 0
	current := "AAA"
	for {
		for _, char := range instruction {
			count++
			switch char {
			case 'L':
				current = nodes[current].left
			case 'R':
				current = nodes[current].right
			}
			if current == "ZZZ" {
				return count
			}
		}
	}
}
