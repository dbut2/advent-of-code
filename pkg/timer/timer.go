package timer

import (
	"fmt"
	"time"
)

type timer struct {
	print bool
	start time.Time
}

func Start(opts ...Option) *timer {
	t := &timer{print: true}
	for _, opt := range opts {
		opt(t)
	}
	t.start = time.Now()
	return t
}

func (t *timer) Stop() time.Duration {
	d := time.Since(t.start)
	if t.print {
		fmt.Println(time.Since(t.start))
	}
	return d
}

func (t *timer) Until(end time.Time) time.Duration {
	return end.Sub(t.start)
}

type Option func(*timer)

func NoPrint() Option {
	return func(t *timer) {
		t.print = false
	}
}
