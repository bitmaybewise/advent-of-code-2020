package main

import (
	"testing"
)

func TestBoardingPassRow(t *testing.T) {
	tests := []struct {
		b              boardingPass
		expectedResult int
	}{
		{b: boardingPass("FBFBBFFRLR"), expectedResult: 44},
		{b: boardingPass("BFFFBBFRRR"), expectedResult: 70},
		{b: boardingPass("FFFBBBFRRR"), expectedResult: 14},
		{b: boardingPass("BBFFBBFRLL"), expectedResult: 102},
	}

	for _, test := range tests {
		if result := test.b.row(); result != test.expectedResult {
			t.Errorf("For boarding pass %s expected row is %d, got %d", test.b, test.expectedResult, result)
		}
	}
}

func TestBoardingPassColumn(t *testing.T) {
	tests := []struct {
		b              boardingPass
		expectedResult int
	}{
		{b: boardingPass("FBFBBFFRLR"), expectedResult: 5},
		{b: boardingPass("BFFFBBFRRR"), expectedResult: 7},
		{b: boardingPass("FFFBBBFRRR"), expectedResult: 7},
		{b: boardingPass("BBFFBBFRLL"), expectedResult: 4},
	}

	for _, test := range tests {
		if result := test.b.column(); result != test.expectedResult {
			t.Errorf("For boarding pass %s expected column is %d, got %d", test.b, test.expectedResult, result)
		}
	}
}

func TestBoardingPassSeatId(t *testing.T) {
	tests := []struct {
		b              boardingPass
		expectedResult int
	}{
		{b: boardingPass("FBFBBFFRLR"), expectedResult: 357},
		{b: boardingPass("BFFFBBFRRR"), expectedResult: 567},
		{b: boardingPass("FFFBBBFRRR"), expectedResult: 119},
		{b: boardingPass("BBFFBBFRLL"), expectedResult: 820},
	}

	for _, test := range tests {
		if result := test.b.seatID(); result != test.expectedResult {
			t.Errorf("For boarding pass %s expected column is %d, got %d", test.b, test.expectedResult, result)
		}
	}
}

func TestAnswer1(t *testing.T) {
	b := []boardingPass{
		boardingPass("FBFBBFFRLR"),
		boardingPass("BFFFBBFRRR"),
		boardingPass("FFFBBBFRRR"),
		boardingPass("BBFFBBFRLL"),
	}
	if result := answer1(b); result != 820 {
		t.Errorf("Highest seat ID should be %d, got %d", 820, result)
	}
}
