package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Byr, Iyr, Eyr, Hgt, Hcl, Ecl, Pid, Cid string
}

var passports []Passport

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

// setField sets field of v with given name to given value.
func setField(v interface{}, name string, value string) error {
	// v must be a pointer to a struct
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("v must be pointer to struct")
	}

	// Dereference pointer
	rv = rv.Elem()

	// Lookup field by name
	fv := rv.FieldByName(name)
	if !fv.IsValid() {
		return fmt.Errorf("not a field name: %s", name)
	}

	// Field must be exported
	if !fv.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}

	// We expect a string field
	if fv.Kind() != reflect.String {
		return fmt.Errorf("%s is not a string field", name)
	}

	// Set the value
	fv.SetString(value)

	return nil
}

func isValid1(p Passport) bool {
	if p.Byr == "" || p.Iyr == "" || p.Eyr == "" || p.Hgt == "" || p.Hcl == "" || p.Ecl == "" || p.Pid == "" {
		return false
	}

	return true
}

func isValid2(p Passport) bool {
	byr, err := strconv.Atoi(p.Byr)
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}

	iyr, err := strconv.Atoi(p.Iyr)
	if err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, err := strconv.Atoi(p.Eyr)
	if err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	hgtRe1 := regexp.MustCompile(`(?P<h>[0-9]{3})cm`)
	hgtRe2 := regexp.MustCompile(`(?P<h>[0-9]{2})in`)

	m1 := hgtRe1.FindStringSubmatch(p.Hgt)
	if len(m1) == 0 {
		m2 := hgtRe2.FindStringSubmatch(p.Hgt)
		if len(m2) == 0 {
			return false
		} else {
			h, _ := strconv.Atoi(m2[1])
			if h < 59 || h > 76 {
				return false
			}
		}
	} else {
		h, _ := strconv.Atoi(m1[1])
		if h < 150 || h > 193 {
			return false
		}
	}

	if m, _ := regexp.Match(`#[0-9a-f]{6}`, []byte(p.Hcl)); !m {
		return false
	}

	if m, _ := regexp.Match(`^(amb|blu|brn|gry|grn|hzl|oth)$`, []byte(p.Ecl)); !m {
		return false
	}

	if m, _ := regexp.Match(`^[0-9]{9}$`, []byte(p.Pid)); !m {
		return false
	}

	return true
}

func main() {
	var passport Passport

	elems, err := readFile("input")
	if err != nil {
		panic(err)
	}

	valid1 := 0
	valid2 := 0

	for _, elem := range elems {
		if elem == "" {
			if isValid1(passport) {
				valid1++

				if isValid2(passport) {
					valid2++
				}
			}

			passports = append(passports, passport)
			passport = Passport{}

			continue
		}

		lSplit := strings.Split(elem, " ")

		for _, iSplit := range lSplit {
			eSplit := strings.Split(iSplit, ":")

			err := setField(&passport, strings.Title(eSplit[0]), eSplit[1])
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Printf("part 1: %#v\n", valid1)
	fmt.Printf("part 2: %#v\n", valid2)
}
