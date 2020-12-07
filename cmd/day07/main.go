package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const colorShinyGold = "shiny gold"

var contentsMatch = regexp.MustCompile(`(?U:(\d+) (.+) bag[s]?)`)

type color string
type bags map[color]bag
type contents map[color]int

type bag struct {
	color    color
	contents map[color]int
}

func main() {
	bs := bags{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		b := bag{color(""), contents{}}
		l := strings.Split(scanner.Text(), " bags contain ")
		b.color = color(l[0])
		for _, c := range contentsMatch.FindAllStringSubmatch(l[1], -1) {
			count, _ := strconv.Atoi(c[1])
			b.contents[color(c[2])] = count
		}
		bs[b.color] = b
	}

	cs := computeContentsOfOuterBags(bs)
	log.Printf("Number of valid outer bags: %d", sumValidOuterBags(colorShinyGold, cs))
	log.Printf("Number of bags inside shiny gold bag: %d", getNumBagsInColor(colorShinyGold, cs))
}

func computeContentsOfOuterBags(bs bags) map[color]contents {
	cs := map[color]contents{}

	for _, b := range bs {
		c := contents{}
		for color, n := range traverseBag(b.color, bs) {
			c[color] += n
		}
		cs[b.color] = c
	}

	return cs
}

func sumValidOuterBags(color color, cs map[color]contents) int {
	var sum int

	for _, cs := range cs {
		if outerBagContainsColor(color, cs) {
			sum++
		}
	}

	return sum
}

func outerBagContainsColor(color color, cs contents) bool {
	for c := range cs {
		if c == color {
			return true
		}
	}
	return false
}

func getNumBagsInColor(color color, contents map[color]contents) int {
	var sum int

	for _, num := range contents[color] {
		sum += num
	}

	return sum
}

func traverseBag(color color, bs bags) contents {
	cs := contents{}

	for color, count := range bs[color].contents {
		cs[color] += count
		for color, num := range traverseBag(color, bs) {
			cs[color] += num * count
		}
	}

	return cs
}
