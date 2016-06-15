package game

import "testing"

func TestIsAlive(t *testing.T) {
	grid := newGrid(3, 3)

	grid[1][1] = true

	if !grid.isAlive(1, 1) {
		t.Error("Expected grid to be alive here")
	}
}

func TestNegativeIndexIsDead(t *testing.T) {
	grid := newGrid(3, 3)

	if grid.isAlive(-1, 1) {
		t.Error("Expected grid to be dead here")
	}

	if grid.isAlive(1, -1) {
		t.Error("Expected grid to be dead here")
	}
}

func TestOutOfBoundsIndexIsDead(t *testing.T) {
	grid := newGrid(3, 3)

	if grid.isAlive(2, 3) {
		t.Error("Expected grid to be dead here")
	}

	if grid.isAlive(3, 2) {
		t.Error("Expected grid to be dead here")
	}
}

func TestLiveNeighborsNone(t *testing.T) {
	g := newGrid(1, 1)

	if liveNeighbors(0, 0, g) != 0 {
		t.Error("Expected no live neighbours")
	}
}

func TestLiveNeighborsTwo(t *testing.T) {
	g := newGrid(3, 3)
	g[1][1] = true
	g[2][2] = true
	g[0][0] = true

	if liveNeighbors(1, 1, g) != 2 {
		t.Error("Expected two live neighbours")
	}
	if liveNeighbors(0, 0, g) != 1 {
		t.Error("Expected one live neighbours")
	}
	if liveNeighbors(2, 2, g) != 1 {
		t.Error("Expected one live neighbours")
	}
	if liveNeighbors(0, 2, g) != 1 {
		t.Error("Expected one live neighbours")
	}
}

func TestLiveNeighbors(t *testing.T) {
	g := newGrid(3, 3)
	g[0][1] = true
	g[1][0] = true
	g[1][2] = true
	g[2][1] = true

	if liveNeighbors(0, 0, g) != 2 {
		t.Error("Expected two live neighbours")
	}
	if liveNeighbors(0, 1, g) != 2 {
		t.Error("Expected two live neighbours")
	}
	if liveNeighbors(0, 2, g) != 2 {
		t.Error("Expected two live neighbours")
	}
	if liveNeighbors(1, 0, g) != 2 {
		t.Error("Expected two live neighbours")
	}
	if liveNeighbors(1, 1, g) != 4 {
		t.Error("Expected four live neighbours")
	}
	if liveNeighbors(1, 2, g) != 2 {
		t.Error("Expected two live neighbours")
	}
	if liveNeighbors(2, 0, g) != 2 {
		t.Error("Expected two live neighbours")
	}
	if liveNeighbors(2, 1, g) != 2 {
		t.Error("Expected two live neighbours")
	}
	if liveNeighbors(2, 2, g) != 2 {
		t.Error("Expected two live neighbours")
	}
}
