package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		set := NewSet()
		set.Add('a')
		set.Add('a')
		set.Add('b')
		set.Add('b')
		set.Add('c')
		if size := set.Size(); size != 3 {
			t.Errorf("Set should have size %d, got %d", 3, size)
		}
	})
	t.Run("Intersection", func(t *testing.T) {
		setA := NewSet()
		setA.Add('a')
		setA.Add('b')
		setB := NewSet()
		setB.Add('c')
		setB.Add('a')
		setC := setA.Intersection(setB)
		if size := setC.Size(); size != 1 {
			t.Errorf("Set should have size %d, got %d", 1, size)
		}
	})
}

func TestAnswer1(t *testing.T) {
	s := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
	if result := answer1(s); result != 11 {
		t.Errorf("Number of answers should be %d, got %d", 11, result)
	}
}

func TestAnswer2(t *testing.T) {
	s := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
	if result := answer2(s); result != 6 {
		t.Errorf("Number of answers should be %d, got %d", 6, result)
	}
}
