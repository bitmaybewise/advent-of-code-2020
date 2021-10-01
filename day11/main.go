/*
--- Day 11: Seating System ---

Your plane lands with plenty of time to spare. The final leg of your journey is a ferry that goes directly to the tropical island where you can finally start your vacation. As you reach the waiting area to board the ferry, you realize you're so early, nobody else has even arrived yet!

By modeling the process people use to choose (or abandon) their seat in the waiting area, you're pretty sure you can predict the best place to sit. You make a quick map of the seat layout (your puzzle input).

The seat layout fits neatly on a grid. Each position is either floor (.), an empty seat (L), or an occupied seat (#). For example, the initial seat layout might look like this:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL

Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely predictable and always follow a simple set of rules. All decisions are based on the number of occupied seats adjacent to a given seat (one of the eight positions immediately up, down, left, right, or diagonal from the seat). The following rules are applied to every seat simultaneously:

    If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
    If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
    Otherwise, the seat's state does not change.

Floor (.) never changes; seats don't move, and nobody sits on the floor.

After one round of these rules, every seat in the example layout becomes occupied:

#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##

After a second round, the seats with four or more occupied adjacent seats become empty again:

#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##

This process continues for three more rounds:

#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##

#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##

#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##

At this point, something interesting happens: the chaos stabilizes and further applications of these rules cause no seats to change state! Once people stop moving around, you count 37 occupied seats.

Simulate your seating area by applying the seating rules repeatedly until no seats change state. How many seats end up occupied?

2418

That's the right answer! You are one gold star closer to saving your

--- Part Two ---

As soon as people start to arrive, you realize your mistake. People don't just care about adjacent seats - they care about the first seat they can see in each of those eight directions!

Now, instead of considering just the eight immediately adjacent seats, consider the first seat in each of those eight directions. For example, the empty seat below would see eight occupied seats:

.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....

The leftmost empty seat below would only see one empty seat, but cannot see any of the occupied ones:

.............
.L.L.#.#.#.#.
.............

The empty seat below would see no occupied seats:

.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.

Also, people seem to be more tolerant than you expected: it now takes five or more visible occupied seats for an occupied seat to become empty (rather than four or more from the previous rules). The other rules still apply: empty seats that see no occupied seats become occupied, seats matching no rule don't change, and floor never changes.

Given the same starting layout as above, these new rules cause the seating area to shift around as follows:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL

#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##

#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#

#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#

#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#

#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#

#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#

Again, at this point, people stop shifting around and the seating area reaches equilibrium. Once this occurs, you count 26 occupied seats.

Given the new visibility method and the rule change for occupied seats becoming empty, once equilibrium is reached, how many seats end up occupied?

2144

That's the right answer! You are one gold star closer to saving your vacation.

You have completed Day 11!
*/
package main

import (
	"fmt"
	"strings"

	"github.com/hlmerscher/advent-of-code-2020/helper"
)

type grid [][]string

const (
	Floor        = "."
	EmptySeat    = "L"
	OccupiedSeat = "#"
)

func main() {
	content := helper.ReadLinesFromInput("day11/input")
	g := make(grid, 0)
	for _, v := range content {
		if v == "" {
			break
		}
		g = append(g, strings.Split(v, ""))
	}
	fmt.Printf("Day 11-1: %d\n", answer1(g))
	fmt.Printf("Day 11-2: %d\n", answer2(g))
}

func answer1(g grid) int {
	cols, rows, moves := len(g), len(g[0]), 0
	gcopy := clone(g, rows, cols)

	for y := 0; y < cols; y++ {
		for x := 0; x < rows; x++ {
			if val := checkSeat(g, y, x, cols, rows); val != g[y][x] {
				gcopy[y][x] = val
				moves++
			}
		}
	}
	if moves == 0 {
		return countOccupied(g, cols, rows)
	}

	return answer1(gcopy)
}

