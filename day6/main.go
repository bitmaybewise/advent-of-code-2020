/*
--- Day 6: Custom Customs ---

As your flight approaches the regional airport where you'll switch to a much larger plane, customs declaration forms are distributed to the passengers.

The form asks a series of 26 yes-or-no questions marked a through z. All you need to do is identify the questions for which anyone in your group answers "yes". Since your group is just you, this doesn't take very long.

However, the person sitting next to you seems to be experiencing a language barrier and asks if you can help. For each of the people in their group, you write down the questions for which they answer "yes", one per line. For example:

abcx
abcy
abcz

In this group, there are 6 questions to which anyone answered "yes": a, b, c, x, y, and z. (Duplicate answers to the same question don't count extra; each question counts at most once.)

Another group asks for your help, then another, and eventually you've collected answers from every group on the plane (your puzzle input). Each group's answers are separated by a blank line, and within each group, each person's answers are on a single line. For example:

abc

a
b
c

ab
ac

a
a
a
a

b

This list represents answers from five groups:

    The first group contains one person who answered "yes" to 3 questions: a, b, and c.
    The second group contains three people; combined, they answered "yes" to 3 questions: a, b, and c.
    The third group contains two people; combined, they answered "yes" to 3 questions: a, b, and c.
    The fourth group contains four people; combined, they answered "yes" to only 1 question, a.
    The last group contains one person who answered "yes" to only 1 question, b.

In this example, the sum of these counts is 3 + 3 + 3 + 1 + 1 = 11.

For each group, count the number of questions to which anyone answered "yes". What is the sum of those counts?

Your puzzle answer was 6532.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---

As you finish the last group's customs declaration, you notice that you misread one word in the instructions:

You don't need to identify the questions to which anyone answered "yes"; you need to identify the questions to which everyone answered "yes"!

Using the same example as above:

abc

a
b
c

ab
ac

a
a
a
a

b

This list represents answers from five groups:

    In the first group, everyone (all 1 person) answered "yes" to 3 questions: a, b, and c.
    In the second group, there is no question to which everyone answered "yes".
    In the third group, everyone answered yes to only 1 question, a. Since some people did not answer "yes" to b or c, they don't count.
    In the fourth group, everyone answered yes to only 1 question, a.
    In the fifth group, everyone (all 1 person) answered "yes" to 1 question, b.

In this example, the sum of these counts is 3 + 0 + 1 + 1 + 1 = 6.

For each group, count the number of questions to which everyone answered "yes". What is the sum of those counts?

3427

That's the right answer! You are one gold star closer to saving your vacation.

You have completed Day 6!
*/
package main

import (
	"fmt"

	"github.com/hlmerscher/advent-of-code-2020/helper"
)

func main() {
	content := helper.ReadLinesFromInput("day6/input")

	fmt.Printf("Day 6-1: %d\n", answer1(content))
	fmt.Printf("Day 6-2: %d\n", answer2(content))
}

type Set struct {
	_data map[rune]bool
}

func NewSet() Set {
	return Set{_data: map[rune]bool{}}
}

func (s *Set) Add(i rune) {
	s._data[i] = true
}

func (s Set) Size() int {
	return len(s._data)
}

func (a Set) Intersection(b Set) Set {
	s := NewSet()
	for k, _ := range a._data {
		if _, found := b._data[k]; found {
			s.Add(k)
		}
	}
	return s
}

func answer1(s []string) (count int) {
	set := NewSet()
	for _, line := range s {
		if line == "" {
			count += set.Size()
			set = NewSet()
		}
		for _, char := range line {
			set.Add(char)
		}
	}
	count += set.Size()
	return
}

func intersection(g []Set) Set {
	set := g[0]
	for i := 1; i < len(g); i++ {
		set = set.Intersection(g[i])
	}
	return set
}

func answer2(s []string) (count int) {
	groups := make([]Set, 0)
	for _, line := range s {
		if line == "" {
			count += intersection(groups).Size()
			groups = make([]Set, 0)
			continue
		}
		set := NewSet()
		for _, char := range line {
			set.Add(char)
		}
		groups = append(groups, set)
	}
	count += intersection(groups).Size()
	return
}
