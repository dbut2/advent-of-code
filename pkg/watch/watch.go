package watch

import (
	"fmt"
	"time"
)

type Watcher[T any] struct {
	val T
}

func Watch[T any](tick time.Duration) *Watcher[T] {
	w := &Watcher[T]{}
	go func() {
		t := time.NewTicker(tick)
		for {
			<-t.C
			fmt.Println(w.val)
		}
	}()
	return w
}

func (w *Watcher[T]) Update(v T) {
	w.val = v
}

func (w *Watcher[T]) Val() T {
	return w.val
}

func Incrementer(tick time.Duration) func() {
	w := Watch[int](tick)
	return func() {
		w.val++
	}
}

func Occurrence(tick time.Duration) func() {
	w := Watch[time.Time](tick)
	return func() {
		w.val = time.Now()
	}
}
