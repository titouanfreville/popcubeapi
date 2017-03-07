package api

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"flag"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/goware/jwtauth"
	"github.com/jinzhu/gorm"
	"github.com/pressly/chi"
	"github.com/pressly/chi/docgen"
	"github.com/pressly/chi/middleware"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/configs"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
	"github.com/titouanfreville/popcubeapi/utils"
	renderPackage "github.com/unrolled/render"
)

type saveDb struct {
	db *gorm.DB
}

// Key type to be sure the context key is the one we want.
type key string

var (
	hmacSampleSecret []byte
	tokenAuth        *jwtauth.JwtAuth
	userToken        *jwt.Token
	encoding         = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")
	render           = renderPackage.New()
	routes           = flag.Bool("routes", false, "Generate router documentation")
	dbStore          = saveDb{}
	error401         = utils.NewAPIError(401, "unauthorized", "You did not login into the app. Please login to access those resources")
	error422         = utils.NewAPIError(422, "parse.request.body", "Request json object not correct.")
	error503         = utils.NewAPIError(503, "database.maintenance", "Database is currently in maintenance state. We are doing our best to get it back online ASAP.")
)

func newRandomString(length int) string {
	var b bytes.Buffer
	str := make([]byte, length+8)
	rand.Read(str)
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(str)
	encoder.Close()
	b.Truncate(length)
	return b.String()
}
func initAuth() {
	secret := newRandomString(100)
	hmacSampleSecret = []byte(secret)
	log.Print(secret)
	tokenAuth = jwtauth.New("HS256", hmacSampleSecret, hmacSampleSecret)
}

// createToken create JWT auth token for current login user
func createToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"name":  user.Username,
		"email": user.Email,
	}
	unsignedToken := *jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := unsignedToken.SignedString(hmacSampleSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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
	router.Use(middleware.Heartbeat("/heartbeat"))
	router.Use(middleware.CloseNotify)
}

// basicRoutes set basic routes for the API
func basicRoutes(router *chi.Mux) {
	// swagger:route GET / Test hello
	//
	// Hello World
	//
	// 	Responses:
	//    200: generalOk
	// 	  default: genericError
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to PopCube api. Let's chat all together :O"))
	})
	// swagger:route GET /ping Test ping
	//
	// Pong
	//
	// Test api ping
	//
	// 	Responses:
	//    200: generalOk
	// 	  default: genericError
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	router.Get("/heartbeat", func(w http.ResponseWriter, r *http.Request) {})
	// swagger:route GET /panic Test panic
	//
	// Should result in 500
	//
	// Test panic cautching
	//
	// 	Responses:
	//    500: genericError
	// 	  default: genericError
	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("C'est la panique, panique, panique. Sur le périphérique")
	})
	// swagger:route POST /login Login login
	//
	// Try to log user in
	//
	// Login user with provided USERNAME && Password
	//
	// Responses:
	// 		200: loginOk
	// 		404: incorrectIds
	// 	  422: wrongEntity
	// 	  503: databaseError
	router.Post("/login", loginMiddleware)
}

// loginMiddleware login funcion providing user && jwt auth token
func loginMiddleware(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Login    string      `json:"login"`
		Password string      `json:"password"`
		OmitID   interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	response := loginOk{}
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err == nil {
		user, err := store.User().Login(data.Login, data.Password, db)
		if err == nil {
			var terr error
			response.User = user
			response.Token, terr = createToken(user)
			if terr == nil {
				render.JSON(w, 200, response)
				return
			}
			render.JSON(w, err.StatusCode, err)
			return
		}
	}
	render.JSON(w, error503.StatusCode, error503)

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
	initAuth()
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
