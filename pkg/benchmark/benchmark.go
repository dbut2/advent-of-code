package benchmark

import (
	"fmt"
	"runtime"
	"time"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/timer"
)

var globalPings = make(map[string]time.Time)

func Run(f func(), cond Condition) {
	var times []time.Duration
	pings := make(map[string][]time.Duration)

	start := time.Now()
	iteration := 0

	for ; cond(time.Since(start), iteration); iteration++ {
		t := timer.Start(timer.NoPrint())
		f()
		d := t.Stop()
		times = append(times, d)
		for point, ping := range globalPings {
			pings[point] = append(pings[point], t.Until(ping))
		}
	}

	took := time.Since(start)

	ot := math.Order(times, false)
	fmt.Println(len(times), "TRIALS IN", took)
	printTable(ot)

	i := 0
	for point, times := range pings {
		i++
		ot := math.Order(times, false)
		fmt.Println()
		fmt.Println("POINT:", i)
		printTable(ot)
		fmt.Println(point)
	}
}

func printTable[T any](t []T) {
	fmt.Println("MIN:", t[0])
	fmt.Println("MED:", t[len(t)/2])
	fmt.Println("MAX:", t[len(t)-1])
}

func Ping() {
	t := time.Now()
	_, file, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("%s:%d", file, line)
	globalPings[loc] = t
}

type Condition func(duration time.Duration, iteration int) bool

func Time(d time.Duration) Condition {
	return func(duration time.Duration, iteration int) bool {
		return d > duration
	}
}

func Count(n int) Condition {
	return func(duration time.Duration, iteration int) bool {
		return n > iteration
	}
}
