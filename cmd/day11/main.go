package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	seatFloor = seat(".")
	seatEmpty = seat("L")
	seatOcc   = seat("#")
)

var neighborVectors = []coordinate{
	coordinate{x: 0, y: -1},  // North
	coordinate{x: 1, y: -1},  // Northeast
	coordinate{x: 1, y: 0},   // East
	coordinate{x: 1, y: 1},   // Southeast
	coordinate{x: 0, y: 1},   // South
	coordinate{x: -1, y: 1},  // Southwest
	coordinate{x: -1, y: 0},  // West
	coordinate{x: -1, y: -1}, // Northwest
}

type coordinate struct {
	x, y int
}

type seat string
type seats map[coordinate]seat
type neighborCounter func(coordinate) int

func main() {
	ss := make(seats)
	scanner := bufio.NewScanner(os.Stdin)

	var y int
	for scanner.Scan() {
		for x, c := range strings.Split(scanner.Text(), "") {
			coord := coordinate{
				x: x,
				y: y,
			}
			ss[coord] = seat(c)
		}
		y++
	}

	temps := make(seats)

	temps.reset(ss)
	log.Printf("Number of occupied seats after reaching stable state (immediate neighbors): %d", temps.runToStable(temps.adjNeighbors, 4))

	temps.reset(ss)
	log.Printf("Number of occupied seats after reaching stable state (visible neighbors): %d", temps.runToStable(temps.visNeighbors, 5))
}

func (ss seats) runToStable(neighborCounter neighborCounter, numToEmpty int) int {
	var stable bool
	for !stable {
		stable = ss.step(neighborCounter, numToEmpty)
	}

	return ss.numOccupied()
}

func (ss seats) reset(os seats) {
	for k, v := range os {
		ss[k] = v
	}
}

func (ss seats) step(neighborCounter neighborCounter, numToEmpty int) bool {
	nss := make(seats)
	for c, s := range ss {
		occupied := neighborCounter(c)
		nss[c] = s
		switch s {
		case seatEmpty:
			if occupied == 0 {
				nss[c] = seatOcc
			}
		case seatOcc:
			if occupied >= numToEmpty {
				nss[c] = seatEmpty
			}
		}
	}
	same := true
	for c, s := range nss {
		if ss[c] != nss[c] {
			same = false
		}
		ss[c] = s
	}

	return same
}

func (ss seats) adjNeighbors(c coordinate) int {
	var sum int
	for _, v := range neighborVectors {
		coord := coordinate{
			x: c.x + v.x,
			y: c.y + v.y,
		}
		if ss[coord] == seatOcc {
			sum++
		}
	}
	return sum
}

func (ss seats) visNeighbors(c coordinate) int {
	var sum int
	for _, v := range neighborVectors {
		var done bool
		for i := 1; !done; i++ {
			s, ok := ss[coordinate{x: c.x + v.x*i, y: c.y + v.y*i}]
			if ok && s != seatFloor {
				if s == seatOcc {
					sum++
				}
			}
			if !ok || s != seatFloor {
				done = true
			}
		}
	}
	return sum
}

func (ss seats) numOccupied() int {
	var sum int
	for _, s := range ss {
		if s == seatOcc {
			sum++
		}
	}
	return sum
}
