package web

import "net/http"

// GameRenderer renders a game onto the supplied writer.
type GameRenderer interface {
	WriteGrid(writer http.ResponseWriter)
}

// Server returns a HTTP handler which renders a game.
func Server(renderer GameRenderer) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		renderer.WriteGrid(w)
	})
	return mux
}
