package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Pers map[rune]int
type Group struct {
	persons []Pers
	part1   Pers
	part2   Pers
}

func readFile(fname string) (elems []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	elems = make([]string, 0, len(lines))

	for _, l := range lines {
		elems = append(elems, l)
	}

	return elems, nil
}

func main() {
	var answ []Group

	elems, err := readFile("input")
	if err != nil {
		panic(err)
	}

	crtGroup := Group{}
	crtGroup.part1 = Pers{}

	for _, elem := range elems {
		if elem == "" {
			answ = append(answ, crtGroup)
			crtGroup = Group{}
			crtGroup.part1 = Pers{}

			continue
		}

		crtPers := Pers{}
		for _, ch := range elem {
			crtPers[ch] = 1
			crtGroup.part1[ch]++
		}

		crtGroup.persons = append(crtGroup.persons, crtPers)
	}

	part1 := 0
	part2 := 0

	for _, a := range answ {
		part1 += len(a.part1)
	}

	for k, a := range answ {
		answ[k].part2 = Pers{}
		for c, _ := range a.part1 {
			found := 1

			for _, p := range a.persons {
				if _, f := p[c]; f == false {
					found = 0
				}
			}

			if found != 0 {
				answ[k].part2[c]++
			}
		}
	}

	for _, a := range answ {
		part2 += len(a.part2)
	}

	fmt.Printf("part 1: %#v\n", part1)
	fmt.Printf("part 2: %#v\n", part2)
}
