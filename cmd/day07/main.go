package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var contentsMatch = regexp.MustCompile(`(?U:(\d+) (.+) bag[s]?)`)

type rule struct {
	color    string
	contents []content
}

type content struct {
	count int
	color string
}

func main() {
	rs := map[string]rule{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ruleString := strings.Split(scanner.Text(), " bags contain ")
		color := ruleString[0]
		contentStrings := contentsMatch.FindAllStringSubmatch(ruleString[1], -1)
		cs := []content{}
		for _, c := range contentStrings {
			count, _ := strconv.Atoi(c[1])
			cs = append(cs, content{
				count: count,
				color: c[2],
			})
		}
		rs[color] = rule{
			color:    color,
			contents: cs,
		}
	}

	log.Printf("Number of valid outer bags: %d", sumValidOuterBags(rs))
	log.Printf("Number of bags inside shiny gold bag: %d", computeContentsOfOuterBagsP2(rs)["shiny gold"])
}

func sumValidOuterBags(rs map[string]rule) int {
	var sum int

	for _, contents := range computeContentsOfOuterBags(rs) {
		if outerBagContainsShinyGold(contents) {
			sum++
		}
	}

	return sum
}

func computeContentsOfOuterBags(rs map[string]rule) map[string][]string {
	bs := map[string][]string{}

	for _, r := range rs {
		bs[r.color] = traverseBag(r.color, rs)
	}

	return bs
}

func computeContentsOfOuterBagsP2(rs map[string]rule) map[string]int {
	bs := map[string]int{}

	for _, r := range rs {
		for _, n := range traverseBagNum(r.color, rs) {
			bs[r.color] += n
		}
	}

	return bs
}

func outerBagContainsShinyGold(cs []string) bool {
	for _, c := range cs {
		if c == "shiny gold" {
			return true
		}
	}
	return false
}

func traverseBag(color string, rs map[string]rule) []string {
	cs := []string{}

	for _, c := range rs[color].contents {
		cs = append(cs, c.color)
		cs = append(cs, traverseBag(c.color, rs)...)
	}

	return cs
}

func traverseBagNum(color string, rs map[string]rule) map[string]int {
	cs := map[string]int{}

	for _, c := range rs[color].contents {
		cs[c.color] += c.count
		for color, num := range traverseBagNum(c.color, rs) {
			cs[color] += num * c.count
		}
	}

	return cs
}
