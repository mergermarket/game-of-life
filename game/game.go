package game

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
	grid          grid
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

func (g *Game) Step() {
	nextGrid := newGrid(g.height, g.width)

	for x, _ := range g.grid {
		for y, _ := range g.grid[x] {
			neighbors := liveNeighbors(x, y, g.grid)
			nextGrid[x][y] = g.grid[x][y]
			if g.grid.isAlive(x, y) && neighbors < 2 {
				nextGrid.killCell(x, y)
			}
			if g.grid.isAlive(x, y) && neighbors > 3 {
				nextGrid.killCell(x, y)
			}
			if g.grid.isAlive(x, y) && (neighbors == 2 || neighbors == 3) {
				// Survives
			}
			if !g.grid.isAlive(x, y) && neighbors == 3 {
				nextGrid.resurrectCell(x, y)
			}
		}
	}
	g.grid = nextGrid
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
