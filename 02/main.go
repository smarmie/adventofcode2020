package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
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

func main() {
	elems, err := readFile("input")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("^(?P<min>[0-9]+)-(?P<max>[0-9]+) (?P<char>[a-z]): (?P<password>[a-z]+)$")

	valid1 := 0
	valid2 := 0

	for _, elem := range elems {
		match := re.FindStringSubmatch(elem)
		min, _ := strconv.Atoi(match[1])
		max, _ := strconv.Atoi(match[2])
		char := match[3]
		pass := match[4]
		lCount := strings.Count(pass, char)

		if lCount >= min && lCount <= max {
			valid1++
		}

		if (pass[min-1] == char[0]) != (pass[max-1] == char[0]) {
			valid2++
		}
	}

	fmt.Printf("part 1: %#v\n", valid1)
	fmt.Printf("part 2: %#v\n", valid2)
}
