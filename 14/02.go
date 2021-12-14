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
	pairs := map[string]int{}
	polymer := s[0]

	for i := 1; i < len(polymer); i++ {
		pairs[string(polymer[i-1])+string(polymer[i])]++
	}

	rules := map[string]string{}

	for i := 2; i < len(s); i++ {
		rule := strings.Split(s[i], " ")
		rules[rule[0]] = rule[2]
	}

	for i := 0; i < 40; i++ {
		pairs = step(pairs, rules)
	}

	counts := map[string]int{}

	for pair, count := range pairs {
		counts[string(pair[1])] += count
	}

	counts[string(polymer[0])]++

	max := -1
	min := -1

	for _, count := range counts {
		if count > max || max == -1 {
			max = count
		}

		if count < min || min == -1 {
			min = count
		}
	}

	return max - min

}

func step(pairs map[string]int, rules map[string]string) map[string]int {
	newPairs := map[string]int{}

	for pair, count := range pairs {
		if match, ok := rules[pair]; ok {
			mols := strings.Split(pair, "")
			newPairs[mols[0]+match] += count
			newPairs[match+mols[1]] += count
		} else {
			newPairs[pair] += count
		}
	}

	return newPairs
}
