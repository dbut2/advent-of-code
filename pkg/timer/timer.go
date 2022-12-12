package timer

import (
	"fmt"
	"time"
)

func Start() func() {
	start := time.Now()
	return func() {
		fmt.Println(time.Since(start))
	}
}
