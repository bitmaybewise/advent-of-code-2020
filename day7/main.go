/*
--- Day 7: Handy Haversacks ---

You land at the regional airport in time for your next flight. In fact, it looks like you'll even have time to grab some food: all flights are currently delayed due to issues in luggage processing.

Due to recent aviation regulations, many rules (your puzzle input) are being enforced about bags and their contents; bags must be color-coded and must contain specific quantities of other color-coded bags. Apparently, nobody responsible for these regulations considered how long they would take to enforce!

For example, consider the following rules:

light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.

These rules specify the required contents for 9 bag types. In this example, every faded blue bag is empty, every vibrant plum bag contains 11 bags (5 faded blue and 6 dotted black), and so on.

You have a shiny gold bag. If you wanted to carry it in at least one other bag, how many different bag colors would be valid for the outermost bag? (In other words: how many colors can, eventually, contain at least one shiny gold bag?)

In the above rules, the following options would be available to you:

    A bright white bag, which can hold your shiny gold bag directly.
    A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
    A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
    A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.

So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.

How many bag colors can eventually contain at least one shiny gold bag? (The list of rules is quite long; make sure you get all of it.)

Your puzzle answer was 378.

The first half of this puzzle is complete! It provides one gold star: *
--- Part Two ---

It's getting pretty expensive to fly these days - not because of ticket prices, but because of the ridiculous number of bags you need to buy!

Consider again your shiny gold bag and the rules from the above example:

    faded blue bags contain 0 other bags.
    dotted black bags contain 0 other bags.
    vibrant plum bags contain 11 other bags: 5 faded blue bags and 6 dotted black bags.
    dark olive bags contain 7 other bags: 3 faded blue bags and 4 dotted black bags.

So, a single shiny gold bag must contain 1 dark olive bag (and the 7 bags within it) plus 2 vibrant plum bags (and the 11 bags within each of those): 1 + 1*7 + 2 + 2*11 = 32 bags!

Of course, the actual rules have a small chance of going several levels deeper than this example; be sure to count all of the bags, even if the nesting becomes topologically impractical!

Here's another example:

shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.

In this example, a single shiny gold bag must contain 126 other bags.

How many individual bags are required inside your single shiny gold bag?

27526

That's the right answer! You are one gold star closer to saving your vacation.

You have completed Day 7!
*/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hlmerscher/advent-of-code-2020/helper"
)

func main() {
	content := helper.ReadLinesFromInput("day7/input")

	fmt.Printf("Day 7-1: %d\n", answer1(content))
	fmt.Printf("Day 7-2: %d\n", answer2(content))
}

type Bag struct {
	name     string
	children Set
	parents  Set
}

func parseBag(s string) (string, int) {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "bags", "", -1)
	s = strings.Replace(s, "bag", "", -1)
	s = strings.Trim(s, " ")
	values := strings.Split(s, " ")
	var containedCapacity int
	if len(values) > 2 {
		containedCapacity, _ = strconv.Atoi(values[0])
		values = values[1:]
	}
	name := strings.Join(values, " ")
	return name, containedCapacity
}

func NewBag(name string) *Bag {
	return &Bag{name: name, parents: NewSet(), children: NewSet()}
}

type Set struct {
	_data map[string]bool
}

func NewSet() Set {
	return Set{_data: map[string]bool{}}
}

func (s *Set) Add(i string) {
	s._data[i] = true
}

func (s Set) Size() int {
	return len(s._data)
}

func answer1(s []string) int {
	bags := make(map[string]*Bag)

	for _, line := range s {
		values := strings.Split(line, "contain")
		contains, contained := strings.Trim(values[0], ""), strings.Trim(values[1], "")
		parentName, _ := parseBag(contains)

		for _, v := range strings.Split(contained, ",") {
			name, _ := parseBag(v)
			bag, found := bags[name]
			if !found {
				bag = NewBag(name)
			}
			bag.parents.Add(parentName)
			bags[name] = bag
		}
	}

	parentsOfParents := NewSet()
	var f func(s string)
	f = func(s string) {
		b, ok := bags[s]
		if !ok {
			return
		}
		for k := range b.parents._data {
			parentsOfParents.Add(k)
			f(k)
		}
	}
	f("shiny gold")

	return parentsOfParents.Size()
}

type BagCapacity struct {
	bag              *Bag
	childrenCapacity map[string]int
}

func NewBagCapacity(s string) *BagCapacity {
	return &BagCapacity{bag: NewBag(s), childrenCapacity: make(map[string]int)}
}

func (bc *BagCapacity) Add(s string, c int) {
	bc.bag.children.Add(s)
	bc.childrenCapacity[s] = c
}

func answer2(s []string) (count int) {
	b := make(map[string]*BagCapacity)

	for _, line := range s {
		values := strings.Split(line, "contain")
		contains, contained := strings.Trim(values[0], ""), strings.Trim(values[1], "")
		parentName, _ := parseBag(contains)
		parent, found := b[parentName]
		if !found {
			parent = NewBagCapacity(parentName)
		}
		b[parentName] = parent

		for _, v := range strings.Split(contained, ",") {
			childrenName, capacity := parseBag(v)
			parent.Add(childrenName, capacity)
		}
	}

	var f func(s string, m int)
	f = func(s string, multiplier int) {
		bc, found := b[s]
		if !found {
			return
		}
		for k, v := range bc.childrenCapacity {
			count += v * multiplier
			f(k, v*multiplier)
		}
	}
	f("shiny gold", 1)

	return count
}
