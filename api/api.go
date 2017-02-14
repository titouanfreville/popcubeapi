package api

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/titouanfreville/popcubeapi/datastores"
)

// Base store Router and dataStore
type Base struct {
	ds     *datastores.DbStore
	router *chi.Mux
}

// BaseConst provide acces to APIBase from all the package api
var BaseConst = Base{}

// newRouter initialise api serveur.
func newRouter() *chi.Mux {
	return chi.NewRouter()
}

// initMiddleware initialise middlewares for router
func (ll Base) initMiddleware() {
	ll.router.Use(middleware.RequestID)
	ll.router.Use(middleware.RealIP)
	ll.router.Use(middleware.Logger)
	ll.router.Use(middleware.Recoverer)
	// When a client closes their connection midway through a request, the
	// http.CloseNotifier will cancel the request context (ctx).
	ll.router.Use(middleware.CloseNotify)
}

// basicRoutes set basic routes for the API
func (ll Base) basicRoutes() {
	ll.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to PopCube api. Let's chat all together :O"))
	})

	ll.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	ll.router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})
}

// StartAPI initialise the api with provided host and port.
func (ll Base) StartAPI(hostname string, port string, ds *datastores.DbStore) {
	ll.ds = ds
	ll.router = newRouter()
	ll.initMiddleware()
	ll.basicRoutes()
	ll.initAvatarRoute()
	http.ListenAndServe(hostname+":"+port, ll.router)
}
