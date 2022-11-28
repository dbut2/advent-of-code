package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n")

	t := buildtree(s)

	ancestors := ChildrenCount(t, "shiny gold")

	fmt.Println(ancestors)
}

func Ancestors(tree map[string][]ReverseRule, color string) []string {
	ancestors := []string{}
	parents, ok := tree[color]
	if !ok {
		return []string{}
	}
	for _, parent := range parents {
		ancestors = union(ancestors, []string{parent.Parent})
		ancestors = union(ancestors, Ancestors(tree, parent.Parent))
	}
	return ancestors
}

func ChildrenCount(tree map[string][]Rule, color string) int {
	count := 0
	rules := tree[color]
	for _, rule := range rules {
		count += rule.Count
		count += rule.Count * ChildrenCount(tree, rule.Color)
	}
	return count
}

func union[T comparable](a, b []T) []T {
	u := map[T]bool{}
	for _, x := range a {
		u[x] = true
	}
	for _, x := range b {
		u[x] = true
	}
	var v []T
	for x := range u {
		v = append(v, x)
	}
	return v
}

type ReverseRule struct {
	Parent string
}

func reverseTree(rules map[string][]Rule) map[string][]ReverseRule {
	rtree := map[string][]ReverseRule{}
	for color, rs := range rules {
		for _, rule := range rs {
			if _, ok := rtree[rule.Color]; !ok {
				rtree[rule.Color] = []ReverseRule{}
			}

			rtree[rule.Color] = append(rtree[rule.Color], ReverseRule{Parent: color})
		}
	}
	return rtree
}

type Rule struct {
	Color string
	Count int
}

func buildtree(s []string) map[string][]Rule {
	rules := map[string][]Rule{}

	for _, str := range s {
		rs := []Rule{}

		r := strings.Split(str, " bags contain ")
		color := r[0]
		contain := strings.Trim(r[1], ".")
		contains := strings.Split(contain, ", ")

		for _, c := range contains {
			split1 := strings.SplitN(c, " ", 2)
			if split1[0] == "no" {
				continue
			}
			count, err := strconv.Atoi(split1[0])
			if err != nil {
				panic(err.Error())
			}
			color2 := split1[1]
			color2 = strings.TrimSuffix(color2, " bags")
			color2 = strings.TrimSuffix(color2, " bag")

			rs = append(rs, Rule{
				Color: color2,
				Count: count,
			})
		}

		rules[color] = rs
	}

	return rules
}
