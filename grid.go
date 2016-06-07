package main

import (
	"errors"
	"strings"
)

var (
	ErrBadGridShape = errors.New("Grids must be rectangular")
	ErrBadGridChar  = errors.New("Characters in grid must either be - or *")
)

type Grid struct {
	width, height int
	grid          [][]bool
}

func NewGrid(width, height int) *Grid {
	g := new(Grid)
	g.width = width
	g.height = height

	grid := make([][]bool, height)

	for i := range grid {
		grid[i] = make([]bool, width)
	}

	g.grid = grid

	return g
}

func NewGridFromString(in string) (*Grid, error) {
	rows := strings.Split(strings.TrimSpace(in), "\n")
	height := len(rows)
	width := len(rows[0])

	g := NewGrid(width, height)

	for rowNumber, row := range rows {
		if len(row) != width {
			return nil, ErrBadGridShape
		}
		for colNumber, col := range row {
			if col == '*' {
				g.grid[rowNumber][colNumber] = true
			} else if col == '-' {
				g.grid[rowNumber][colNumber] = false
			} else {
				return nil, ErrBadGridChar
			}
		}
	}

	return g, nil
}

func (g *Grid) Step() {
	//todo
}

func (g *Grid) String() string {
	var out string

	for i, _ := range g.grid {
		out += "\n"
		for _, y := range g.grid[i] {
			if y == true {
				out += "*"
			} else {
				out += "-"
			}
		}
	}
	out += "\n"
	return out
}
