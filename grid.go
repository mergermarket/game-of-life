package main

import (
	"strings"
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

func NewGridFromString(in string) *Grid {
	rows := strings.Split(strings.TrimSpace(in), "\n")
	height := len(rows)
	width := len(rows[0]) //todo: assuming all rows are same length, they wont be! return an error

	g := NewGrid(width, height)

	for rowNumber, row := range rows {
		for colNumber, col := range row {
			if col == '*' {
				g.grid[rowNumber][colNumber] = true
			}
		}
	}

	return g
}

func (g *Grid) String() string {
	var out string

	for i, _ := range g.grid{
		out +="\n"
		for _, y := range g.grid[i]{
			if y==true{
				out += "*"
			}else{
				out +="-"
			}
		}
	}
	out +="\n"
	return out
}
