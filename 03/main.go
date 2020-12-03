package main

import (
	"fmt"
	"io/ioutil"
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

func getSlopeTrees(elems []string, stepR int, stepD int) int {
	trees := 0
	pos := 0
	width := len(elems[0])

	for line, elem := range elems {
		if (line % stepD) != 0 {
			continue
		}

		if elem[pos:pos+1] == "#" {
			trees++
		}

		pos = (pos + stepR) % width
	}

	return trees
}

func main() {
	elems, err := readFile("input")
	if err != nil {
		panic(err)
	}

	total := getSlopeTrees(elems, 1, 1) *
		getSlopeTrees(elems, 3, 1) *
		getSlopeTrees(elems, 5, 1) *
		getSlopeTrees(elems, 7, 1) *
		getSlopeTrees(elems, 1, 2)

	fmt.Printf("part 1: %#v\n", getSlopeTrees(elems, 3, 1))
	fmt.Printf("part 2: %#v\n", total)
}
