package std

import (
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/strings"
)

type Grid = space.Grid[byte]

func Abs[N math.Number](n N) N {
	return math.Abs(n)
}

func Fill[T any](x int, def T) []T {
	return lists.Fill(x, def)
}

func Filter[S ~[]E, E any](s S, predicate func(E) bool) S {
	return lists.Filter(s, predicate)
}

func GCD[T math.Int](a, b T) T {
	return math.GCD(a, b)
}

func Int(s string) int {
	return sti.Int(s)
}

func Ints(s string) []int {
	return strings.Ints(s)
}

func LCM[T math.Int](a, b T) T {
	return math.LCM(a, b)
}

func Pow[N math.Number, M math.Int](x N, y M) N {
	return math.Pow(x, y)
}

func Product[T math.Number](nums ...T) T {
	return math.Product(nums...)
}

func Reduce[T any, U any](s []T, f func(U, T) U) U {
	return lists.Reduce(s, f)
}

func Sign[N math.Snumber](a N) N {
	return math.Sign(a)
}

func Sum[T math.Number](s ...T) T {
	return math.Sum(s...)
}
