package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regexMem = regexp.MustCompile(`mem\[(\d*)\] = (\d*)`)

type memory map[adr]val
type adr int
type val int
type mask string
type masker struct {
	mem memory
	m   mask
}

func main() {
	var masker masker
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		switch l := scanner.Text(); strings.Contains(l, "mask") {
		case true:
			masker.setMask(mask(strings.Split(l, " = ")[1]))
		case false:
			memLine := regexMem.FindAllStringSubmatch(l, -1)
			a, _ := strconv.Atoi(memLine[0][1])
			v, _ := strconv.Atoi(memLine[0][2])
			masker.store(adr(a), val(v))
		}
	}

	log.Printf("Sum of memory after all operations: %d", masker.sum())
}

func (masker *masker) setMask(m mask) {
	masker.m = m
}

func (masker *masker) store(a adr, v val) {
	if masker.mem == nil {
		masker.mem = make(memory)
	}

	masker.mem[a] = masker.mask(v)
}

func (masker masker) mask(v val) val {
	b := fmt.Sprintf("%036b", int(v))
	var newB string
	for i, m := range masker.m {
		switch m {
		case '0':
			newB += "0"
		case '1':
			newB += "1"
		case 'X':
			newB += string(b[i])
		}
	}
	newV, _ := strconv.ParseInt(newB, 2, 64)
	return val(newV)
}

func (masker masker) sum() int {
	var sum int
	for _, v := range masker.mem {
		sum += int(v)
	}
	return sum
}
