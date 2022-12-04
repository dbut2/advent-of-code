package utils

import (
	"sort"
	"strconv"
)

func Order(s []int, desc bool) []int {
	t := s
	sort.Ints(t)
	if desc {
		t = Reverse(t)
	}
	return t
}

func OrderMap[T any](s []T, f func(T) int, desc bool) []T {
	t := s
	sort.Slice(s, func(i, j int) bool {
		return f(t[i]) < f(t[j])
	})
	if desc {
		t = Reverse(t)
	}
	return t
}

func LargestN(s []int, n int) []int {
	return Order(s, true)[:n]
}

func largest(s []int) int {
	return LargestN(s, 1)[0]
}

func LargestNMap[T any](s []T, f func(T) int, n int) []T {
	return OrderMap(s, f, true)[:n]
}

func LargestMap[T any](s []T, f func(T) int) T {
	return LargestNMap(s, f, 1)[0]
}

func SmallestN(s []int, n int) []int {
	return Order(s, false)[:n]
}

func Smallest(s []int) int {
	return SmallestN(s, 1)[0]
}

func SmallestNMap[T any](s []T, f func(T) int, n int) []T {
	return OrderMap(s, f, false)[:n]
}

func SmallestMap[T any](s []T, f func(T) int) T {
	return SmallestNMap(s, f, 1)[0]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func Sum(s []int) int {
	t := 0
	for _, i := range s {
		t += i
	}
	return t
}

func Sti(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func Stis(s []string) []int {
	var i []int
	for _, str := range s {
		i = append(i, Sti(str))
	}
	return i
}

func Intersection[T comparable](a, b []T) []T {
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

func Range(a, b int) []int {
	var l []int
	for i := a; i <= b; i++ {
		l = append(l, i)
	}
	return l
}
