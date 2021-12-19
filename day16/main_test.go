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
