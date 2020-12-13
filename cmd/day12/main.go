package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	aoc "github.com/mleone10/advent-of-code-2020"
)

type action string
type value int
type rotationDir int
type heading int

const (
	actionNorth   = action("N")
	actionSouth   = action("S")
	actionEast    = action("E")
	actionWest    = action("W")
	actionLeft    = action("L")
	actionRight   = action("R")
	actionForward = action("F")
)

const (
	rotateRight = rotationDir(-1)
	rotateLeft  = rotationDir(1)
)

const (
	headingEast = heading(iota)
	headingNorth
	headingWest
	headingSouth
)

var headings = []heading{
	headingEast,
	headingNorth,
	headingWest,
	headingSouth,
}

var dirNorth = coordinate{x: 0, y: -1} // North
var dirSouth = coordinate{x: 0, y: 1}  // South
var dirEast = coordinate{x: 1, y: 0}   // East
var dirWest = coordinate{x: -1, y: 0}  // West

type coordinate struct {
	x, y int
}

type ship struct {
	head heading
	loc  coordinate
}

type wayship struct {
	s, wp ship
}

func main() {
	var s ship
	var ws wayship
	ws.wp.loc.x, ws.wp.loc.y = 10, -1
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		i := scanner.Text()
		a := action(i[0])
		v, _ := strconv.Atoi(i[1:])
		s.apply(a, value(v))
		ws.apply(a, value(v))
	}

	log.Printf("Normal ship manhattan distance from origin: %d", s.loc.x+s.loc.y)
	log.Printf("Wayship manhattan distance from origin: %d", aoc.Abs(ws.s.loc.x)+aoc.Abs(ws.s.loc.y))
}

func (s *ship) apply(a action, v value) {
	switch a {
	case actionNorth:
		s.move(dirNorth, v)
	case actionSouth:
		s.move(dirSouth, v)
	case actionEast:
		s.move(dirEast, v)
	case actionWest:
		s.move(dirWest, v)
	case actionLeft:
		s.turn(rotateLeft, v)
	case actionRight:
		s.turn(rotateRight, v)
	case actionForward:
		switch headings[s.head] {
		case headingNorth:
			s.move(dirNorth, v)
		case headingSouth:
			s.move(dirSouth, v)
		case headingEast:
			s.move(dirEast, v)
		case headingWest:
			s.move(dirWest, v)
		}
	}
}

func (s *ship) move(dir coordinate, val value) {
	s.loc = coordinate{x: s.loc.x + dir.x*int(val), y: s.loc.y + dir.y*int(val)}
}

func (s *ship) turn(dir rotationDir, val value) {
	s.head = headings[mod(int(s.head)+int(dir)*(int(val)/90), 4)]
}

func (ws *wayship) apply(a action, v value) {
	switch a {
	case actionNorth:
		ws.wp.move(dirNorth, v)
	case actionSouth:
		ws.wp.move(dirSouth, v)
	case actionEast:
		ws.wp.move(dirEast, v)
	case actionWest:
		ws.wp.move(dirWest, v)
	case actionLeft:
		ws.turn(rotateLeft, v)
	case actionRight:
		ws.turn(rotateRight, v)
	case actionForward:
		ws.move(v)
	}
}

func (ws *wayship) turn(dir rotationDir, val value) {
	for i := 0; i < int(val)/90; i++ {
		switch dir {
		case rotateLeft:
			ws.wp.loc.x, ws.wp.loc.y = ws.wp.loc.y, -1*ws.wp.loc.x
		case rotateRight:
			ws.wp.loc.x, ws.wp.loc.y = -1*ws.wp.loc.y, ws.wp.loc.x
		}
	}
}

func (ws *wayship) move(val value) {
	for i := 0; i < int(val); i++ {
		ws.s.loc.x = ws.s.loc.x + ws.wp.loc.x
		ws.s.loc.y = ws.s.loc.y + ws.wp.loc.y
	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}
