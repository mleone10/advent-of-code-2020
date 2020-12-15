package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regexMem = regexp.MustCompile(`mem\[(\d*)\] = (\d*)`)

type adr int
type val int
type memory map[adr]val
type mask string
type valMasker struct {
	memory
	m mask
}
type adrMasker struct {
	memory
	m     mask
	masks []mask
}

func main() {
	var vm valMasker
	var am adrMasker
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		switch l := scanner.Text(); strings.Contains(l, "mask") {
		case true:
			m := strings.Split(l, " = ")[1]
			vm.setMask(m)
			am.setMask(m)
		case false:
			memLine := regexMem.FindAllStringSubmatch(l, -1)
			a, _ := strconv.Atoi(memLine[0][1])
			v, _ := strconv.Atoi(memLine[0][2])
			vm.store(adr(a), val(v))
			am.store(adr(a), val(v))
		}
	}

	log.Printf("Sum of memory after value masking: %d", vm.sum())
	log.Printf("Sum of memory after address masking: %d", am.sum())
}

func (m *valMasker) store(a adr, v val) {
	if m.memory == nil {
		m.memory = make(memory)
	}

	m.memory.store(a, m.maskVal(v))
}

func (m *valMasker) setMask(s string) {
	m.m = mask(s)
}

func (m valMasker) maskVal(v val) val {
	bVal := fmt.Sprintf("%036b", int(v))
	var newB string
	for i, b := range m.m {
		switch b {
		case '0':
			newB += "0"
		case '1':
			newB += "1"
		case 'X':
			newB += string(bVal[i])
		}
	}
	newV, _ := strconv.ParseInt(newB, 2, 64)
	return val(newV)
}

func (m *adrMasker) store(a adr, v val) {
	if m.memory == nil {
		m.memory = make(memory)
	}

	for _, ma := range m.maskAdr(a) {
		m.memory.store(ma, v)
	}
}

func (m *adrMasker) maskAdr(a adr) []adr {
	return m.genAdrPermutations(m.applyMask(a))
}

func (m *adrMasker) applyMask(a adr) mask {
	bAdr := fmt.Sprintf("%036b", int(a))
	var nbAdr string
	for i, b := range m.m {
		switch b {
		case '0':
			nbAdr += string(bAdr[i])
		default:
			nbAdr += string(b)
		}
	}
	return mask(nbAdr)
}

func (m *adrMasker) genAdrPermutations(a mask) []adr {
	as := []adr{}
	xs := strings.Count(string(a), "X")
	mx := int(math.Pow(2.0, float64(xs)))
	for i := 0; i < mx; i++ {
		bX := fmt.Sprintf("%0*b", len(fmt.Sprintf("%b", mx-1)), i)
		sAdr := string(a)
		for _, b := range bX {
			sAdr = strings.Replace(sAdr, "X", string(b), 1)
		}
		dAdr, _ := strconv.ParseInt(sAdr, 2, 64)
		as = append(as, adr(dAdr))
	}
	return as
}

func (m *adrMasker) setMask(s string) {
	m.m = mask(s)
}

func (m memory) store(a adr, v val) {
	m[a] = v
}

func (m memory) sum() int {
	var sum int
	for _, v := range m {
		sum += int(v)
	}
	return sum
}
