package timer

import (
	"fmt"
	"time"
)

type timer struct {
	print bool
}

func Start(opts ...Option) func() time.Duration {
	t := &timer{print: true}
	for _, opt := range opts {
		opt(t)
	}
	start := time.Now()
	return func() time.Duration {
		d := time.Since(start)
		if t.print {
			fmt.Println(time.Since(start))
		}
		return d
	}
}

type Option func(*timer)

func NoPrint() Option {
	return func(t *timer) {
		t.print = false
	}
}
