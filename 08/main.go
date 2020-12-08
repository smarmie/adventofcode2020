package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type I struct {
	inst    string
	val     int
	visited bool
}

var acumulator int
var pIndex int
var p []I

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

func process(mutate int) (int, bool) {
	crtacc := acumulator
	crtpos := pIndex
	crtinst := p[pIndex].inst

	if pIndex == mutate {
		switch crtinst {
		case "jmp":
			crtinst = "nop"
		case "nop":
			crtinst = "jmp"
		}
	}

	switch crtinst {
	case "acc":
		acumulator += p[pIndex].val
		pIndex++
	case "jmp":
		pIndex += p[pIndex].val
	case "nop":
		pIndex++
	}

	p[crtpos].visited = true

	if pIndex >= len(p) {
		return acumulator, true
	}

	if p[pIndex].visited {
		return crtacc, true
	}

	return 0, false
}

func resetVisited() {
	for k := range p {
		p[k].visited = false
	}
}

func main() {
	elems, err := readFile("input")
	if err != nil {
		panic(err)
	}

	for _, elem := range elems {
		parts := strings.Split(elem, " ")
		off, _ := strconv.Atoi(parts[1])
		p = append(p, I{parts[0], off, false})
	}

	part1 := 0

	for finished := false; !finished; {
		part1, finished = process(-1)
	}

	part2 := 0

	for mutate := 0; mutate < len(p) && pIndex < len(p); mutate++ {
		acumulator = 0
		pIndex = 0

		resetVisited()

		for finished := false; !finished; {
			part2, finished = process(mutate)
		}
	}

	fmt.Printf("part 1: %#v\n", part1)
	fmt.Printf("part 2: %#v\n", part2)
}
