\--- Day 21: Monkey Math ---
----------

The [monkeys](11) are back! You're worried they're going to try to steal your stuff again, but it seems like they're just holding their ground and making various monkey noises at you.

Eventually, one of the elephants realizes you don't speak monkey and comes over to interpret. As it turns out, they overheard you talking about trying to find the grove; they can show you a shortcut if you answer their *riddle*.

Each monkey is given a *job*: either to *yell a specific number* or to *yell the result of a math operation*. All of the number-yelling monkeys know their number from the start; however, the math operation monkeys need to wait for two other monkeys to yell a number, and those two other monkeys might *also* be waiting on other monkeys.

Your job is to *work out the number the monkey named `root` will yell* before the monkeys figure it out themselves.

For example:

```
root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32

```

Each line contains the name of a monkey, a colon, and then the job of that monkey:

* A lone number means the monkey's job is simply to yell that number.
* A job like `aaaa + bbbb` means the monkey waits for monkeys `aaaa` and `bbbb` to yell each of their numbers; the monkey then yells the sum of those two numbers.
* `aaaa - bbbb` means the monkey yells `aaaa`'s number minus `bbbb`'s number.
* Job `aaaa * bbbb` will yell `aaaa`'s number multiplied by `bbbb`'s number.
* Job `aaaa / bbbb` will yell `aaaa`'s number divided by `bbbb`'s number.

So, in the above example, monkey `drzm` has to wait for monkeys `hmdt` and `zczc` to yell their numbers. Fortunately, both `hmdt` and `zczc` have jobs that involve simply yelling a single number, so they do this immediately: `32` and `2`. Monkey `drzm` can then yell its number by finding `32` minus `2`: `*30*`.

Then, monkey `sjmn` has one of its numbers (`30`, from monkey `drzm`), and already has its other number, `5`, from `dbpl`. This allows it to yell its own number by finding `30` multiplied by `5`: `*150*`.

This process continues until `root` yells a number: `*152*`.

However, your actual situation involves considerably more monkeys. *What number will the monkey named `root` yell?*

To begin, [get your puzzle input](21/input).

Answer:

You can also [Shareon [Twitter](https://twitter.com/intent/tweet?text=%22Monkey+Math%22+%2D+Day+21+%2D+Advent+of+Code+2022&url=https%3A%2F%2Fadventofcode%2Ecom%2F2022%2Fday%2F21&related=ericwastl&hashtags=AdventOfCode) [Mastodon](javascript:void(0);)] this puzzle.