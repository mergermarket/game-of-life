package game

type grid [][]bool

type cell struct {
	x, y int
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

func newGrid(width, height int) grid {
	g := make([][]bool, height)

	for i := range g {
		g[i] = make([]bool, width)
	}

	return g
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
