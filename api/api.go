package api

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

// newRouter initialise api serveur.
func newRouter() *chi.Mux {
	return chi.NewRouter()
}

// initMiddleware initialise middlewares for router
func initMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// A good base middleware stack
	// When a client closes their connection midway through a request, the
	// http.CloseNotifier will cancel the request context (ctx).
	r.Use(middleware.CloseNotify)
}

// basicRoutes set basic routes for the API
func basicRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to PopCube api. Let's chat all together :O"))
	})
}

// StartAPI initialise the api with provided host and port.
func StartAPI(hostname string, port string) {
	r := newRouter()
	initMiddleware(r)
	basicRoutes(r)
	http.ListenAndServe(hostname+":"+port, r)
}
