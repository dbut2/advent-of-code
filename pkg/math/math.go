package math

import (
	"math"
)

const MaxInt = math.MaxInt
const MinInt = math.MinInt

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

func Sum[T Number](s ...T) T {
	var t T
	for _, i := range s {
		t += i
	}
	return t
}

func Product[T Number](nums ...T) T {
	p := T(1)
	for _, v := range nums {
		p *= v
	}
	return p
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

func Pow[N Number, M Int](x N, y M) N {
	if y < 0 {
		return 1 / Pow(x, -y)
	}

	result := N(1)
	for y > 0 {
		if y&1 == 1 {
			result *= x
		}
		y >>= 1
		x *= x
	}
	return result
}

func GCD[T Int](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM[T Int](a, b T) T {
	l := a * b / GCD(a, b)
	return l
}
