package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	ns := []int{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		ns = append(ns, n)
	}

	log.Printf("Product of two elements that sum to 2020: %d", findDoubleProduct(ns))
	log.Printf("Product of three elements that sum to 2020: %d", findTripleProduct(ns))

}

func findDoubleProduct(ns []int) int {
	for i, n := range ns {
		for _, m := range ns[i+1:] {
			if n+m == 2020 {
				return n * m
			}
		}
	}
	return -1
}

func findTripleProduct(ns []int) int {
	for i, n := range ns {
		for _, m := range ns[i+1:] {
			for _, o := range ns[i+2:] {
				if n+m+o == 2020 {
					return n * m * o
				}
			}
		}
	}
	return -1
}
