package api

import (
	"net/http"

	"github.com/pressly/chi"
)

// NewRouter initialise api serveur.
func NewRouter() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
