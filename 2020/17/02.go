package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 848)
	fmt.Println(solve(input))
}

type Space struct {
	s sets.Set[string]
}

func (s *Space) Get(x, y, z, w int) bool {
	return s.s.Has(key(x, y, z, w))
}
func (s *Space) Set(x, y, z, w int, set bool) {
	if set {
		s.s.Add(key(x, y, z, w))
	} else {
		s.s.Remove(key(x, y, z, w))
	}
}

func key(x, y, z, w int) string {
	return fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)
}
func dekey(key string) (int, int, int, int) {
	vs := sti.Stis(strings.Split(key, ","))
	return vs[0], vs[1], vs[2], vs[3]
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	var space Space

	for x, str := range s {
		for y, active := range strings.Split(str, "") {
			if active == "#" {
				space.Set(x, y, 0, 0, true)
			}
		}
	}

	minx, maxx := 0, len(s)-1
	miny, maxy := 0, len(strings.Split(s[0], ""))-1
	minz, maxz := 0, 0
	minw, maxw := 0, 0

	for i := 1; i <= 6; i++ {
		minx--
		maxx++
		miny--
		maxy++
		minz--
		maxz++
		minw--
		maxw++

		space = cycle(space, minx, maxx, miny, maxy, minz, maxz, minw, maxw)
	}

	return len(space.s)
}

func cycle(space Space, minx, maxx, miny, maxy, minz, maxz, minw, maxw int) Space {
	var newSpace Space

	for i := minx; i <= maxx; i++ {
		for j := miny; j <= maxy; j++ {
			for k := minz; k <= maxz; k++ {
				for l := minw; l <= maxw; l++ {
					count := 0
					for x := i - 1; x <= i+1; x++ {
						for y := j - 1; y <= j+1; y++ {
							for z := k - 1; z <= k+1; z++ {
								for w := l - 1; w <= l+1; w++ {
									if x == i && y == j && z == k && l == w {
										continue
									}

									if space.Get(x, y, z, w) {
										count++
									}
								}
							}
						}
					}

					switch space.Get(i, j, k, l) {
					case true:
						if count == 2 || count == 3 {
							newSpace.Set(i, j, k, l, true)
						}
					case false:
						if count == 3 {
							newSpace.Set(i, j, k, l, true)
						}
					}
				}
			}
		}
	}

	return newSpace
}
