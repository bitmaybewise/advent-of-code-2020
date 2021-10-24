package main

import (
	"testing"
)

func TestApplyBitmask1(t *testing.T) {
	table := []struct {
		mask            string
		input, expected int
	}{
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 11, 73},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 101, 101},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 0, 64},
	}
	for _, val := range table {
		result := applyBitmask1(val.mask, val.input)
		if result != val.expected {
			t.Errorf("Mask applied should result into %d, got %d", val.expected, result)
		}
	}
}

func TestApplyBitmask2(t *testing.T) {
	table := []struct {
		mask            string
		input, expected int
	}{
		{"000000000000000000000000000000X1001X", 42, 59},
		{"00000000000000000000000000000000X0XX", 26, 27},
	}
	for _, val := range table {
		result := applyBitmask2(val.mask, val.input, nil)
		if result != val.expected {
			t.Errorf("Mask applied should result into %d, got %d", val.expected, result)
		}
	}
}

func TestAnswer1(t *testing.T) {
	input := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	result := answer1(input)
	expected := 165
	if result != expected {
		t.Errorf("Sum should be %d, got %d", expected, result)
	}
}

func TestAnswer2(t *testing.T) {
	input := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
	result := answer2(input)
	expected := 208
	if result != expected {
		t.Errorf("Sum should be %d, got %d", expected, result)
	}
}
