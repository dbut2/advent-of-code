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
	polymer := s[0]

	rules := map[string]string{}

	for i := 2; i < len(s); i++ {
		rule := strings.Split(s[i], " ")

		rules[rule[0]] = rule[2]

	}

	for i := 0; i < 10; i++ {
		polymer = step(polymer, rules)
	}

	counts := map[string]int{}

	for _, mol := range strings.Split(polymer, "") {
		counts[mol]++
	}

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

func step(polymer string, rules map[string]string) string {
	mols := strings.Split(polymer, "")

	newPolymer := mols[0]

	for i := 1; i < len(mols); i++ {

		if match, ok := rules[mols[i-1]+mols[i]]; ok {
			newPolymer += match
		}
		newPolymer += mols[i]

	}

	return newPolymer
}
