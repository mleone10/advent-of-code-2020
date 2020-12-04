package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type passport map[string]string

func main() {
	ps := []passport{}
	scanner := bufio.NewScanner(os.Stdin)

	p := passport{}
	for scanner.Scan() {
		fs := strings.Fields(scanner.Text())
		if len(fs) == 0 {
			ps = append(ps, p)
			p = passport{}
			continue
		}

		for _, f := range fs {
			p[strings.Split(f, ":")[0]] = strings.Split(f, ":")[1]
		}
	}

	log.Printf("Valid passports: %d", countValidPassports(ps))
}

func countValidPassports(ps []passport) int {
	var sum int

	for _, p := range ps {
		if isValid(p) {
			sum++
		} else {
		}
	}

	return sum
}

func isValid(p passport) bool {
	ks := keys(p)
	_, cidPresent := p["cid"]
	return (len(ks) == 8 || (len(ks) == 7 && !cidPresent)) && allFieldsValid(p)
}

func allFieldsValid(p passport) bool {
	if byr, _ := strconv.Atoi(p["byr"]); byr < 1920 || byr > 2002 {
		return false
	}
	if iyr, _ := strconv.Atoi(p["iyr"]); iyr < 2010 || iyr > 2020 {
		return false
	}
	if eyr, _ := strconv.Atoi(p["eyr"]); eyr < 2020 || eyr > 2030 {
		return false
	}
	if p["hgt"][len(p["hgt"])-2:] != "cm" && p["hgt"][len(p["hgt"])-2:] != "in" {
		return false
	}
	if p["hgt"][len(p["hgt"])-2:] == "cm" {
		if hgt, _ := strconv.Atoi(p["hgt"][:len(p["hgt"])-2]); hgt < 150 || hgt > 193 {
			return false
		}
	}
	if p["hgt"][len(p["hgt"])-2:] == "in" {
		if hgt, _ := strconv.Atoi(p["hgt"][:len(p["hgt"])-2]); hgt < 59 || hgt > 76 {
			return false
		}
	}
	if string(p["hcl"][0]) != "#" {
		return false
	} else {
		_, err := strconv.ParseUint(p["hcl"][1:], 16, 64)
		if err != nil {
			return false
		}
	}
	if p["ecl"] != "amb" && p["ecl"] != "blu" && p["ecl"] != "gry" && p["ecl"] != "grn" && p["ecl"] != "hzl" && p["ecl"] != "oth" && p["ecl"] != "brn" {
		return false
	}
	if len(p["pid"]) != 9 {
		return false
	}
	if _, err := strconv.Atoi(p["pid"]); err != nil {
		return false
	}
	return true
}

func keys(p passport) []string {
	var ks []string
	for k := range p {
		ks = append(ks, k)
	}
	return ks
}