func answer2(g grid) int {
	cols, rows, moves := len(g), len(g[0]), 0
	gcopy := clone(g, rows, cols)

	for y := 0; y < cols; y++ {
		for x := 0; x < rows; x++ {
			val := checkSeat2(g, y, x, cols, rows)
			if val != g[y][x] {
				gcopy[y][x] = val
				moves++
			}
		}
	}
	if moves == 0 {
		return countOccupied(g, cols, rows)
	}

	return answer2(gcopy)
}

func clone(g grid, rows, cols int) grid {
	gridCopy := make(grid, cols)
	for y := 0; y < cols; y++ {
		row := make([]string, rows)
		for x := 0; x < rows; x++ {
			row[x] = g[y][x]
		}
		gridCopy[y] = row
	}
	return gridCopy
}

func countOccupied(g grid, cols, rows int) (count int) {
	for y := 0; y < cols; y++ {
		for x := 0; x < rows; x++ {
			count += occupied(g, y, x, cols, rows)
		}
	}
	return
}

func checkSeat(g grid, y, x, cols, rows int) string {
	val := g[y][x]
	adjacents := immediatelyAdjacentsOccupied(g, y, x, cols, rows)
	if val == EmptySeat && adjacents == 0 {
		return OccupiedSeat
	}
	if val == OccupiedSeat && adjacents > 3 {
		return EmptySeat
	}
	return val
}

func checkSeat2(g grid, y, x, cols, rows int) string {
	val := g[y][x]
	if val == Floor {
		return val
	}
	visible := farAdjacentsOccupied(g, y, x, cols, rows)
	if val == EmptySeat && visible == 0 {
		return OccupiedSeat
	}
	if val == OccupiedSeat && visible > 4 {
		return EmptySeat
	}
	return val
}

func immediatelyAdjacentsOccupied(g grid, y, x, cols, rows int) int {
	if g[y][x] == Floor {
		return 0
	}
	return occupied(g, y, x+1, cols, rows) + // right
		occupied(g, y+1, x+1, cols, rows) + // right bottom
		occupied(g, y+1, x, cols, rows) + // bottom
		occupied(g, y+1, x-1, cols, rows) + // left bottom
		occupied(g, y, x-1, cols, rows) + // left
		occupied(g, y-1, x-1, cols, rows) + // top left
		occupied(g, y-1, x, cols, rows) + // top
		occupied(g, y-1, x+1, cols, rows) // top right
}

func farAdjacentsOccupied(g grid, y, x, cols, rows int) int {
	var count int
	inc := func(y, x int) bool {
		if g[y][x] == EmptySeat {
			return true
		}
		val := occupied(g, y, x, cols, rows)
		if val == 1 {
			count += val
			return true
		}
		return false
	}
	// right
	for i := x + 1; i < rows; i++ {
		if inc(y, i) {
			break
		}
	}
	// right bottom
	for i, j := x+1, y+1; i < rows && j < cols; {
		if inc(j, i) {
			break
		}
		i++
		j++
	}
	// bottom
	for j := y + 1; j < cols; j++ {
		if inc(j, x) {
			break
		}
	}
	// bottom left
	for i, j := x-1, y+1; i >= 0 && j < cols; {
		if inc(j, i) {
			break
		}
		i--
		j++
	}
	// left
	for i := x - 1; i >= 0; i-- {
		if inc(y, i) {
			break
		}
	}
	// top left
	for i, j := x-1, y-1; i >= 0 && j >= 0; {
		if inc(j, i) {
			break
		}
		i--
		j--
	}
	// top
	for j := y - 1; j >= 0; j-- {
		if inc(j, x) {
			break
		}
	}
	// top right
	for i, j := x+1, y-1; i < rows && j >= 0; {
		if inc(j, i) {
			break
		}
		i++
		j--
	}
	return count
}

func occupied(g grid, y, x, cols, rows int) int {
	if x >= 0 && y >= 0 && x < rows && y < cols && g[y][x] == OccupiedSeat {
		return 1
	}
	return 0
}
