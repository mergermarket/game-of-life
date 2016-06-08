package main

import (
	"errors"
	"strings"
)

var (
	ErrBadGridShape = errors.New("Grids must be rectangular")
	ErrBadGridChar  = errors.New("Characters in grid must either be - or *")
)

type Game struct {
	width, height int
	grid          [][]bool
}

func NewGame(width, height int) *Game {
	g := new(Game)
	g.width = width
	g.height = height

	g.grid = newGrid(g.width, g.height)

	return g
}

func NewGameFromString(in string) (*Game, error) {
	rows := strings.Split(strings.TrimSpace(in), "\n")
	height := len(rows)
	width := len(rows[0])

	g := NewGame(width, height)

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

func totalLivingCells(grid [][]bool) int {
	total := 0

	for i, _ := range grid {
		for _, y := range grid[i] {
			if y {
				total++
			}
		}
	}

	return total
}

func (g *Game) Step() {
	originalGrid := newGrid(g.height, g.width)
	for i := range originalGrid {
		copy(originalGrid[i], g.grid[i])
	}

	for i, _ := range originalGrid {
		for j, _ := range originalGrid[i] {
			if originalGrid[i][j] {
				if liveNeighbors(i, j, originalGrid) < 2 {
					g.grid[i][j] = false
				}
			}
		}
	}
}

func (g *Game) String() string {
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
