package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, ",")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {

	var fishes []*LanternFish

	fishes = []*LanternFish{}

	for _, f := range s {
		left, err := strconv.Atoi(f)
		if err != nil {
			panic(err.Error())
		}
		fish := &LanternFish{
			timer: left,
		}
		fishes = append(fishes, fish)
	}

	days := 80

	for i := 0; i < days; i++ {
		for _, fish := range fishes {
			newFish := fish.Progress()
			if newFish != nil {
				fishes = append(fishes, newFish)
			}
		}
	}

	return len(fishes)
}

type LanternFish struct {
	timer int
}

func (l *LanternFish) Progress() *LanternFish {
	if l.timer > 0 {
		l.timer--
		return nil
	} else {
		l.timer = 6
		return &LanternFish{
			timer: 8,
		}
	}
}
