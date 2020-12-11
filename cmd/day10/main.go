package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	ints := []int{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, i)
	}

	sort.Ints(ints)
	ints = append([]int{0}, ints...)
	ints = append(ints, ints[len(ints)-1]+3)

	log.Printf("Sum of 1-jolt diffs multiplied by sum of 3-jold diffs: %d", calcProductJoltDiffSums(ints))
	log.Printf("Possible combinations of adapters: %d", countRoutes(0, ints, map[int]int{}))
}

func calcProductJoltDiffSums(ints []int) int {
	var ones, threes int

	for i := 0; i < len(ints)-1; i++ {
		if ints[i+1]-ints[i] == 1 {
			ones++
		} else if ints[i+1]-ints[i] == 3 {
			threes++
		}
	}

	return ones * threes
}

func countRoutes(index int, ints []int, visited map[int]int) int {
	if index >= len(ints)-3 {
		return 1
	}

	origin := ints[index]
	if v, ok := visited[origin]; ok {
		return v
	}

	var num int
	for i := 1; i <= 3; i++ {
		candidate := ints[index+i]
		if match(origin, candidate) {
			num += countRoutes(index+i, ints, visited)
		}
	}

	visited[origin] = num
	return num
}

func match(low, high int) bool {
	return low+1 == high || low+2 == high || low+3 == high
}
