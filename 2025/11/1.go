package main

import (
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []string) int {
	edges := map[string][]string{}

	for _, line := range input {
		parts := strings.Split(line, " ")
		key := parts[0]
		key = key[:len(key)-1]
		rest := parts[1:]

		edges[key] = rest
	}

	start, end := "you", "out"

	var dfs func(string) int
	dfs = func(node string) int {
		if node == end {
			return 1
		}

		count := 0
		for _, nextNode := range edges[node] {
			count += dfs(nextNode)
		}
		return count
	}

	return dfs(start)
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 5)
	h.Run()
}
