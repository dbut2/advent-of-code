package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	connections := map[string][]string{}

	for _, line := range s {
		split := strings.Split(line, "-")

		connections[split[0]] = append(connections[split[0]], split[1])
		connections[split[1]] = append(connections[split[1]], split[0])
	}

	paths := getPaths(connections, []string{"start"})

	return len(paths)
}

func getPaths(connections map[string][]string, stack []string) [][]string {
	var paths [][]string

	current := stack[len(stack)-1]

	if current == "end" {
		return [][]string{stack}
	}

	proposals := filterProposals(stack, connections[current])
	for _, match := range proposals {
		tempStack := append(stack, match)
		paths = append(paths, getPaths(connections, tempStack)...)
	}

	return paths
}

func filterProposals(stack []string, proposals []string) []string {
	var matches []string

	for _, proposal := range proposals {
		if proposal == "start" {
			continue
		}

		if stack[len(stack)-1] == "end" {
			continue
		}

		lowerAndSeenTwiceBefore := false
		if proposal == strings.ToLower(proposal) {
			seen := 0
			for _, cave := range stack {
				if proposal == cave {
					seen++
				}
			}

			canSeeTwo := true
			caves := map[string]int{}
			for _, cave := range stack {
				if cave == strings.ToUpper(cave) {
					continue
				}
				caves[cave]++
			}
			for _, count := range caves {
				if count >= 2 {
					canSeeTwo = false
				}
			}
			if !canSeeTwo && seen >= 1 {
				lowerAndSeenTwiceBefore = true
			}

			if seen >= 2 {
				lowerAndSeenTwiceBefore = true
			}
		}

		if lowerAndSeenTwiceBefore {
			continue
		}

		matches = append(matches, proposal)
	}

	return matches
}
