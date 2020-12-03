package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		nums = append(nums, n)
	}

	return nums, nil
}

func main() {
	nums, err := readFile("input")
	if err != nil {
		panic(err)
	}

	for i, num1 := range nums {
		for _, num2 := range nums[i:] {
			if num1+num2 == 2020 {
				fmt.Printf("part 1: %#v\n", num1*num2)
			}
		}
	}

	for i, num1 := range nums {
		for j, num2 := range nums[i:] {
			for _, num3 := range nums[i+j:] {
				if num1+num2+num3 == 2020 {
					fmt.Printf("part 2: %#v\n", num1*num2*num3)
				}
			}
		}
	}
}
