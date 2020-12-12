package aoc

import (
	"fmt"
)

// Coordinate represents a two-dimensional (x,y) position on the grid.
type Coordinate struct {
	X, Y int
}

// Add adds Coordinate s to the given Coordinate
func (c Coordinate) Add(s Coordinate) Coordinate {
	return Coordinate{
		X: c.X + s.X,
		Y: c.Y + s.Y,
	}
}

// Subtract substracts Coordinate s from the given Coordinate.
func (c Coordinate) Subtract(s Coordinate) Coordinate {
	return Coordinate{
		X: c.X - s.X,
		Y: c.Y - s.Y,
	}
}

// Neighbors returns a list of the eight coordinates directly adjacent to a given coordinate.
func (c Coordinate) Neighbors() []Coordinate {
	return []Coordinate{
		Coordinate{X: c.X, Y: c.Y - 1},     // North
		Coordinate{X: c.X + 1, Y: c.Y - 1}, // Northeast
		Coordinate{X: c.X + 1, Y: c.Y},     // East
		Coordinate{X: c.X + 1, Y: c.Y + 1}, // Southeast
		Coordinate{X: c.X, Y: c.Y + 1},     // South
		Coordinate{X: c.X - 1, Y: c.Y + 1}, // Southwest
		Coordinate{X: c.X - 1, Y: c.Y},     // West
		Coordinate{X: c.X - 1, Y: c.Y - 1}, // Northwest
	}
}

// Grid represents a two dimensional array of Stringers.  Since Grids may be printed, String() methods for elements stored in the grid should return a uniform length string (ideally len() == 1).
type Grid struct {
	Points                 map[Coordinate]fmt.Stringer
	minX, minY, maxX, maxY int
}

// Set stores a value at a given coordinate.
func (g *Grid) Set(c Coordinate, v fmt.Stringer) {
	if g.Points == nil {
		g.Points = make(map[Coordinate]fmt.Stringer)
	}

	g.Points[c] = v
	g.minX = Min(g.minX, c.X)
	g.minY = Min(g.minY, c.Y)
	g.maxX = Max(g.maxX, c.X)
	g.maxY = Max(g.maxY, c.Y)
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
