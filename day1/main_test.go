package main

import "testing"

func TestAnswer1(t *testing.T) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	if n := answer1(values); n != 514579 {
		t.Errorf("Result %d, expected %d", n, 514579)
	}
}

func TestAnswer2(t *testing.T) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	if n := answer2(values); n != 241861950 {
		t.Errorf("Result %d, expected %d", n, 241861950)
	}
}
