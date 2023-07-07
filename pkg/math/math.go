package math

import (
	"sort"
)

func Max[N Number](a, b N) N {
	if a > b {
		return a
	}
	return b
}

func Min[N Number](a, b N) N {
	if a < b {
		return a
	}
	return b
}

func Abs[N Number](a N) N {
	if a < 0 {
		return -a
	}
	return a
}

func Sign[N Snumber](a N) N {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

func Order[T Number](s []T, desc bool) []T {
	t := s
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
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

func LargestN[N Number](s []N, n int) []N {
	return Order(s, true)[:n]
}

func Largest[N Number](s []N) N {
	return LargestN(s, 1)[0]
}

func LargestNMap[T any](s []T, f func(T) int, n int) []T {
	return OrderMap(s, f, true)[:n]
}

func LargestMap[T any](s []T, f func(T) int) T {
	return LargestNMap(s, f, 1)[0]
}

func SmallestN[N Number](s []N, n int) []N {
	return Order(s, false)[:n]
}

func Smallest[N Number](s []N) N {
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

type Number interface {
	Sint | Uint | Float
}

type Snumber interface {
	Sint | Float
}

type Int interface {
	Sint | Uint
}

type Sint interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type SizedInt interface {
	SizedSint | SizedUint
}

type SizedSint interface {
	~int8 | ~int16 | ~int32 | ~int64
}

type SizedUint interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Sum[T Number](s []T) T {
	var t T
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

func Pow[N Int, M Uint](x N, y M) N {
	result := N(1)
	for y > 0 {
		if y&1 == 1 {
			result *= x
		}
		y = y >> 1
		x *= x
	}
	return result
}
