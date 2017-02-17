package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/titouanfreville/popcubeapi/datastores"
)

type testDb struct {
	db *gorm.DB
}

var dbStore = testDb{}

// newRouter initialise api serveur.
func newRouter() *chi.Mux {
	return chi.NewRouter()
}

// initMiddleware initialise middlewares for router
func initMiddleware(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// When a client closes their connection midway through a request, the
	// http.CloseNotifier will cancel the request context (ctx).
	router.Use(middleware.CloseNotify)
}

// basicRoutes set basic routes for the API
func basicRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to PopCube api. Let's chat all together :O"))
	})

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("C'est la panique, panique, panique. Sur le périphérique")
	})
}

// StartAPI initialise the api with provided host and port.
func StartAPI(hostname string, port string) {
	router := newRouter()
	dbStore.db = datastores.NewStore().InitConnection("root", "popcube_test", "popcube_dev", "0.0.0.0", "3306")
	initMiddleware(router)
	basicRoutes(router)
	initAvatarRoute(router)
	http.ListenAndServe(hostname+":"+port, router)
}
