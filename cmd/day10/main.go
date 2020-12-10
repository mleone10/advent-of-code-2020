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
	log.Printf("Possible combinations of adapters: %d", computeCombinations(ints))
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

func computeCombinations(ints []int) int {
	return 0
}
