package game

import "testing"

func TestItCanCreateGridFromString(t *testing.T) {
	input := `
-*-
---
---
`
	grid, _ := NewGridFromString(input)

	if grid == nil {
		t.Error("Expected a grid to be returned but got nil")
	}

	if grid[0][1] != true {
		t.Error("Expected grid item to be set to true but it wasnt", grid)
	}
}

func TestItCanCreateGridFromStringAgain(t *testing.T) {
	input := `
*--
-*-
--*
`
	grid, _ := NewGridFromString(input)

	if grid == nil {
		t.Error("Expected a grid to be returned but got nil")
	}
	if !(grid[0][0] && grid[1][1] && grid[2][2]) {
		t.Error("Expected grid item to be set to true but it wasnt", grid)
	}

	if grid[0][1] && grid[0][2] &&
		grid[1][0] && grid[1][2] &&
		grid[2][0] && grid[2][1] {
		t.Error("Expected grid item to be set to false but it wasnt", grid)
	}
}

func TestItRejectsNonRectangularGrids(t *testing.T) {
	input := `
-*-----
---
---
`
	_, err := NewGridFromString(input)

	if err == nil {
		t.Error("Expected an error")
	}

	if err != ErrBadGridShape {
		t.Error("Expected", ErrBadGridShape, "but got", err)
	}
}

func TestItRejectsUnexpectedCharsInGrids(t *testing.T) {
	input := `
---
-2-
---
`
	_, err := NewGridFromString(input)

	if err == nil {
		t.Error("Expected an error")
	}

	if err != ErrBadGridChar {
		t.Error("Expected", ErrBadGridChar, "but got", err)
	}
}

func TestItCanOutputStringGrid(t *testing.T) {
	input := `
---
---
---
`
	grid, _ := NewGridFromString(input)

	grid[0][0] = true

	expected := `
*--
---
---
`
	if grid.String() != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(grid.String())
		t.Error("Grid was not outputted as expected")
	}
}

func TestIsAlive(t *testing.T) {
	grid := NewGrid(3, 3)

	grid[1][1] = true

	if !grid.IsAlive(cell{1, 1}) {
		t.Error("Expected grid to be alive here")
	}
}

func TestNegativeIndexIsDead(t *testing.T) {
	grid := NewGrid(3, 3)

	if grid.IsAlive(cell{-1, 1}) {
		t.Error("Expected grid to be dead here")
	}

	if grid.IsAlive(cell{1, -1}) {
		t.Error("Expected grid to be dead here")
	}
}

func TestOutOfBoundsIndexIsDead(t *testing.T) {
	grid := NewGrid(3, 3)

	if grid.IsAlive(cell{2, 3}) {
		t.Error("Expected grid to be dead here")
	}

	if grid.IsAlive(cell{3, 2}) {
		t.Error("Expected grid to be dead here")
	}
}

func TestLiveNeighborsNone(t *testing.T) {
	g := NewGrid(1, 1)

	if g.GetAliveNeighbours(cell{0, 0}) != 0 {
		t.Error("Expected no live neighbours")
	}
}

func TestLiveNeighborsTwo(t *testing.T) {
	g := NewGrid(3, 3)
	g[1][1] = true
	g[2][2] = true
	g[0][0] = true

	if g.GetAliveNeighbours(cell{1, 1}) != 2 {
		t.Error("Expected two live neighbours")
	}
	if g.GetAliveNeighbours(cell{0, 0}) != 1 {
		t.Error("Expected one live neighbours")
	}
	if g.GetAliveNeighbours(cell{2, 2}) != 1 {
		t.Error("Expected one live neighbours")
	}
	if g.GetAliveNeighbours(cell{0, 2}) != 1 {
		t.Error("Expected one live neighbours")
	}
}

func TestLiveNeighbors(t *testing.T) {
	g := NewGrid(3, 3)
	g[0][1] = true
	g[1][0] = true
	g[1][2] = true
	g[2][1] = true

	if g.GetAliveNeighbours(cell{0, 0}) != 2 {
		t.Error("Expected two live neighbours")
	}
	if g.GetAliveNeighbours(cell{0, 1}) != 2 {
		t.Error("Expected two live neighbours")
	}
	if g.GetAliveNeighbours(cell{0, 2}) != 2 {
		t.Error("Expected two live neighbours")
	}
	if g.GetAliveNeighbours(cell{1, 0}) != 2 {
		t.Error("Expected two live neighbours")
	}
	if g.GetAliveNeighbours(cell{1, 1}) != 4 {
		t.Error("Expected four live neighbours")
	}
	if g.GetAliveNeighbours(cell{1, 2}) != 2 {
		t.Error("Expected two live neighbours")
	}
	if g.GetAliveNeighbours(cell{2, 0}) != 2 {
		t.Error("Expected two live neighbours")
	}
	if g.GetAliveNeighbours(cell{2, 1}) != 2 {
		t.Error("Expected two live neighbours")
	}
	if g.GetAliveNeighbours(cell{2, 2}) != 2 {
		t.Error("Expected two live neighbours")
	}
}
