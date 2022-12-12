package benchmark

import (
	"fmt"
	"time"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/timer"
)

func Run(n int, f func()) {
	var times []time.Duration

	for i := 0; i < n; i++ {
		stop := timer.Start(timer.NoPrint())
		f()
		d := stop()
		times = append(times, d)
	}

	ot := math.Order(times, false)
	fmt.Println("TRIALS:", len(ot))
	fmt.Println("MIN:", ot[0])
	fmt.Println("MED:", ot[len(ot)/2])
	fmt.Println("MAX:", ot[len(ot)-1])
}
