package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const toFind = "shiny gold"

type Contents map[string]int
type Bags map[string]Contents

var bags Bags

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

func getContents(b string, s string) {
	subRE := regexp.MustCompile(`(?P<count>[1-9]) (?P<name>[a-z]+ [a-z]+) bags?, (?P<rest>.+\.)`)

	match := subRE.FindStringSubmatch(s)
	if match == nil {
		lastRE := regexp.MustCompile(`(?P<count>[1-9]) (?P<name>[a-z]+ [a-z]+) bags?\.`)

		matchLast := lastRE.FindStringSubmatch(s)
		if matchLast != nil {
			if bags[b] == nil {
				bags[b] = make(Contents)
			}

			count, _ := strconv.Atoi(matchLast[1])
			bags[b][matchLast[2]] = count
		} else {
			bags[b] = nil
		}
	} else {
		if bags[b] == nil {
			bags[b] = make(Contents)
		}

		count, _ := strconv.Atoi(match[1])
		bags[b][match[2]] = count

		getContents(b, match[3])
	}
}

func findBag(b string) bool {
	for k, _ := range bags[b] {
		if k == toFind || findBag(k) {
			return true
		}
	}

	return false
}

func countBags(b string) int {
	count := 1

	if bags[b] != nil {
		for k, v := range bags[b] {
			count += v * countBags(k)
		}
	}

	return count
}

func main() {
	bags = make(Bags)

	elems, err := readFile("input")
	if err != nil {
		panic(err)
	}

	bagnameRE := regexp.MustCompile(`(?P<name>[a-z]+ [a-z]+) bags contain (?P<contents>.+\.)`)

	for _, elem := range elems {
		match := bagnameRE.FindStringSubmatch(elem)
		getContents(match[1], match[2])
	}

	part1 := 0
	for k, _ := range bags {
		if findBag(k) {
			part1++
		}
	}

	part2 := countBags(toFind) - 1

	fmt.Printf("part 1: %#v\n", part1)
	fmt.Printf("part 2: %#v\n", part2)
}
