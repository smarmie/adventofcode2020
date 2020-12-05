package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func readFile(fname string) (elems []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	elems = make([]string, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		elems = append(elems, l)
	}

	return elems, nil
}

func binsearch(dir string, down, up rune, start, end int) int {
	if start == end {
		return start
	}

	if rune(dir[0]) == up {
		return binsearch(dir[1:], down, up, start+int(math.Ceil(float64(end-start)/float64(2))), end)
	}

	return binsearch(dir[1:], down, up, start, start+int(math.Floor(float64(end-start)/float64(2))))
}

func main() {
	var idList []int

	var ticket int

	elems, err := readFile("input")
	if err != nil {
		panic(err)
	}

	high := 0

	for _, elem := range elems {
		row := binsearch(elem[0:7], 'F', 'B', 0, 127)
		column := binsearch(elem[7:10], 'L', 'R', 0, 7)

		seatID := row*8 + column
		if seatID > high {
			high = seatID
		}

		idList = append(idList, seatID)
	}

	sort.Ints(idList)

	for pos, id := range idList {
		if pos == 0 {
			continue
		}

		if id-idList[pos-1] == 2 {
			ticket = id - 1
		}
	}

	fmt.Printf("part 1: %#v\n", high)
	fmt.Printf("part 2: %#v\n", ticket)
}
