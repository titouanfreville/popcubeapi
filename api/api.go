package api

import (
	"flag"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pressly/chi"
	"github.com/pressly/chi/docgen"
	"github.com/pressly/chi/middleware"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/utils"
)

type testDb struct {
	db *gorm.DB
}

type deleteMessage struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Object  interface{} `json:"removed_object, omitempty"`
}

var (
	routes   = flag.Bool("routes", false, "Generate router documentation")
	dbStore  = testDb{}
	error422 = utils.NewApiError(422, "parse.request.body", "Request json object not correct.")
	error503 = utils.NewApiError(503, "database.maintenance", "Database is currently in maintenance state. We are doing our best to get it back online ASAP.")
)

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
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Timeout(5 * 1000))
	router.Use(middleware.Heartbeat("/ping"))
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
	initChannelRoute(router)
	initEmojiRoute(router)
	initFolderRoute(router)
	initMessageRoute(router)
	initOrganisationRoute(router)
	initParameterRoute(router)
	initRoleRoute(router)
	initUserRoute(router)

	// Passing -routes to the program will generate docs for the above
	// router definition. See the `routes.json` file in this folder for
	// the output.
	log.Println(docgen.JSONRoutesDoc(router))
	log.Println(docgen.BuildDoc(router))
	log.Println(docgen.MarkdownRoutesDoc(router, docgen.MarkdownOpts{
		ProjectPath: "github.com/titouanfreville/popcubeapi",
		Intro:       "Welcomme to popcube user api.",
	}))

	http.ListenAndServe(hostname+":"+port, router)
}
