package utils

import (
	"sort"
	"strconv"
	"strings"
)

func ParseInput(s string) []string {
	return strings.Split(s, "\n")
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

func Stiss(s []string) [][]int {
	var i [][]int
	for _, line := range s {
		i = append(i, Stis(strings.Split(line, "")))
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

func Map[T, U any](s []T, f func(T) U) []U {
	var l []U
	for _, v := range s {
		l = append(l, f(v))
	}
	return l
}

func MapMap[T, U any, V, W comparable](s map[V]T, f func(V, T) (W, U)) map[W]U {
	l := make(map[W]U)
	for k, v := range s {
		nk, nv := f(k, v)
		l[nk] = nv
	}
	return l
}

func MapToSlice[T comparable, U any](s map[T]U) []Pair[T, U] {
	var l []Pair[T, U]
	for k, v := range s {
		l = append(l, Pair[T, U]{A: k, B: v})
	}
	return l
}

type Pair[T, U any] struct {
	A T
	B U
}

func Fill2D[T comparable](x, y int, def T) [][]T {
	var a [][]T
	for i := 0; i < x; i++ {
		var b []T
		for j := 0; j < y; j++ {
			b = append(b, def)
		}
		a = append(a, b)
	}
	return a
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
