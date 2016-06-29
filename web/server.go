package web

import (
	"net/http"
	"strconv"
)

// Server returns a HTTP handler which renders a game.
func Server(render func(http.ResponseWriter, int)) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		size, _ := strconv.Atoi(req.URL.Query().Get("size"))
		render(w, size)
	})
	return mux
}
