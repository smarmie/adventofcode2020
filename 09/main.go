package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const P = 25

var c []int

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

	for _, elem := range elems {
		ic, _ := strconv.Atoi(elem)
		c = append(c, ic)
	}

	part1 := 0

	for k1, v1 := range c[P:] {
		found := false

		for k2, v2 := range c[k1 : k1+P] {
			for _, v3 := range c[k1+k2+1 : k1+P] {
				if v2+v3 == v1 {
					found = true
				}
			}
		}

		if !found {
			part1 = v1
		}

		if part1 != 0 {
			break
		}
	}

	found := false
	is := 0
	ie := 0

	for k1 := range c {
		if found {
			break
		}

		for k2 := k1 + 1; k2 < len(c); k2++ {
			s := 0

			for k3 := k1; k3 <= k2; k3++ {
				s += c[k3]
			}

			if s > part1 {
				break
			}

			if s == part1 {
				found = true
				is = k1
				ie = k2
			}
		}
	}

	min := c[is]
	max := c[is]

	for i := is; i <= ie; i++ {
		if c[i] < min {
			min = c[i]
		}

		if c[i] > max {
			max = c[i]
		}
	}

	part2 := min + max

	print(found)
	fmt.Printf("part 1: %#v\n", part1)
	fmt.Printf("part 2: %#v\n", part2)
}
