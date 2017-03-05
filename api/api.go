package api

import (
	"flag"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pressly/chi"
	"github.com/pressly/chi/docgen"
	"github.com/pressly/chi/middleware"
	"github.com/titouanfreville/popcubeapi/configs"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/utils"
)

type saveDb struct {
	db *gorm.DB
}

// Key type to be sure the context key is the one we want.
type key string

var (
	routes   = flag.Bool("routes", false, "Generate router documentation")
	dbStore  = saveDb{}
	error422 = utils.NewAPIError(422, "parse.request.body", "Request json object not correct.")
	error503 = utils.NewAPIError(503, "database.maintenance", "Database is currently in maintenance state. We are doing our best to get it back online ASAP.")
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
	// router.Use(middleware.Heartbeat("/heartbeat"))
	router.Use(middleware.CloseNotify)
}

// basicRoutes set basic routes for the API
func basicRoutes(router *chi.Mux) {
	// swagger:route GET / Test getter
	//
	// Hello World
	//
	// 	Responses:
	//    200
	// 	  default: genericError
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to PopCube api. Let's chat all together :O"))
	})
	// swagger:route GET /ping Test getter
	//
	// Pong
	//
	// Test api ping
	//
	// 	Responses:
	//    200
	// 	  default: genericError
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	// swagger:route GET /panic Test getter
	//
	// Should result in 500
	//
	// Test panic cautching
	//
	// 	Responses:
	//    500
	// 	  default: genericError
	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("C'est la panique, panique, panique. Sur le périphérique")
	})
}

// StartAPI initialise the api with provided host and port.
func StartAPI(hostname string, port string, DbConnectionInfo *configs.DbConnection) {
	router := newRouter()
	// Init DB connection
	user := DbConnectionInfo.User
	db := DbConnectionInfo.Database
	pass := DbConnectionInfo.Password
	host := DbConnectionInfo.Host
	dbport := DbConnectionInfo.Port
	dbStore.db = datastores.Store().InitConnection(user, db, pass, host, dbport)

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
	if *routes {
		log.Println(docgen.JSONRoutesDoc(router))
		log.Println(docgen.BuildDoc(router))
		log.Println(docgen.MarkdownRoutesDoc(router, docgen.MarkdownOpts{
			ProjectPath: "github.com/titouanfreville/popcubeapi",
			Intro:       "Welcomme to popcube user api.",
		}))
	}

	http.ListenAndServe(hostname+":"+port, router)
}
