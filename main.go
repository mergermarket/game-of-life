package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/mergermarket/game-of-life/game"
	"github.com/mergermarket/game-of-life/gif"
)

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	grid := game.NewGrid(201, 201)

	for x, _ := range grid {
		for y, _ := range grid {
			grid[x][y] = r.Int()%2 == 0
		}
	}

	myAwesomeGame := game.NewGame(grid)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		// fmt.Fprintf(w, "Welcome to the home page!")
		gif.Lissajous(myAwesomeGame, w)

	})
	http.ListenAndServe(":8080", mux)
}
