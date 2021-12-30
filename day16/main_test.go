package main

import (
	"testing"
)

func Test_Day16_Answer1(t *testing.T) {
	content := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",

		"your ticket:",
		"7,1,14",

		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}
	result := answer1(content, 3, 7, 10)
	expected := 71
	if result != expected {
		t.Errorf("Error rate should be %d, got %d", expected, result)
	}
}

func Test_Day16_What(t *testing.T) {
	content := []string{
		"class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"seat: 0-13 or 16-19",

		"your ticket:",
		"11,12,13",

		"nearby tickets:",
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}
	t.Run("returns which row is what", func(t *testing.T) {
		result := what(content, 3, 6, 9)
		if result["row"] != 0 {
			t.Errorf("Row 0 should be row, got %d", result["row"])
		}
		if result["class"] != 1 {
			t.Errorf("Row 1 should be class, got %d", result["class"])
		}
		if result["seat"] != 2 {
			t.Errorf("Row 2 should be class, got %d", result["seat"])
		}
		// if result[0] != "row" {
		// 	t.Errorf("Row 0 should be row, got %s", result[0])
		// }
		// if result[1] != "class" {
		// 	t.Errorf("Row 1 should be class, got %s", result[1])
		// }
		// if result[2] != "seat" {
		// 	t.Errorf("Row 2 should be class, got %s", result[2])
		// }
	})
}

func Test_Day16_Answer2(t *testing.T) {
	content := []string{
		"departure class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"departure seat: 0-13 or 16-19",

		"your ticket:",
		"11,12,13",

		"nearby tickets:",
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}
	result := answer2(content, 3, 6, 9, 4)
	expected := 156
	if result != expected {
		t.Errorf("Departure sum should be %d, got %d", expected, result)
	}
}
