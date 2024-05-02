with open("input.txt") as f:
    lines: list[list[str]] = [list(line.strip()) for line in f.read().strip().split("\n")]
    # nums: list[int] = list(map(int, lines))

# Wrote for the input
STEPS = 26501365

GRID_WIDTH = 131

ABOVE = (-1, 65)
LEFT = (65, -1)
RIGHT = (65, 131)
BELOW = (131, 56)

GRID_RADIUS = (STEPS - (GRID_WIDTH - 1) // 2) // GRID_WIDTH - 1
LAST_LEG = STEPS - (GRID_RADIUS * GRID_WIDTH + (GRID_WIDTH - 1) // 2)
FILLED_GRIDS = (2 * GRID_RADIUS * (GRID_RADIUS + 1)) + 1

start_i, start_j = 0, 0

found = False
for i, line in enumerate(lines):
    for j, c in enumerate(line):
        if c == "S":
            start_i = i
            start_j = j
            found = True
            break

    if found:
        break

pos = {(start_i, start_j)}
next_pos = set()

odd_set = set()
even_set = set()

odd = False
# First run, determines how many are in the grid after an even or odd number of steps
while pos:
    this_set: set = even_set if odd else odd_set
    odd = not odd
    for (i, j) in pos:
        for (di, dj) in [(0, -1), (-1, 0), (1, 0), (0, 1)]:
            ni, nj = i + di, j + dj
            if not (0 <= ni < len(lines) and 0 <= nj < len(lines[0])):
                continue
            nc = lines[ni][nj]
            if nc == "#":
                continue
            nt = (ni, nj)

            if nt in this_set:
                continue

            this_set.add(nt)
            next_pos.add(nt)

    pos = next_pos
    next_pos = set()

# Total within all the filled grids
total = ((GRID_RADIUS + 1) // 2 * 2) ** 2 * len(even_set) + (GRID_RADIUS // 2 * 2 + 1) ** 2 * len(odd_set)

possible_starts = [ABOVE, BELOW, LEFT, RIGHT]

spr: dict[tuple, dict[int, dict[int, set]]] = {}

for start in possible_starts:
    start_i, start_j = start

    pos = {(start_i, start_j)}
    next_pos = set()

    for step in range(GRID_WIDTH):

        for (i, j) in pos:
            for (di, dj) in [(0, -1), (-1, 0), (1, 0), (0, 1)]:
                ni, nj = i + di, j + dj
                ri, rj = ni % GRID_WIDTH, nj % GRID_WIDTH
                nc = lines[ri][rj]
                if nc == "#":
                    continue
                next_pos.add((ni, nj))

        pos = next_pos
        next_pos = set()

    spr[start] = {}
    for p in pos:
        ni, nj = p
        ri, rdi = ni % GRID_WIDTH, ni // GRID_WIDTH
        rj, rdj = nj % GRID_WIDTH, nj // GRID_WIDTH

        if rdi not in spr[start]:
            spr[start][rdi] = {}
        if rdj not in spr[start][rdi]:
            spr[start][rdi][rdj] = set()
        spr[start][rdi][rdj].add((ri, rj))

    total += len(spr[start][0][0])

for corner in [
    (BELOW, LEFT), (BELOW, RIGHT), (ABOVE, RIGHT), (ABOVE, LEFT)
]:
    vs, hs = corner
    vss, hss = spr[vs][0][0], spr[hs][0][0]
    vos, hos = spr[vs][0][-1 if hs is RIGHT else 1], spr[hs][-1 if vs is BELOW else 1][0]

    total += (GRID_RADIUS * len(vss | hss)) + ((GRID_RADIUS + 1) * len(vos | hos))

print(total)