package main

import "log"

type game struct {
	init, hist []int
}

func main() {
	g := initGame(0, 1, 4, 13, 15, 12, 16)

	for len(g.hist) < 2020 {
		g.step()
	}
	log.Printf("Last number spoken on turn 2020 %d", g.last())

	g.reset()

	for len(g.hist) < 30000000 {
		g.step()
	}
	log.Printf("Last number spoken on turn 30000000: %d", g.last())
}

func initGame(ns ...int) *game {
	g := game{}
	for _, n := range ns {
		g.init = append(g.init, n)
	}
	g.reset()

	return &g
}

func (g *game) step() {
	p := g.last()
	ts := g.turnIndexes(p)
	say := 0
	if len(ts) != 1 {
		say = ts[len(ts)-1] - ts[len(ts)-2]
	}
	g.load(say)
}

func (g *game) turnIndexes(n int) []int {
	ts := []int{}
	for i, m := range g.hist {
		if m == n {
			ts = append(ts, i)
		}
	}
	return ts
}

func (g *game) reset() {
	g.hist = make([]int, len(g.init))
	copy(g.hist, g.init)
}

func (g *game) load(n int) {
	g.hist = append(g.hist, n)
}

func (g *game) last() int {
	return g.hist[len(g.hist)-1]
}
