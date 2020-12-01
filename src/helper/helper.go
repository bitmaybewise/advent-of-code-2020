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
	return strings.Split(string(input), "\n")
}

func LinesToInt(s []string) []int {
	ints := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			continue
		}
		n, err := strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}
	return ints
}
