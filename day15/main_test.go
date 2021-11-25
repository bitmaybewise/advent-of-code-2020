package main

import (
	"fmt"
	"testing"
)

func Test_Day15_Answer(t *testing.T) {
	table := []struct {
		input           []int
		expected, until int
	}{
		// part 1
		{input: []int{0, 3, 6}, expected: 436, until: 2020},
		{input: []int{0, 13, 1, 8, 6, 15}, expected: 1618, until: 2020},
		// {input: []int{1, 3, 2}, expected: 1, until: 2020},
		// {input: []int{2, 1, 3}, expected: 10},
		// {input: []int{1, 2, 3}, expected: 27},
		// {input: []int{2, 3, 1}, expected: 78},
		// {input: []int{3, 2, 1}, expected: 438},
		// {input: []int{3, 1, 2}, expected: 1836},

		// part 2
		{input: []int{0, 3, 6}, expected: 175594, until: 30_000_000},
		{input: []int{0, 13, 1, 8, 6, 15}, expected: 548531, until: 30_000_000},
		// {input: []int{1, 3, 2}, expected: 2578, until: 30_000_000},
		// {input: []int{2, 1, 3}, expected: 3_544_142, until: 30_000_000},
		// {input: []int{1, 2, 3}, expected: 261_214, until: 30_000_000},
		// {input: []int{2, 3, 1}, expected: 6_895_259, until: 30_000_000},
		// {input: []int{3, 2, 1}, expected: 18, until: 30_000_000},
		// {input: []int{3, 1, 2}, expected: 362, until: 30_000_000},
	}

	for _, value := range table {
		t.Run(fmt.Sprintf("given %v", value.input), func(t *testing.T) {
			result := answer(value.input, value.until)
			if result != value.expected {
				t.Errorf("%dth number should be %d, got %d", value.until, value.expected, result)
			}
		})
	}
}
