\--- Day 24: Blizzard Basin ---
----------

With everything replanted for next year (and with elephants and monkeys to tend the grove), you and the Elves leave for the extraction point.

Partway up the mountain that shields the grove is a flat, open area that serves as the extraction point. It's a bit of a climb, but nothing the expedition can't handle.

At least, that would normally be true; now that the mountain is covered in snow, things have become more difficult than the Elves are used to.

As the expedition reaches a valley that must be traversed to reach the extraction site, you find that strong, turbulent winds are pushing small *blizzards* of snow and sharp ice around the valley. It's a good thing everyone packed warm clothes! To make it across safely, you'll need to find a way to avoid them.

Fortunately, it's easy to see all of this from the entrance to the valley, so you make a map of the valley and the blizzards (your puzzle input). For example:

```
#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#

```

The walls of the valley are drawn as `#`; everything else is ground. Clear ground - where there is currently no blizzard - is drawn as `.`. Otherwise, blizzards are drawn with an arrow indicating their direction of motion: up (`^`), down (`v`), left (`<`), or right (`>`).

The above map includes two blizzards, one moving right (`>`) and one moving down (`v`). In one minute, each blizzard moves one position in the direction it is pointing:

```
#.#####
#.....#
#.>...#
#.....#
#.....#
#...v.#
#####.#

```

Due to conservation of blizzard energy, as a blizzard reaches the wall of the valley, a new blizzard forms on the opposite side of the valley moving in the same direction. After another minute, the bottom downward-moving blizzard has been replaced with a new downward-moving blizzard at the top of the valley instead:

```
#.#####
#...v.#
#..>..#
#.....#
#.....#
#.....#
#####.#

```

Because blizzards are made of tiny snowflakes, they pass right through each other. After another minute, both blizzards temporarily occupy the same position, marked `2`:

```
#.#####
#.....#
#...2.#
#.....#
#.....#
#.....#
#####.#

```

After another minute, the situation resolves itself, giving each blizzard back its personal space:

```
#.#####
#.....#
#....>#
#...v.#
#.....#
#.....#
#####.#

```

Finally, after yet another minute, the rightward-facing blizzard on the right is replaced with a new one on the left facing the same direction:

```
#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#

```

This process repeats at least as long as you are observing it, but probably forever.

Here is a more complex example:

```
#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#

```

Your expedition begins in the only non-wall position in the top row and needs to reach the only non-wall position in the bottom row. On each minute, you can *move* up, down, left, or right, or you can *wait* in place. You and the blizzards act *simultaneously*, and you cannot share a position with a blizzard.

In the above example, the fastest way to reach your goal requires `*18*` steps. Drawing the position of the expedition as `E`, one way to achieve this is:

```
Initial state:
#E######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#

Minute 1, move down:
#.######
#E>3.<.#
#<..<<.#
#>2.22.#
#>v..^<#
######.#

Minute 2, move down:
#.######
#.2>2..#
#E^22^<#
#.>2.^>#
#.>..<.#
######.#

Minute 3, wait:
#.######
#<^<22.#
#E2<.2.#
#><2>..#
#..><..#
######.#

Minute 4, move up:
#.######
#E<..22#
#<<.<..#
#<2.>>.#
#.^22^.#
######.#

Minute 5, move right:
#.######
#2Ev.<>#
#<.<..<#
#.^>^22#
#.2..2.#
######.#

Minute 6, move right:
#.######
#>2E<.<#
#.2v^2<#
#>..>2>#
#<....>#
######.#

Minute 7, move down:
#.######
#.22^2.#
#<vE<2.#
#>>v<>.#
#>....<#
######.#

Minute 8, move left:
#.######
#.<>2^.#
#.E<<.<#
#.22..>#
#.2v^2.#
######.#

Minute 9, move up:
#.######
#<E2>>.#
#.<<.<.#
#>2>2^.#
#.v><^.#
######.#

Minute 10, move right:
#.######
#.2E.>2#
#<2v2^.#
#<>.>2.#
#..<>..#
######.#

Minute 11, wait:
#.######
#2^E^2>#
#<v<.^<#
#..2.>2#
#.<..>.#
######.#

Minute 12, move down:
#.######
#>>.<^<#
#.<E.<<#
#>v.><>#
#<^v^^>#
######.#

Minute 13, move down:
#.######
#.>3.<.#
#<..<<.#
#>2E22.#
#>v..^<#
######.#

Minute 14, move right:
#.######
#.2>2..#
#.^22^<#
#.>2E^>#
#.>..<.#
######.#

Minute 15, move right:
#.######
#<^<22.#
#.2<.2.#
#><2>E.#
#..><..#
######.#

Minute 16, move right:
#.######
#.<..22#
#<<.<..#
#<2.>>E#
#.^22^.#
######.#

Minute 17, move down:
#.######
#2.v.<>#
#<.<..<#
#.^>^22#
#.2..2E#
######.#

Minute 18, move down:
#.######
#>2.<.<#
#.2v^2<#
#>..>2>#
#<....>#
######E#

```

*What is the fewest number of minutes required to avoid the blizzards and reach the goal?*

To begin, [get your puzzle input](24/input).

Answer:

You can also [Shareon [Twitter](https://twitter.com/intent/tweet?text=%22Blizzard+Basin%22+%2D+Day+24+%2D+Advent+of+Code+2022&url=https%3A%2F%2Fadventofcode%2Ecom%2F2022%2Fday%2F24&related=ericwastl&hashtags=AdventOfCode) [Mastodon](javascript:void(0);)] this puzzle.