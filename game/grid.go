package game

import (
	"errors"
	"strings"
)

var (
	ErrBadGridShape = errors.New("Grids must be rectangular")
	ErrBadGridChar  = errors.New("Characters in grid must either be - or *")
)

type grid [][]bool

type cell struct {
	x, y int
}

func (c cell) splat() (x int, y int) {
	return c.x, c.y
}

func NewGridFromString(in string) (grid, error) {
	rows := strings.Split(strings.TrimSpace(in), "\n")
	height := len(rows)
	width := len(rows[0])

	g := newGrid(width, height)

	for rowNumber, row := range rows {
		if len(row) != width {
			return nil, ErrBadGridShape
		}
		for colNumber, col := range row {
			if col == '*' {
				g[rowNumber][colNumber] = true
			} else if col == '-' {
				g[rowNumber][colNumber] = false
			} else {
				return nil, ErrBadGridChar
			}
		}
	}

	return g, nil
}

func (g grid) IsAlive(cell cell) bool {
	x, y := cell.splat()
	if x < 0 || y < 0 {
		return false
	}

	if x >= len(g) || y >= len(g[0]) {
		return false
	}

	return g[x][y]
}

func (g grid) KillCell(cell cell) {
	g[cell.x][cell.y] = false
}

func (g grid) ResurrectCell(cell cell) {
	g[cell.x][cell.y] = true
}

//todo: testme
func (g grid) GetCells() []cell {
	var cells []cell

	for x := range g {
		for y := range g[x] {
			cells = append(cells, cell{x, y})
		}
	}

	return cells
}

func (g grid) String() string {
	var out string

	for i := range g {
		out += "\n"
		for _, y := range g[i] {
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

func (g grid) Copy() GameWorld {
	cpy := newGrid(len(g), len(g[0]))

	for x := range g {
		for y, cell := range g[x] {
			cpy[x][y] = cell
		}
	}
	return cpy
}

func newGrid(width, height int) grid {
	g := make([][]bool, height)

	for i := range g {
		g[i] = make([]bool, width)
	}

	return g
}

func (g grid) GetAliveNeighbours(c cell) int {
	total := 0
	x, y := c.splat()

	if g.IsAlive(cell{x - 1, y - 1}) {
		total++
	}

	if g.IsAlive(cell{x, y - 1}) {
		total++
	}

	if g.IsAlive(cell{x + 1, y - 1}) {
		total++
	}

	if g.IsAlive(cell{x - 1, y}) {
		total++
	}

	if g.IsAlive(cell{x + 1, y}) {
		total++
	}

	if g.IsAlive(cell{x - 1, y + 1}) {
		total++
	}

	if g.IsAlive(cell{x, y + 1}) {
		total++
	}

	if g.IsAlive(cell{x + 1, y + 1}) {
		total++
	}

	return total
}

func (g grid) ToGrid() [][]bool {
	return g
}
