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

	visitedSmall := 0

	for _, path := range paths {
		seenSmall := false
		for _, cave := range path {
			if cave == "start" {
				continue
			}

			if cave == "end" {
				continue
			}

			if cave == strings.ToLower(cave) {
				seenSmall = true
			}
		}

		if seenSmall {
			visitedSmall++
		}
	}

	return visitedSmall
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

		lowerAndSeenBefore := false
		if proposal == strings.ToLower(proposal) {
			for _, cave := range stack {
				if proposal == cave {
					lowerAndSeenBefore = true
				}
			}
		}
		if lowerAndSeenBefore {
			continue
		}

		matches = append(matches, proposal)
	}

	return matches
}
