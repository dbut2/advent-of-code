package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/utils"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	fmt.Println("Test")
	fmt.Println(do(test))
	fmt.Println()
	fmt.Println("Solution")
	fmt.Println(do(input))
}

func do(s string) string {
	strs := strings.Split(s, "\n")
	return solve(strs)
}

func solve(s []string) string {

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

		count := utils.Sti(rule[1])
		from := utils.Sti(rule[3])
		to := utils.Sti(rule[5])

		for j := 0; j < count; j++ {
			removed, char := pop(stacks[from])
			stacks[from] = removed
			stacks[to] = char + stacks[to]
		}
	}

	req := ""

	for i := 1; i <= len(stacks); i++ {
		stack := stacks[i]
		req += string(stack[0])
	}

	return req
}

func pop(s string) (string, string) {
	return s[1:], s[:1]
}