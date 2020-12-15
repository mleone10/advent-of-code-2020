package main

import "log"

const (
	targetLow  = 2020
	targetHigh = 30000000
)

type history map[int]int
type game struct {
	init      []int
	hist      history
	turn, say int
}

func main() {
	g := initGame(0, 1, 4, 13, 15, 12, 16)

	log.Printf("Last number spoken on turn %d: %d", targetLow, g.saidOnTurnN(targetLow))
	log.Printf("Last number spoken on turn %d: %d", targetHigh, g.saidOnTurnN(targetHigh))
}

func initGame(ns ...int) *game {
	g := game{}
	for _, n := range ns {
		g.init = append(g.init, n)
	}
	g.reset()

	return &g
}

func (g *game) saidOnTurnN(n int) int {
	g.reset()
	for g.turn < n {
		g.step()
	}
	return g.say
}

func (g *game) step() {
	say := 0
	if t, ok := g.hist[g.say]; ok {
		say = g.turn - t
	}
	g.load(say)
}

func (g *game) reset() {
	g.hist = history{}
	g.turn = 1
	g.say = g.init[0]
	for _, n := range g.init[1:] {
		g.load(n)
	}
}

func (g *game) load(n int) {
	g.hist[g.say] = g.turn
	g.say = n
	g.turn++
}
