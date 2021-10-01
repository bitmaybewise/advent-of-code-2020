package main

import (
	"strings"
	"testing"
)

func TestAnswer1(t *testing.T) {
	grid := grid{
		strings.Split("L.LL.LL.LL", ""),
		strings.Split("LLLLLLL.LL", ""),
		strings.Split("L.L.L..L..", ""),
		strings.Split("LLLL.LL.LL", ""),
		strings.Split("L.LL.LL.LL", ""),
		strings.Split("L.LLLLL.LL", ""),
		strings.Split("..L.L.....", ""),
		strings.Split("LLLLLLLLLL", ""),
		strings.Split("L.LLLLLL.L", ""),
		strings.Split("L.LLLLL.LL", ""),
	}
	result := answer1(grid)
	expected := 37
	if result != expected {
		t.Errorf("Seats occupied should be %d, got %d", expected, result)
	}
}

func TestAnswer2(t *testing.T) {
	grid := grid{
		strings.Split("L.LL.LL.LL", ""),
		strings.Split("LLLLLLL.LL", ""),
		strings.Split("L.L.L..L..", ""),
		strings.Split("LLLL.LL.LL", ""),
		strings.Split("L.LL.LL.LL", ""),
		strings.Split("L.LLLLL.LL", ""),
		strings.Split("..L.L.....", ""),
		strings.Split("LLLLLLLLLL", ""),
		strings.Split("L.LLLLLL.L", ""),
		strings.Split("L.LLLLL.LL", ""),
	}
	result := answer2(grid)
	expected := 26
	if result != expected {
		t.Errorf("Seats occupied should be %d, got %d", expected, result)
	}
}

func TestFarAdjacentsOccupied(t *testing.T) {
	g := grid{
		strings.Split(".......#.", ""),
		strings.Split("...#.....", ""),
		strings.Split(".#.......", ""),
		strings.Split(".........", ""),
		strings.Split("..#L....#", ""),
		strings.Split("....#....", ""),
		strings.Split(".........", ""),
		strings.Split("#........", ""),
		strings.Split("...#.....", ""),
	}
	result := farAdjacentsOccupied(g, 4, 3, 9, 9)
	expected := 8
	if result != expected {
		t.Errorf("Seats occupied should be %d, got %d", expected, result)
	}

	g = grid{
		strings.Split(".##.##.", ""),
		strings.Split("#.#.#.#", ""),
		strings.Split("##...##", ""),
		strings.Split("...L...", ""),
		strings.Split("##...##", ""),
		strings.Split("#.#.#.#", ""),
		strings.Split(".##.##.", ""),
	}
	result = farAdjacentsOccupied(g, 3, 3, 7, 7)
	expected = 0
	if result != expected {
		t.Errorf("Seats occupied should be %d, got %d", expected, result)
	}
}
