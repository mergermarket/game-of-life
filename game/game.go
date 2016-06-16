package game

import (
	"errors"
)

var (
	ErrBadGridShape = errors.New("Grids must be rectangular")
	ErrBadGridChar  = errors.New("Characters in grid must either be - or *")
)

type Game struct {
	grid          grid
}

func NewGame(grid grid) *Game {
	g := new(Game)
	g.grid = grid
	return g
}

func (g *Game) Step() {
	nextGrid := g.grid.cleanClone()

	for _, cell := range g.grid.getCells() {
		x := cell.x
		y := cell.y
		neighbors := liveNeighbors(cell, g.grid)
		nextGrid[x][y] = g.grid[x][y]

		if g.grid.isAlive(cell) && neighbors < 2 {
			nextGrid.killCell(x, y)
		}
		if g.grid.isAlive(cell) && neighbors > 3 {
			nextGrid.killCell(x, y)
		}
		if g.grid.isAlive(cell) && (neighbors == 2 || neighbors == 3) {
			// Survives
		}
		if !g.grid.isAlive(cell) && neighbors == 3 {
			nextGrid.resurrectCell(x, y)
		}
	}
	g.grid = nextGrid
}
