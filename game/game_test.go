package game

import "testing"

func TestFewerThanTwoLiveCellsDiesOnGeneration(t *testing.T) {
	input := `
---
-*-
---
`
	grid, _ := NewGridFromString(input)
	game := NewGame(grid)

	game.Step()

	expected := `
---
---
---
`
	result := game.grid.String()
	if result != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(result)
		t.Error("Grid was not outputted as expected")
	}
}

func TestTwoByTwoSquareLivesToNextGeneration(t *testing.T) {
	input := `
---
-**
-**
`
	grid, _ := NewGridFromString(input)
	game := NewGame(grid)

	game.Step()

	expected := `
---
-**
-**
`
	result := game.grid.String()
	if result != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(result)
		t.Error("Grid was not outputted as expected")
	}
}

func TestTwoByDiagonalTripleDiesExceptCenterCell(t *testing.T) {
	input := `
--*
-*-
*--
`
	grid, _ := NewGridFromString(input)
	game := NewGame(grid)

	game.Step()

	expected := `
---
-*-
---
`
	result := game.grid.String()
	if result != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(result)
		t.Error("Grid was not outputted as expected")
	}
}

func TestCenterCellWith4NeighboursDies(t *testing.T) {
	input := `
*-*
***
*-*
`
	grid, _ := NewGridFromString(input)
	game := NewGame(grid)

	game.Step()

	expected := `
*-*
*-*
*-*
`
	result := game.grid.String()
	if result != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(result)
		t.Error("Grid was not outputted as expected")
	}
}

func TestDeadCenterCellWith3NeighboursResurrects(t *testing.T) {
	input := `
*-*
---
*--
`
	grid, _ := NewGridFromString(input)
	game := NewGame(grid)

	game.Step()

	expected := `
---
-*-
---
`
	result := game.grid.String()
	if result != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(result)
		t.Error("Grid was not outputted as expected")
	}
}
