package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"

	aoc "github.com/mleone10/advent-of-code-2020"
)

func main() {
	ps := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ps = append(ps, strings.TrimSpace(scanner.Text()))
	}

	log.Printf("Highest seat ID on any boarding pass: %d", calcMaxSeatID(ps))
	log.Printf("My seat ID: %d", calcMySeatID(ps))
}

func calcMaxSeatID(ps []string) int {
	return aoc.IntSliceMax(calcSeatIDs(ps))
}

func calcMySeatID(ps []string) int {
	ids := calcSeatIDs(ps)
	sort.Ints(ids)

	prevID := ids[0]
	for _, id := range ids[1:] {
		expected := prevID + 1
		if expected != id {
			return expected
		}
		prevID = id
	}

	return 0
}

func calcSeatIDs(ps []string) []int {
	ids := []int{}

	for _, p := range ps {
		ids = append(ids, calcSeatID(p))
	}

	return ids
}

func calcSeatID(p string) int {
	row := calcRow(p[:7])
	col := calcCol(p[7:])

	return row*8 + col
}

func calcRow(p string) int {
	rs := aoc.InitIntSlice(0, 127)

	for _, c := range p {
		switch string(c) {
		case "F":
			rs = rs[:len(rs)/2]
		case "B":
			rs = rs[len(rs)/2:]
		}
	}

	return rs[0]
}

func calcCol(p string) int {
	rs := aoc.InitIntSlice(0, 7)

	for _, c := range p {
		switch string(c) {
		case "L":
			rs = rs[:len(rs)/2]
		case "R":
			rs = rs[len(rs)/2:]
		}
	}

	return rs[0]
}
