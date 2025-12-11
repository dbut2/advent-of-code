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

	start, end := "svr", "out"

	type key struct {
		node        string
		passedWords [2]bool
	}
	cache := map[key]int{}

	var dfs func(node string, passedWords [2]bool) int
	dfs = func(node string, passedWords [2]bool) int {
		key := key{node: node, passedWords: passedWords}
		if v, ok := cache[key]; ok {
			return v
		}

		if node == "dac" {
			passedWords[0] = true
		}
		if node == "fft" {
			passedWords[1] = true
		}

		if node == end {
			if passedWords[0] && passedWords[1] {
				return 1
			}
			return 0
		}

		count := 0
		for _, nextNode := range edges[node] {
			count += dfs(nextNode, passedWords)
		}
		cache[key] = count
		return count
	}

	return dfs(start, [2]bool{})
}

func main() {
	h := harness.New(solve)
	h.Expect(2, 2)
	h.Run()
}
