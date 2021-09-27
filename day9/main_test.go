package main

import (
	"testing"
)

func TestAnswer1(t *testing.T) {
	numbers := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	if result := answer1(numbers, 5); result != 127 {
		t.Errorf("Number should be %d, got %d", 127, result)
	}
}

func TestAnswer2(t *testing.T) {
	numbers := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	a := answer1(numbers, 5)

	if result := answer2(numbers, a); result != 62 {
		t.Errorf("Number should be %d, got %d", 62, result)
	}
}
