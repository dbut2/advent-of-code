package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	sums := []int{}
	for _, str := range s {
		sums = append(sums, sum(stis(strings.Split(str, "\n"))))
	}

	sums = order(sums, true)

	return sum(sums[:3])
}

func order(s []int, desc bool) []int {
	t := s
	sort.Ints(t)
	if desc {
		t = reverse(t)
	}
	return t
}

func orderMap[T any](s []T, f func(T) int, desc bool) []T {
	t := s
	sort.Slice(s, func(i, j int) bool {
		return f(t[i]) < f(t[j])
	})
	if desc {
		t = reverse(t)
	}
	return t
}

func largestN(s []int, n int) []int {
	return order(s, true)[:n]
}

func largest(s []int) int {
	return largestN(s, 1)[0]
}

func largestNMap[T any](s []T, f func(T) int, n int) []T {
	return orderMap(s, f, true)[:n]
}

func largestMap[T any](s []T, f func(T) int) T {
	return largestNMap(s, f, 1)[0]
}

func smallestN(s []int, n int) []int {
	return order(s, false)[:n]
}

func smallest(s []int) int {
	return smallestN(s, 1)[0]
}

func smallestNMap[T any](s []T, f func(T) int, n int) []T {
	return orderMap(s, f, false)[:n]
}

func smallestMap[T any](s []T, f func(T) int) T {
	return smallestNMap(s, f, 1)[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func sum(s []int) int {
	t := 0
	for _, i := range s {
		t += i
	}
	return t
}

func sti(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func stis(s []string) []int {
	var i []int
	for _, str := range s {
		i = append(i, sti(str))
	}
	return i
}