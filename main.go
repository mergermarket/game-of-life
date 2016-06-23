package main

import (
	"os"

	"github.com/mergermarket/game-of-life/game"
	"github.com/mergermarket/game-of-life/gif"
)

func main() {

	input := `
---
***
---
`

	grid, _ := game.NewGridFromString(input)
	myAwesomeGame := game.NewGame(grid)
	gif.Lissajous(myAwesomeGame, os.Stdout)
}
