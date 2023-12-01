package math

import (
	"math"
	"testing"
)

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pow(100, 100)
	}
}

func BenchmarkOldPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(float64(100), float64(100))
	}
}
