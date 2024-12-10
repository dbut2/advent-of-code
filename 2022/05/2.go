package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), "MCD")
	fmt.Println(solve(input))
}

func solve(input string) string {
	s := utils.ParseInput(input)

	var stacks = make(map[int]string)

	rulesStart := 0

	for i, str := range s {

		line := strings.Split(str, "")

		if line[1] == "1" {
			rulesStart = i + 2
			break
		}

		for j := 1; j < len(line); j += 4 {
			index := (j + 3) / 4

			if line[j] != " " {

				if _, ok := stacks[index]; !ok {
					stacks[index] = ""
				}

				stacks[index] = stacks[index] + line[j]
			}
		}
	}

	for i := rulesStart; i < len(s); i++ {

		rule := strings.Split(s[i], " ")

		count := sti.Sti(rule[1])
		from := sti.Sti(rule[3])
		to := sti.Sti(rule[5])

		removed, char := pop(stacks[from], count)
		stacks[from] = removed
		stacks[to] = char + stacks[to]
	}

	req := ""

	for i := 1; i <= len(stacks); i++ {
		stack := stacks[i]
		req += string(stack[0])
	}

	return req
}

func pop(s string, n int) (string, string) {
	return s[n:], s[:n]
}
