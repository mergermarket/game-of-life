package main

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
