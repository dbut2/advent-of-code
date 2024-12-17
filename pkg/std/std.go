package std

import (
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/strings"
)

type Grid = space.Grid[byte]

var Abs = math.Abs
var Fill = lists.Fill
var Filter = lists.Filter
var GCD = math.GCD
var Int = sti.Int
var Ints = strings.Ints
var LCM = math.LCM
var Pow = math.Pow
var Product = math.Product
var Reduce = lists.Reduce
var Sign = math.Sign
var Sum = math.Sum
