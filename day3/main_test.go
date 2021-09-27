package main

import (
	"testing"
)

func TestAnswer1(t *testing.T) {
	grid := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	if result := answer1(grid); result != 7 {
		t.Errorf("%d trees should have been found, got %d", 7, result)
	}
}

func TestAnswer2(t *testing.T) {
	grid := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	if result := answer2(grid); result != 336 {
		t.Errorf("%d trees should have been found, got %d", 336, result)
	}
}
