package web

import (
	"net/http"
	"strconv"
)

const defaultGridSize = 100

// Server returns a HTTP handler which renders a game.
func Server(render func(http.ResponseWriter, int)) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		sizeParam := req.URL.Query().Get("size")
		var size int
		if sizeParam == "" {
			size = defaultGridSize
		} else {
			size, _ = strconv.Atoi(sizeParam)
		}
		render(w, size)
	})
	return mux
}
