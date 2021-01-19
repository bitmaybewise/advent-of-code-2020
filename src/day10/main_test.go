package main

import (
	"testing"
)

func TestAnswer1(t *testing.T) {
	numbers := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}
	if result := answer1(numbers); result != 35 {
		t.Errorf("Jolt difference should be %d, got %d", 35, result)
	}
	numbers = []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	if result := answer1(numbers); result != 220 {
		t.Errorf("Jolt difference should be %d, got %d", 220, result)
	}
}

func TestAnswer2(t *testing.T) {
}
