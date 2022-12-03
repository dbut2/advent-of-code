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
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {

	count := 0

	for i := 0; i < len(s); i += 3 {

		l1 := strings.Split(s[i], "")
		l2 := strings.Split(s[i+1], "")
		l3 := strings.Split(s[i+2], "")

		inboth := intersection(intersection(l1, l2), l3)[0]

		count += priority(inboth)
	}

	return count
}

func priority(a string) int {
	b := []byte(a)[0]

	if b > 96 {
		b -= 58
	}

	b -= 38

	return int(b)
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

func intersection[T comparable](a, b []T) []T {
	var i []T
	for _, x := range a {
		for _, y := range b {
			if x == y {
				i = append(i, x)
			}
		}
	}
	return i
}
