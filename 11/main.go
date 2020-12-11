package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var P [][]rune

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

func getStep1(p [][]rune, k1, k2 int) rune {
	var adj []rune

	if p[k1][k2] == '.' {
		return '.'
	}

	for i := k1 - 1; i <= k1+1; i++ {
		for j := k2 - 1; j <= k2+1; j++ {
			if i < 0 || i >= len(p) {
				continue
			}

			if j < 0 || j >= len(p[k1]) {
				continue
			}

			if i == k1 && j == k2 {
				continue
			}

			adj = append(adj, p[i][j])
		}
	}

	adjOcc := 0

	for _, v := range adj {
		if v == '#' {
			adjOcc++
		}
	}

	switch p[k1][k2] {
	case 'L':
		if adjOcc == 0 {
			return '#'
		}
	case '#':
		if adjOcc >= 4 {
			return 'L'
		}
	}

	return p[k1][k2]
}

func getStep2(p [][]rune, k1, k2 int) rune {
	DIRS := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	if p[k1][k2] == '.' {
		return '.'
	}

	occ := 0

	for _, dir := range DIRS {
		for i, j := k1+dir[0], k2+dir[1]; i >= 0 && j >= 0 && i < len(p) && j < len(p[k1]); i, j = i+dir[0], j+dir[1] {
			if p[i][j] == 'L' {
				break
			}
			if p[i][j] == '#' {
				occ++
				break
			}
		}
	}

	switch p[k1][k2] {
	case 'L':
		if occ == 0 {
			return '#'
		}
	case '#':
		if occ >= 5 {
			return 'L'
		}
	}

	return p[k1][k2]
}

func main() {
	elems, err := readFile("input")
	if err != nil {
		panic(err)
	}

	for k1, elem := range elems {
		P = append(P, make([]rune, len(elem)))
		for k2, s := range elem {
			P[k1][k2] = s
		}
	}

	s := make([][]rune, len(P))
	d := make([][]rune, len(P))

	part1 := 0

	for i := range P {
		s[i] = make([]rune, len(P[i]))
		d[i] = make([]rune, len(P[i]))
		copy(s[i], P[i])
		copy(d[i], P[i])
	}

	for i := 0; ; i++ {
		// take a step, build dest array
		for k1 := 0; k1 < len(s); k1++ {
			for k2 := 0; k2 < len(s[k1]); k2++ {
				m := getStep1(s, k1, k2)
				d[k1][k2] = m
			}
		}

		// check if source == dest
		equal := true

		for k1 := range s {
			for k2 := range s[k1] {
				if s[k1][k2] != d[k1][k2] {
					equal = false
					break
				}
			}
		}

		if equal {
			break
		}

		for i := range P {
			copy(s[i], d[i])
		}
	}

	for k1 := range s {
		for k2 := range s[k1] {
			if s[k1][k2] == '#' {
				part1++
			}
		}
	}

	part2 := 0

	for i := range P {
		copy(s[i], P[i])
		copy(d[i], P[i])
	}

	for i := 0; ; i++ {
		// take a step, build dest array
		for k1 := 0; k1 < len(s); k1++ {
			for k2 := 0; k2 < len(s[k1]); k2++ {
				m := getStep2(s, k1, k2)
				d[k1][k2] = m
			}
		}

		// check if source == dest
		equal := true

		for k1 := range s {
			for k2 := range s[k1] {
				if s[k1][k2] != d[k1][k2] {
					equal = false
					break
				}
			}
		}

		if equal {
			break
		}

		for i := range P {
			copy(s[i], d[i])
		}
	}

	for k1 := range s {
		for k2 := range s[k1] {
			if s[k1][k2] == '#' {
				part2++
			}
		}
	}

	fmt.Printf("part 1: %#v\n", part1)
	fmt.Printf("part 2: %#v\n", part2)
}
