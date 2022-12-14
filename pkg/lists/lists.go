package lists

import (
	"github.com/dbut2/advent-of-code/pkg/math"
)

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
	min := math.Min(a, b)
	max := math.Max(a, b)

	var l []int
	for i := min; i <= max; i++ {
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

func Fill[T any](x int, def T) []T {
	var a []T
	for i := 0; i < x; i++ {
		a = append(a, def)
	}
	return a
}

func Fill2D[T any](x, y int, def T) [][]T {
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

func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func Contains[T comparable](s []T, i T) bool {
	for _, item := range s {
		if item == i {
			return true
		}
	}
	return false
}
