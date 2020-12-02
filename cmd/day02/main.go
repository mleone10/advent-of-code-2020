package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	min, max int
	c, pwd   string
}

func main() {
	ls := []line{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		l := strings.ReplaceAll(scanner.Text(), ":", " ")
		l = strings.ReplaceAll(l, "-", " ")
		parts := strings.Fields(l)
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		ls = append(ls, line{
			min: min,
			max: max,
			c:   parts[2],
			pwd: parts[3],
		})
	}

	log.Printf("Number of valid passwords by count: %d", countValidPasswords(ls, isValidCount))
	log.Printf("Number of valid passwords by index: %d", countValidPasswords(ls, isValidIndex))
}

func countValidPasswords(ls []line, isValid func(l line) bool) int {
	var sum int

	for _, l := range ls {
		if isValid(l) {
			sum++
		}
	}

	return sum
}

func isValidCount(l line) bool {
	c := strings.Count(l.pwd, l.c)
	return c >= l.min && c <= l.max
}

func isValidIndex(l line) bool {
	return (string(l.pwd[l.min-1]) == l.c) != (string(l.pwd[l.max-1]) == l.c)
}
