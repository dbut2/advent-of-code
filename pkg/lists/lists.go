package lists

import (
	"slices"

	"github.com/dbut2/advent-of-code/pkg/math"
)

func Intersection[T comparable](a, b []T) []T {
	var i []T
	for _, x := range a {
		if slices.Contains(b, x) {
			i = append(i, x)
		}
	}
	return i
}

func Range[N math.Number](a, b N) []N {
	low := min(a, b)
	high := max(a, b)

	l := make([]N, 0, int(high-low+1))
	for i := low; i <= high; i++ {
		l = append(l, i)
	}
	return l
}

func Map[T, U any](s []T, f func(T) U) []U {
	l := make([]U, len(s))
	for i, v := range s {
		l[i] = f(v)
	}
	return l
}

func MapMap[T, U any, V, W comparable](s map[V]T, f func(V, T) (W, U)) map[W]U {
	l := make(map[W]U, len(s))
	for k, v := range s {
		nk, nv := f(k, v)
		l[nk] = nv
	}
	return l
}

func MapToSlice[T comparable, U any](s map[T]U) Pairs[T, U] {
	l := make(Pairs[T, U], 0, len(s))
	for k, v := range s {
		l = append(l, Pair[T, U]{A: k, B: v})
	}
	return l
}

type Pair[T, U any] struct {
	A T
	B U
}

type Pairs[T, U any] []Pair[T, U]

func (p Pairs[T, U]) Keys() []T {
	l := make([]T, len(p))
	for i := range p {
		l[i] = p[i].A
	}
	return l
}

func (p Pairs[T, U]) Vals() []U {
	l := make([]U, len(p))
	for i := range p {
		l[i] = p[i].B
	}
	return l
}

func Fill[T any](x int, def T) []T {
	a := make([]T, x)
	for i := 0; i < x; i++ {
		a[x] = def
	}
	return a
}

func Fill2D[T any](x, y int, def T) [][]T {
	a := make([][]T, x)
	for i := 0; i < x; i++ {
		b := make([]T, y)
		for j := 0; j < y; j++ {
			b[j] = def
		}
		a[i] = b
	}
	return a
}

func Reverse[T any](s []T) []T {
	r := slices.Clone(s)
	slices.Reverse(r)
	return r
}

func Contains[T comparable](s []T, i T) bool {
	return slices.Contains(s, i)
}
