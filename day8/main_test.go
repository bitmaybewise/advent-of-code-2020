package main

import (
	"testing"
)

func TestAnswer1(t *testing.T) {
	instructions := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	if result := answer1(instructions); result != 5 {
		t.Errorf("Accumulator should be %d, got %d", 5, result)
	}
}

func TestAnswer2(t *testing.T) {
	instructions := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	if result := answer2(instructions); result != 8 {
		t.Errorf("Accumulator should be %d, got %d", 8, result)
	}
}
