package cache_test

import (
	"fmt"
	"testing"

	"github.com/dbut2/advent-of-code/pkg/cache"
)

func TestCache(t *testing.T) {
	c := cache.New()

	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 1, 2)
	c.Call(add, 3, 4)

	c.Call(fmt.Println, "Hello, world!")
	c.Call(fmt.Println, "Hello, world!")

	fmt.Println(c.Stats())
}

func BenchmarkCache(b *testing.B) {
	c := cache.New()

	fast := func(a, b int) int {
		return a + b
	}

	slow := func(a, b int) int {
		sum := 0
		for range 1000000 {
			sum += a + b
		}
		return sum
	}

	b.Run("fast", func(b *testing.B) {
		b.Run("no cache", func(b *testing.B) {
			for range b.N {
				fast(1, 2)
			}
		})

		b.Run("cache", func(b *testing.B) {
			for range b.N {
				c.Call(fast, 1, 2)
			}
		})
	})

	b.Run("slow", func(b *testing.B) {
		b.Run("no cache", func(b *testing.B) {
			for range b.N {
				slow(1, 2)
			}
		})

		b.Run("cache", func(b *testing.B) {
			for range b.N {
				c.Call(slow, 1, 2)
			}
		})
	})
}

func add(a, b int) int {
	return a + b
}
