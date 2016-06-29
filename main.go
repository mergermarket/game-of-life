package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/mergermarket/game-of-life/game"
	"github.com/mergermarket/game-of-life/gif"
	"github.com/mergermarket/game-of-life/web"
)

// UltimateGame is.
type UltimateGame struct {
}

// WriteGrid is.
func (u UltimateGame) WriteGrid(writer http.ResponseWriter) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	grid := game.NewGrid(201, 201)
	for x := range grid {
		for y := range grid {
			grid[x][y] = r.Int()%2 == 0
		}
	}

	myAwesomeGame := game.NewGame(grid)
	gif.WriteGrid(myAwesomeGame, writer)
}

func main() {
	game := UltimateGame{}
	mux := web.Server(game)
	http.ListenAndServe(":8080", mux)
}
