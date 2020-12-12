package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Ship1 struct {
	x, y, heading int
}

type Ship2 struct {
	x, y, wpx, wpy int
}

type Move struct {
	dir  string
	dist int
}

var ship1 = Ship1{
	x:       0,
	y:       0,
	heading: 0,
}

var ship2 = Ship2{
	x:   0,
	y:   0,
	wpx: 10,
	wpy: 1,
}

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

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func moveShip1(move Move) {
	switch move.dir {
	case "N":
		ship1.y += move.dist
	case "S":
		ship1.y -= move.dist
	case "E":
		ship1.x += move.dist
	case "W":
		ship1.x -= move.dist
	case "R":
		ship1.heading = (ship1.heading - move.dist) % 360
		if ship1.heading < 0 {
			ship1.heading += 360
		}
	case "L":
		ship1.heading = (ship1.heading + move.dist) % 360
	case "F":
		switch ship1.heading {
		case 0:
			moveShip1(Move{
				dir:  "E",
				dist: move.dist,
			})
		case 90:
			moveShip1(Move{
				dir:  "N",
				dist: move.dist,
			})
		case 180:
			moveShip1(Move{
				dir:  "W",
				dist: move.dist,
			})
		case 270:
			moveShip1(Move{
				dir:  "S",
				dist: move.dist,
			})
		}
	}
}

func moveShip2(move Move) {
	switch move.dir {
	case "N":
		ship2.wpy += move.dist
	case "S":
		ship2.wpy -= move.dist
	case "E":
		ship2.wpx += move.dist
	case "W":
		ship2.wpx -= move.dist
	case "F":
		ship2.x += move.dist * ship2.wpx
		ship2.y += move.dist * ship2.wpy
	case "R":
		switch move.dist {
		case 90:
			ship2.wpx, ship2.wpy = ship2.wpy, 0-ship2.wpx
		case 180:
			ship2.wpx, ship2.wpy = 0-ship2.wpx, 0-ship2.wpy
		case 270:
			ship2.wpx, ship2.wpy = 0-ship2.wpy, ship2.wpx
		}
	case "L":
		switch move.dist {
		case 90:
			ship2.wpx, ship2.wpy = 0-ship2.wpy, ship2.wpx
		case 180:
			ship2.wpx, ship2.wpy = 0-ship2.wpx, 0-ship2.wpy
		case 270:
			ship2.wpx, ship2.wpy = ship2.wpy, 0-ship2.wpx
		}
	}
}

func main() {
	elems, err := readFile("input")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("(?P<dir>[A-Z])(?P<dist>[0-9]+)")

	for _, elem := range elems {
		match := re.FindStringSubmatch(elem)
		dist, _ := strconv.Atoi(match[2])
		move := Move{
			dir:  match[1],
			dist: dist,
		}
		// moves = append(moves, move)
		moveShip1(move)
		moveShip2(move)
	}

	fmt.Printf("part 1: %#v\n", Abs(ship1.x)+Abs(ship1.y))
	fmt.Printf("part 1: %#v\n", Abs(ship2.x)+Abs(ship2.y))
}
