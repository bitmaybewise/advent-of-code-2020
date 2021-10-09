package main

import (
	"testing"
)

func TestAnswer1(t *testing.T) {
	earliest := 939
	buses := []string{"7", "13", "x", "x", "59", "x", "31", "19"}
	result := answer1(earliest, buses)
	if result != 295 {
		t.Errorf("The bus ID multiple by the # of minutes should be %d, got %d", 295, result)
	}
}

func TestAnswer2(t *testing.T) {
	table := []struct {
		input  []string
		output int
	}{
		{[]string{"7", "13", "x", "x", "59", "x", "31", "19"}, 1068781},
		{[]string{"17", "x", "13", "19"}, 3417},
		{[]string{"67", "7", "59", "61"}, 754018},
		{[]string{"67", "x", "7", "59", "61"}, 779210},
		{[]string{"67", "7", "x", "59", "61"}, 1261476},
		{[]string{"1789", "37", "47", "1889"}, 1202161486},
	}
	for _, value := range table {
		result := answer2(value.input)
		if result != value.output {
			t.Errorf("The earliest timestamp should be %d, got %d", value.output, result)
		}
	}
}
