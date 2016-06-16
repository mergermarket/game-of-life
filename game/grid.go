package game

import "strings"

type grid [][]bool

type cell struct {
	x, y int
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

func (g grid) isAlive(x int, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if x >= len(g) || y >= len(g[0]) {
		return false
	}

	return g[x][y]
}

func (g grid) killCell(x int, y int) {
	g[x][y] = false
}

func (g grid) resurrectCell(x int, y int) {
	g[x][y] = true
}

//todo: testme
func (g grid) getCells() []cell {
	var cells []cell

	for x, _ := range g {
		for y, _ := range g[x] {
			cells = append(cells, cell{x, y})
		}
	}

	return cells
}

func (g grid) String() string {
	var out string

	for i, _ := range g {
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

func newGrid(width, height int) grid {
	g := make([][]bool, height)

	for i := range g {
		g[i] = make([]bool, width)
	}

	return g
}

func (g grid) cleanClone() grid{
	height := len(g)
	width := len(g[0])
	return newGrid(width, height)
}

func liveNeighbors(x int, y int, g grid) int {
	total := 0

	if g.isAlive(x-1, y-1) {
		total++
	}

	if g.isAlive(x, y-1) {
		total++
	}

	if g.isAlive(x+1, y-1) {
		total++
	}

	if g.isAlive(x-1, y) {
		total++
	}

	if g.isAlive(x+1, y) {
		total++
	}

	if g.isAlive(x-1, y+1) {
		total++
	}

	if g.isAlive(x, y+1) {
		total++
	}

	if g.isAlive(x+1, y+1) {
		total++
	}

	return total
}
