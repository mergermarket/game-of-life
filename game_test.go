package main

import "testing"

func TestInitialisesAnEmptyGrid(t *testing.T) {
	width := 5
	height := 8
	game := NewGame(width, height)

	if game.grid[height-1][width-1] != false {
		t.Error("Expected bottom right corner to be set to false")
	}

	if len(game.grid) != height {
		t.Error("Incorrect value for height, got", len(game.grid), "expected", height)
	}

	if len(game.grid[0]) != width {
		t.Error("Incorrect value for width, got", len(game.grid[0]), "expected", width)
	}
}

func TestItCanCreateGridFromString(t *testing.T) {
	input := `
-*-
---
---
`
	game, _ := NewGameFromString(input)

	if game == nil {
		t.Error("Expected a grid to be returned but got nil")
	}

	if game.grid[0][1] != true {
		t.Error("Expected grid item to be set to true but it wasnt", game.grid)
	}

}

func TestItCanOutputStringGrid(t *testing.T) {
	input := `
---
---
---
`
	game, _ := NewGameFromString(input)

	game.grid[0][0] = true

	expected := `
*--
---
---
`
	if game.String() != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(game.String())
		t.Error("Grid was not outputted as expected")
	}
}

func TestFewerThanTwoLiveCellsDiesOnGeneration(t *testing.T) {
	input := `
---
-*-
---
`
	game, _ := NewGameFromString(input)

	game.Step()

	expected := `
---
---
---
`
	if game.String() != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(game.String())
		t.Error("Grid was not outputted as expected")
	}
}

func TestTwoByTwoSquareLivesToNextGeneration(t *testing.T) {
	input := `
---
-**
-**
`
	game, _ := NewGameFromString(input)

	game.Step()

	expected := `
---
-**
-**
`
	if game.String() != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(game.String())
		t.Error("Grid was not outputted as expected")
	}
}

func TestTwoByDiagonalTripleDiesExceptCenterCell(t *testing.T) {
	input := `
--*
-*-
*--
`
	game, _ := NewGameFromString(input)

	game.Step()

	expected := `
---
-*-
---
`
	if game.String() != expected {
		t.Log("Expected")
		t.Log(expected)
		t.Log("Got")
		t.Log(game.String())
		t.Error("Grid was not outputted as expected")
	}
}

func TestItRejectsNonRectangularGrids(t *testing.T) {
	input := `
-*-----
---
---
`
	_, err := NewGameFromString(input)

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
	_, err := NewGameFromString(input)

	if err == nil {
		t.Error("Expected an error")
	}

	if err != ErrBadGridChar {
		t.Error("Expected", ErrBadGridChar, "but got", err)
	}
}
