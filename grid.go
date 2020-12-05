package aoc

import (
	"fmt"
)

// Coordinate represents a two-dimensional (x,y) position on the grid.
type Coordinate struct {
	X, Y int
}

// Add adds Coordinate s to the given Coordinate
func (r Coordinate) Add(s Coordinate) Coordinate {
	return Coordinate{
		X: r.X + s.X,
		Y: r.Y + s.Y,
	}
}

// Subtract substracts Coordinate s from the given Coordinate.
func (r Coordinate) Subtract(s Coordinate) Coordinate {
	return Coordinate{
		X: r.X - s.X,
		Y: r.Y - s.Y,
	}
}

// Grid represents a two dimensional array of Stringers.  Since Grids may be printed, String() methods for elements stored in the grid should return a uniform length string (ideally len() == 1).
type Grid struct {
	Points                 map[Coordinate]fmt.Stringer
	minX, minY, maxX, maxY int
}

// Print displays the entire grid to STDOUT using the grid's designated MapperFunc
func (g Grid) Print() {
	output := [][]string{}
	h, w := g.maxY-g.minY, g.maxX-g.minX

	for i := 0; i <= h; i++ {
		row := []string{}
		for j := 0; j <= w; j++ {
			row = append(row, " ")
		}
		output = append(output, row)
	}

	for l, c := range g.Points {
		output[l.Y+Abs(g.minY)][l.X+Abs(g.minX)] = c.String()
	}

	for i := range output {
		for j := range output[i] {
			fmt.Print(output[i][j])
		}
		fmt.Print("\n")
	}
}
