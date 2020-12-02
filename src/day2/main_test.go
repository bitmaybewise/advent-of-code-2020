package main

import "testing"

func TestIsValidOldJob(t *testing.T) {
	t.Run("When valid", func(t *testing.T) {
		a := password{min: 1, max: 3, letter: "a", value: "abcde"}
		b := password{min: 2, max: 9, letter: "c", value: "ccccccccc"}
		for _, p := range []password{a, b} {
			if !isValidOldJob(p) {
				t.Errorf("Password %s should be valid!", p.value)
			}
		}
	})
	t.Run("When invalid", func(t *testing.T) {
		a := password{min: 1, max: 3, letter: "b", value: "cdefg"}
		b := password{min: 2, max: 3, letter: "a", value: "acdefg"}
		for _, p := range []password{a, b} {
			if isValidOldJob(p) {
				t.Errorf("Password %s should be invalid!", p.value)
			}
		}
	})
}

func TestParsePassword(t *testing.T) {
	t.Run("With simple password", func(t *testing.T) {
		line := "1-3 a: abcde"
		p := parsePassword(line)
		if p.min != 1 {
			t.Errorf("Min should be %d, but got %d", 1, p.min)
		}
		if p.max != 3 {
			t.Errorf("Min should be %d, but got %d", 3, p.max)
		}
		if p.letter != "a" {
			t.Errorf("Letter should be %s, but got %s", "a", p.letter)
		}
		if p.value != "abcde" {
			t.Errorf("Value should be %s, but got %s", "a", p.value)
		}
	})
	t.Run("With bigger limits", func(t *testing.T) {
		line := "33-666 a: abcde"
		p := parsePassword(line)
		if p.min != 33 {
			t.Errorf("Min should be %d, but got %d", 33, p.min)
		}
		if p.max != 666 {
			t.Errorf("Min should be %d, but got %d", 666, p.max)
		}
		if p.letter != "a" {
			t.Errorf("Letter should be %s, but got %s", "a", p.letter)
		}
		if p.value != "abcde" {
			t.Errorf("Value should be %s, but got %s", "a", p.value)
		}
	})
}

func TestAnswer1(t *testing.T) {
	a := password{min: 1, max: 3, letter: "a", value: "abcde"}
	b := password{min: 1, max: 3, letter: "b", value: "cdefg"}
	c := password{min: 2, max: 9, letter: "c", value: "ccccccccc"}
	d := password{min: 14, max: 18, letter: "m", value: "mmmmmmmmmmmmmmmmmmmm"}
	p := []password{a, b, c, d}
	if result := answer1(p); result != 2 {
		t.Errorf("Expected valid passwords is %d, but got %d", 2, result)
	}
}

func TestAnswer2(t *testing.T) {
	a := password{min: 1, max: 3, letter: "a", value: "abcde"}
	b := password{min: 1, max: 3, letter: "c", value: "abcde"}
	c := password{min: 1, max: 3, letter: "b", value: "cdefg"}
	d := password{min: 2, max: 9, letter: "c", value: "ccccccccc"}
	p := []password{a, b, c, d}
	if result := answer2(p); result != 2 {
		t.Errorf("Expected valid passwords is %d, but got %d", 2, result)
	}
}

func TestIsValidToboggan(t *testing.T) {
	t.Run("When valid", func(t *testing.T) {
		a := password{min: 1, max: 3, letter: "a", value: "abcde"}
		b := password{min: 1, max: 3, letter: "c", value: "abcde"}
		for _, p := range []password{a, b} {
			if !isValidToboggan(p) {
				t.Errorf("Password %s should be valid!", p.value)
			}
		}
	})
	t.Run("When invalid", func(t *testing.T) {
		a := password{min: 1, max: 3, letter: "b", value: "cdefg"}
		b := password{min: 2, max: 9, letter: "c", value: "ccccccccc"}
		for _, p := range []password{a, b} {
			if isValidToboggan(p) {
				t.Errorf("Password %s should be invalid!", p.value)
			}
		}
	})
}
