package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/mergermarket/game-of-life/game"
	"github.com/mergermarket/game-of-life/gif"
	"github.com/mergermarket/game-of-life/web"
)

// WriteGrid is.
func WriteGrid(writer http.ResponseWriter, size int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	grid := game.NewGrid(size, size)
	for x := range grid {
		for y := range grid {
			grid[x][y] = r.Int()%2 == 0
		}
	}

	myAwesomeGame := game.NewGame(grid)
	gif.WriteGrid(myAwesomeGame, writer)
}

func main() {
	mux := web.Server(WriteGrid)
	http.ListenAndServe(":8080", mux)
}
