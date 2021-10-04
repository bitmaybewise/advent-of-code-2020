package main

import (
	"testing"
)

func TestAnswer1(t *testing.T) {
	instructions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	result := answer1(instructions)
	if result != 25 {
		t.Errorf("Manhattan distance should be %d, got %d", 25, result)
	}
}

func TestAnswer2(t *testing.T) {
	instructions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	result := answer2(instructions)
	if result != 286 {
		t.Errorf("Manhattan distance should be %d, got %d", 286, result)
	}
}
