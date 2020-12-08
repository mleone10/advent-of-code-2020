package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	comAcc = command("acc")
	comJmp = command("jmp")
	comNop = command("nop")
)

type command string

type program struct {
	instructions, init []instruction
	pc, acc            int
}

type instruction struct {
	command command
	value   int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	is := []instruction{}
	for scanner.Scan() {
		l := strings.Fields(scanner.Text())
		v, _ := strconv.Atoi(l[1])
		i := instruction{
			command: command(l[0]),
			value:   v,
		}
		is = append(is, i)
	}

	p := program{
		instructions: is,
		init:         is,
		pc:           0,
		acc:          0,
	}

	p.doesLoop()
	log.Printf("Value in acc just before loop: %d", p.acc)

	p.reset()

	p.permutateToHalt()
	log.Printf("Value in acc after termination: %d", p.acc)

}

func (p *program) doesLoop() bool {
	pcs := map[int]bool{}

	for {
		p.step()
		if pcs[p.pc] {
			return true
		}
		if p.pc >= len(p.instructions) {
			return false
		}
		pcs[p.pc] = true
	}
}

func (p *program) permutateToHalt() {
	for i, inst := range p.instructions {
		if inst.command == comJmp {
			p.instructions[i].command = comNop
		} else if inst.command == comNop {
			p.instructions[i].command = comJmp
		}
		if !p.doesLoop() {
			return
		}
		p.reset()
	}
}

func (p *program) reset() {
	r := make([]instruction, len(p.init))
	copy(p.instructions, p.init)
	p.instructions = r
	p.acc = 0
	p.pc = 0
}

func (p *program) step() {
	switch com := p.instructions[p.pc]; com.command {
	case comAcc:
		p.acc += com.value
		p.pc++
	case comJmp:
		p.pc += com.value
	case comNop:
		p.pc++
	}
}
