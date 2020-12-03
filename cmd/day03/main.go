package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const tree = "#"

type loc bool

type grid struct {
	locs [][]loc
}

type slope struct {
	x, y int
}

var slopes = [5]slope{
	slope{1, 1},
	slope{3, 1},
	slope{5, 1},
	slope{7, 1},
	slope{1, 2},
}

func main() {
	g := grid{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		r := []loc{}
		for _, l := range strings.Split(scanner.Text(), "") {
			r = append(r, (l == tree))
		}
		g.locs = append(g.locs, r)
	}

	log.Printf("Trees encountered with slope 3, 1: %d", numTreesOnSlope(g, slopes[1]))
	log.Printf("Product of trees encountered across all slopes: %d", prodOfAllSlopes(g))
}

func numTreesOnSlope(g grid, s slope) int {
	var x, y, sum int

	for y < len(g.locs) {
		if g.locs[y][x%len(g.locs[y])] {
			sum++
		}
		x += s.x
		y += s.y
	}

	return sum
}

func prodOfAllSlopes(g grid) int {
	prod := 1

	for _, s := range slopes {
		prod *= numTreesOnSlope(g, s)
	}

	return prod
}
