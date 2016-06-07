package main

import "testing"

func TestInitialisesAnEmptyGrid(t *testing.T) {
	width := 5
	height := 8
	grid := NewGrid(width, height)

	if grid.grid[height-1][width-1] != false {
		t.Error("Expected bottom right corner to be set to false")
	}

	if len(grid.grid) != height {
		t.Error("Incorrect value for height, got", len(grid.grid), "expected", height)
	}

	if len(grid.grid[0]) != width {
		t.Error("Incorrect value for width, got", len(grid.grid[0]), "expected", width)
	}
}

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

	if grid.grid[0][1] != true {
		t.Error("Expected grid item to be set to true but it wasnt", grid.grid)
	}

}

func TestItRejectsNonRectangularGrids(t *testing.T){
	input := `
-*-----
---
---
`
	_, err := NewGridFromString(input)

	if err == nil{
		t.Error("Expected an error")
	}
}

func TestItCanOutputStringGrid(t *testing.T) {
	input := `
---
---
---
`
	grid, _ := NewGridFromString(input)

	grid.grid[0][0] = true

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
