/*
--- Day 16: Ticket Translation ---

As you're walking to yet another connecting flight, you realize that one of the legs of your re-routed trip coming up is on a high-speed train. However, the train ticket you were given is in a language you don't understand. You should probably figure out what it says before you get to the train station after the next flight.

Unfortunately, you can't actually read the words on the ticket. You can, however, read the numbers, and so you figure out the fields these tickets must have and the valid ranges for values in those fields.

You collect the rules for ticket fields, the numbers on your ticket, and the numbers on other nearby tickets for the same train service (via the airport security cameras) together into a single document you can reference (your puzzle input).

The rules for ticket fields specify a list of fields that exist somewhere on the ticket and the valid ranges of values for each field. For example, a rule like class: 1-3 or 5-7 means that one of the fields in every ticket is named class and can be any value in the ranges 1-3 or 5-7 (inclusive, such that 3 and 5 are both valid in this field, but 4 is not).

Each ticket is represented by a single line of comma-separated values. The values are the numbers on the ticket in the order they appear; every ticket has the same format. For example, consider this ticket:

.--------------------------------------------------------.
| ????: 101    ?????: 102   ??????????: 103     ???: 104 |
|                                                        |
| ??: 301  ??: 302             ???????: 303      ??????? |
| ??: 401  ??: 402           ???? ????: 403    ????????? |
'--------------------------------------------------------'

Here, ? represents text in a language you don't understand. This ticket might be represented as 101,102,103,104,301,302,303,401,402,403; of course, the actual train tickets you're looking at are much more complicated. In any case, you've extracted just the numbers in such a way that the first number is always the same specific field, the second number is always a different specific field, and so on - you just don't know what each position actually means!

Start by determining which tickets are completely invalid; these are tickets that contain values which aren't valid for any field. Ignore your ticket for now.

For example, suppose you have the following notes:

class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12

It doesn't matter which position corresponds to which field; you can identify invalid nearby tickets by considering only whether tickets contain values that are not valid for any field. In this example, the values on the first nearby ticket are all valid for at least one field. This is not true of the other three nearby tickets: the values 4, 55, and 12 are are not valid for any field. Adding together all of the invalid values produces your ticket scanning error rate: 4 + 55 + 12 = 71.

Consider the validity of the nearby tickets you scanned. What is your ticket scanning error rate?

25788

That's the right answer! You are one gold star closer to saving your vacation.

--- Part Two ---

Now that you've identified which tickets contain invalid values, discard those tickets entirely. Use the remaining valid tickets to determine which field is which.

Using the valid ranges for each field, determine what order the fields appear on the tickets. The order is consistent between all tickets: if seat is the third field, it is the third field on every ticket, including your ticket.

For example, suppose you have the following notes:

class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9

Based on the nearby tickets in the above example, the first position must be row, the second position must be class, and the third position must be seat; you can conclude that in your ticket, class is 12, row is 11, and seat is 13.

Once you work out which field is which, look for the six fields on your ticket that start with the word departure. What do you get if you multiply those six values together?

*/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hlmerscher/advent-of-code-2020/helper"
)

func main() {
	content := helper.ReadLinesFromInput("day16/input")
	fmt.Printf("Day 16-1: %d\n", answer1(content, 20, 25, 262))
	fmt.Printf("Day 16-2: %d\n", answer2(content, 20, 25, 262, 22))
}

func answer1(content []string, headerSize, nearbyTicketsStart, nearbyTicketsEnd int) int {
	_, valids := build(content, headerSize)
	_, errorRate := nearby(content, valids, nearbyTicketsStart, nearbyTicketsEnd)
	return errorRate
}

func answer2(content []string, headerSize, nearbyTicketsStart, nearbyTicketsEnd, yourTickets int) int {
	mytickets := make([]int, 0)
	values := strings.Split(content[yourTickets], ",")
	for _, val := range values {
		n := toInt(val)
		mytickets = append(mytickets, n)
	}
	mapping := what(content, headerSize, nearbyTicketsStart, nearbyTicketsEnd)
	sum := 1
	for field, i := range mapping {
		if strings.Contains(field, "departure") {
			sum *= mytickets[i]
		}
		// if strings.Contains(field, "departure") {
		// 	sum *= mytickets[i]
		// }
	}
	return sum
}

type Set struct {
	_data map[string]bool
}

func NewSet(fields map[string][]int) Set {
	set := Set{_data: map[string]bool{}}
	for k := range fields {
		set.Add(k)
	}
	return set
}

func (s *Set) Add(i string) {
	s._data[i] = true
}

func (s *Set) Del(i string) {
	for k := range s._data {
		if k == i {
			delete(s._data, k)
		}
	}
}

func (s Set) Size() int {
	return len(s._data)
}

func (a Set) Intersection(b Set) Set {
	s := NewSet(nil)
	for k := range a._data {
		if _, found := b._data[k]; found {
			s.Add(k)
		}
	}
	return s
}

func (a Set) First() string {
	for k := range a._data {
		return k
	}
	return ""
}

func contains(items []int, n int) bool {
	for _, val := range items {
		if val == n {
			return true
		}
	}
	return false
}

func toInt(val string) int {
	n, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return n
}

func build(content []string, headerSize int) (fields map[string][]int, valids []int) {
	fields = make(map[string][]int)
	valids = make([]int, 0)

	for _, line := range content[0:headerSize] {
		values := strings.Split(line, ": ")
		key := values[0]
		fields[key] = make([]int, 0)

		numbers := strings.Split(values[1], " or ")
		for _, val := range numbers {
			values := strings.Split(val, "-")
			a := toInt(values[0])
			b := toInt(values[1])
			for ; a <= b; a++ {
				fields[key] = append(fields[key], a)
				valids = append(valids, a)
			}
		}
	}

	return
}

func nearby(content []string, valids []int, nearbyTicketsStart, nearbyTicketsEnd int) (validNearby [][]int, errorRate int) {
	validNearby = make([][]int, nearbyTicketsEnd-nearbyTicketsStart)

	for i, line := range content[nearbyTicketsStart:nearbyTicketsEnd] {
		values := strings.Split(line, ",")
		validNearby[i] = make([]int, 0)
		for _, val := range values {
			n := toInt(val)
			if !contains(valids, n) {
				errorRate += n
				continue
			}
			validNearby[i] = append(validNearby[i], n)
		}
	}

	return validNearby, errorRate
}

func what(content []string, headerSize, nearbyTicketsStart, nearbyTicketsEnd int) map[string]int {
	fields, valids := build(content, headerSize)
	validNearby, _ := nearby(content, valids, nearbyTicketsStart, nearbyTicketsEnd)

	maxCol := len(validNearby)
	maxRow := len(validNearby[0])
	// mapping := make(map[int]string)
	mapping := make(map[string]int)

	for col := 0; col < maxCol; col++ {
		set := NewSet(fields)

		for i := 0; i < maxRow; i++ {
			itemset := NewSet(nil)
			for k := range fields {
				if _, ok := mapping[k]; ok {
					continue
				}
				if contains(fields[k], validNearby[i][col]) {
					itemset.Add(k)
				}
			}
			set = set.Intersection(itemset)
		}

		for k := range mapping {
			set.Del(k)
		}

		if set.Size() > 1 {
			panic(fmt.Sprintf("too many matches: %v", set))
			// fmt.Println(fmt.Sprintf("too many matches: %v", set.Get()))
		}
		mapping[set.First()] = col
		// mapping[col] = set.First()
		// fmt.Println("intersection", set)
	}
	// for k, v := range mapping {
	// 	fmt.Println(k, v)
	// }

	return mapping
}
