package helper

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadLinesFromInput(path string) []string {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	withLastLineEmpty := strings.Split(string(input), "\n")
	withoutLastLineEmpty := make([]string, len(withLastLineEmpty)-1)
	for i := range withoutLastLineEmpty {
		withoutLastLineEmpty[i] = withLastLineEmpty[i]
	}
	return withoutLastLineEmpty
}

func LinesToInt(s []string) []int {
	ints := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}
	return ints
}
