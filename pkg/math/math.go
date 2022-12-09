package math

import (
	"sort"
)

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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Sign(a int) int {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

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

func SumMap[T comparable](s map[T]int) int {
	t := 0
	for _, i := range s {
		t += i
	}
	return t
}

func SumMapIf[T comparable](s map[T]int, predicate func(T) bool) int {
	t := 0
	for k, v := range s {
		if predicate(k) {
			t += v
		}
	}
	return t
}
