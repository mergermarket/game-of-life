package main

import (
	"os"

	"github.com/mergermarket/game-of-life/game"
	"github.com/mergermarket/game-of-life/gif"
	"math/rand"
	"time"
)

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	grid := game.NewGrid(100, 100)

	for x, _ := range grid {
		for y, _ := range grid {
			grid[x][y] = r.Int()%2 == 0
		}
	}

	myAwesomeGame := game.NewGame(grid)
	gif.Lissajous(myAwesomeGame, os.Stdout)
}
