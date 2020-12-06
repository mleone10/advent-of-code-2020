package main

import (
	"bufio"
	"log"
	"os"
)

type person string
type group []person

func main() {
	gs := []group{}
	scanner := bufio.NewScanner(os.Stdin)

	g := group{}
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			gs = append(gs, g)
			g = group{}
			continue
		}
		g = append(g, person(l))
	}
	gs = append(gs, g)

	log.Printf("Sum of questions in each group: %d", sumGroupCounts(gs, countQuestions))
	log.Printf("Sum of all-yes questions in each group: %d", sumGroupCounts(gs, countAllYes))
}

func sumGroupCounts(gs []group, countFunc func(g group) int) int {
	var sum int

	for _, g := range gs {
		sum += countFunc(g)
	}

	return sum
}

func countQuestions(g group) int {
	rs := aggGroupResponses(g)
	return len(rs)
}

func countAllYes(g group) int {
	rs := aggGroupResponses(g)

	sum := 0
	for _, c := range rs {
		if c == len(g) {
			sum++
		}
	}

	return sum
}

func aggGroupResponses(g group) map[string]int {
	rs := map[string]int{}

	for _, p := range g {
		for _, r := range p {
			rs[string(r)]++
		}
	}

	return rs
}
