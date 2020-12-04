package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var ecls = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
	numFs                                  int
}

func main() {
	ps := []passport{}
	scanner := bufio.NewScanner(os.Stdin)

	acc := map[string]string{}
	for scanner.Scan() {
		fs := strings.Fields(scanner.Text())
		if len(fs) == 0 {
			// If no fields were extracted, this is an empty line
			// Create a new passport with the accumulated fields, then reset the map
			ps = append(ps, newPassport(acc))
			acc = map[string]string{}
			continue
		}

		for _, f := range fs {
			kv := strings.Split(f, ":")
			acc[kv[0]] = kv[1]
		}
	}
	// Last passport - we accumulated it, but never created it
	ps = append(ps, newPassport(acc))

	log.Printf("Valid passports by field count: %d", countValidPassports(ps, validateFields))
	log.Printf("Valid passports by data: %d", countValidPassports(ps, validateData))
}

func newPassport(fs map[string]string) passport {
	return passport{
		byr:   fs["byr"],
		iyr:   fs["iyr"],
		eyr:   fs["eyr"],
		hgt:   fs["hgt"],
		hcl:   fs["hcl"],
		ecl:   fs["ecl"],
		pid:   fs["pid"],
		cid:   fs["cid"],
		numFs: len(fs),
	}
}

func countValidPassports(ps []passport, isValid func(p passport) bool) int {
	var sum int

	for _, p := range ps {
		if isValid(p) {
			sum++
		} else {
		}
	}

	return sum
}

func validateFields(p passport) bool {
	return p.numFs == 8 || p.numFs == 7 && isEmpty(p.cid)
}

func validateData(p passport) bool {
	return p.byrIsValid() &&
		p.iyrIsValid() &&
		p.eyrIsValid() &&
		p.hgtIsValid() &&
		p.hclIsValid() &&
		p.eclIsValid() &&
		p.pidIsValid()
}

func (p passport) byrIsValid() bool {
	byr, _ := strconv.Atoi(p.byr)
	return byr >= 1920 && byr <= 2002
}

func (p passport) iyrIsValid() bool {
	iyr, _ := strconv.Atoi(p.iyr)
	return iyr >= 2010 && iyr <= 2020
}

func (p passport) eyrIsValid() bool {
	eyr, _ := strconv.Atoi(p.eyr)
	return eyr >= 2020 && eyr <= 2030
}

func (p passport) hgtIsValid() bool {
	if isEmpty(p.hgt) {
		return false
	}

	lenHgt := len(p.hgt)
	switch p.hgt[lenHgt-2:] {
	case "cm":
		if hgt, _ := strconv.Atoi(p.hgt[:lenHgt-2]); hgt >= 150 && hgt <= 193 {
			return true
		}
	case "in":
		if hgt, _ := strconv.Atoi(p.hgt[:lenHgt-2]); hgt >= 59 && hgt <= 76 {
			return true
		}
	default:
		return false
	}
	return false
}

func (p passport) hclIsValid() bool {
	if isEmpty(p.hcl) {
		return false
	}

	if string(p.hcl[0]) != "#" {
		return false
	}
	_, err := strconv.ParseUint(p.hcl[1:], 16, 64)
	if err != nil {
		return false
	}

	return true
}

func (p passport) eclIsValid() bool {
	return ecls[p.ecl]
}

func (p passport) pidIsValid() bool {
	if len(p.pid) != 9 {
		return false
	}
	if _, err := strconv.Atoi(p.pid); err != nil {
		return false
	}

	return true
}

func isEmpty(f string) bool {
	return f == ""
}
