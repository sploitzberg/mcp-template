package http

import (
	"net/http"
	"strings"
)

// Router builds an http.ServeMux for the resource API.
func Router(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /resources", h.Create)
	mux.HandleFunc("GET /resources/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/resources/")
		if id == "" {
			http.NotFound(w, r)
			return
		}
		h.GetByID(w, r, id)
	})
	return mux
}
