package api

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base32"
	"flag"
	"log"
	"net/http"
	"os"
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/pressly/chi"
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

// Token A JWT Token.  Different fields will be used depending on whether you're
// creating or parsing/verifying a token.
// type Token struct {
// 	Raw       string                 // The raw token.  Populated when you Parse a token
// 	Method    SigningMethod          // The signing method used or to be used
// 	Header    map[string]interface{} // The first segment of the token
// 	Claims    Claims                 // The second segment of the token
// 	Signature string                 // The third segment of the token.  Populated when you Parse a token
// 	Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
// }

var (
	apiVersionKey    key = "version"
	secret           string
	hmacSampleSecret []byte
	tokenAuth        *JwtAuth
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
	hmacSampleSecret = []byte(secret)
	tokenAuth = New("HS256", hmacSampleSecret, hmacSampleSecret)
	claims := jwt.MapClaims{
		"organisation_name":   "PopCube",
		"organisation_stack":  1,
		"organisation_domain": "chat.popcube.xyz",
		"public":              false,
		"owner":               "TCMJJ",
		"owner_mail":          "owned@popcube.xyz",
		"owner_password":      "popcube",
		"type":                "neworganisation",
		"authorise":           "this token let you create new organisation and a new user in an iner DB",
	}
	unsignedToken := *jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := unsignedToken.SignedString(hmacSampleSecret)
	log.Print("Tolken new organisation && role : ")
	log.Print(tokenString)
}

// createUserToken create JWT auth token for current login user
func createUserToken(user models.User, role models.Role) (string, error) {
	claims := jwt.MapClaims{
		"name":  user.Username,
		"email": user.Email,
		"type":  "userauth",
	}
	unsignedToken := *jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := unsignedToken.SignedString(hmacSampleSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// createInviteToken create JWT auth token for current invitation
func createInviteToken(inviteMail string, organisationName string) (string, error) {
	claims := jwt.MapClaims{
		"email":        inviteMail,
		"organisation": organisationName,
		"type":         "invitation",
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

// initVersionRouting manage Version routing through go serveur
func initVersionRouting(router *chi.Mux) {
	router.Route("/alpha", func(router chi.Router) {
		router.Use(apiVersionContext("alpha"))
		initAvatarRoute(router)
		initChannelRoute(router)
		initEmojiRoute(router)
		initFolderRoute(router)
		initMessageRoute(router)
		initOrganisationRoute(router)
		initParameterRoute(router)
		initRoleRoute(router)
		initUserRoute(router)
		initDevGetter(router)
	})
}

// apiVersionContext Set Current ctx api version
func apiVersionContext(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), apiVersionKey, version))
			next.ServeHTTP(w, r)
		})
	}
}

// basicRoutes set basic routes for the API
func basicRoutes(router *chi.Mux) {
	router.Use(tokenAuth.Verifier)
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
	// 	  default: genericError
	router.Post("/login", loginMiddleware)
	// swagger:route POST /initorganisation Init initOrganisation
	//
	// Try to log user in
	//
	// Login user with provided USERNAME && Password
	//
	// Responses:
	// 		200: initOk
	// 	  503: databaseError
	// 	  default: genericError
	router.Post("/initorganisation", initOrganisation)
	router.Route("/publicuser", func(r chi.Router) {
		// swagger:route POST /publicuser/new Users newPublicUser
		//
		// New user
		//
		// This will create an user for organisation if organisation is Public OR Email match parametetered emails
		//
		// 	Responses:
		//    201: userObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newPublicUser)
		r.Route("/newfrominvite", func(r chi.Router) {
			r.Use(tokenAuth.Verifier)
			r.Use(allowUserCreationFromToken)
			// swagger:route POST /publicuser/newfrominvite Users newInvitedUser
			//
			// New user
			//
			// This will create an user for organisation if user was invited
			//
			// 	Responses:
			//    201: userObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Post("/", newUser)
		})
	})
}

func initDevGetter(router chi.Router) {
	env := os.Getenv("POPCUBE_API_ENV")
	if env == "prod" || env == "test" || env == "beta" || env == "alpha" || env == "production" {
		return
	}
	log.Print("<><><><><><><> Using DEV routes <><><><><><><> \n")
	router.Route("/devgetters", func(r chi.Router) {
		r.Get("/avatar", getAllAvatar)
		r.Get("/channel", getAllChannel)
		r.Get("/emoji", getAllEmoji)
		r.Get("/folder", getAllFolder)
		r.Get("/member", getAllMember)
		r.Get("/message", getAllMessage)
		r.Get("/organisation", getAllOrganisation)
		r.Get("/parameter", getAllParameter)
		r.Get("/role", getAllRole)
		r.Get("/user", getAllUser)
	})
}

// loginRequestObject
type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (lR *loginRequest) Bind(r *http.Request) error {
	return nil
}

// loginMiddleware login funcion providing user && jwt auth token
func loginMiddleware(w http.ResponseWriter, r *http.Request) {
	// var data struct {
	// 	Login    string      `json:"login"`
	// 	Password string      `json:"password"`
	// 	OmitID   interface{} `json:"id,omitempty"`
	// }
	store := datastores.Store()
	response := loginOk{}
	db := dbStore.db
	data := &loginRequest{}
	err := chiRender.Bind(r, data)
	if err != nil {
		log.Print("422 Here - loginMiddleware")
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err == nil {
		user, err := store.User().Login(data.Login, data.Password, db)
		if err == nil {
			var terr error
			// role can't be empty if user exist => foreign key constraint
			role := datastores.Store().Role().GetByID(user.IDRole, dbStore.db)
			response.User = user
			response.Token, terr = createUserToken(user, role)
			if terr == nil {
				render.JSON(w, 200, response)
				return
			}
			render.JSON(w, 422, "Could not generate token")
		}
		render.JSON(w, err.StatusCode, err)
		return
	}
	render.JSON(w, error503.StatusCode, error503)

}

func initOrganisation(w http.ResponseWriter, r *http.Request) {
	// Verify token
	ctx := r.Context()
	if jwtErr, ok := ctx.Value(jwtErrorKey).(error); ok {
		if jwtErr != nil {
			render.JSON(w, 401, "Token not found. You Are not allowed to proceed without token.")
			return
		}
	}
	jwtToken, ok := ctx.Value(jwtTokenKey).(*jwt.Token)
	if !ok || jwtToken == nil || !jwtToken.Valid {
		render.JSON(w, 401, "token is not valid or does not exist")
		return
	}
	tokenType, ok := jwtToken.Claims.(jwt.MapClaims)["type"]
	if !ok {
		render.JSON(w, 401, "Token is not valid. Type is undifined")
		return
	}
	if tokenType != "neworganisation" {
		render.JSON(w, 401, "Token is not an init organisation one")
		return
	}
	// Token passed. Initialising organisation
	store := datastores.Store()
	db := dbStore.db
	organisation := models.Organisation{
		OrganisationName: jwtToken.Claims.(jwt.MapClaims)["organisation_name"].(string),
		DockerStack:      jwtToken.Claims.(jwt.MapClaims)["organisation_stack"].(int),
		Domain:           jwtToken.Claims.(jwt.MapClaims)["organisation_domain"].(string),
		Public:           jwtToken.Claims.(jwt.MapClaims)["public"].(bool),
	}
	user := models.User{
		Username: jwtToken.Claims.(jwt.MapClaims)["owner"].(string),
		Email:    jwtToken.Claims.(jwt.MapClaims)["owner_mail"].(string),
		Password: jwtToken.Claims.(jwt.MapClaims)["owner_password"].(string),
		// Owner role should always have ID 1 as it is the first one created into the DB.
		IDRole: 1,
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	appErr := store.Organisation().Save(&organisation, db)
	if appErr != nil {
		render.JSON(w, appErr.StatusCode, appErr)
		return
	}
	appErr = store.User().Save(&user, db)
	if appErr != nil {
		render.JSON(w, appErr.StatusCode, appErr)
		return
	}
	res := initOk{
		Organisation: organisation,
		Owner:        user,
	}
	render.JSON(w, 201, res)
}

func newPublicUser(w http.ResponseWriter, r *http.Request) {
	var User models.User
	store := datastores.Store()
	db := dbStore.db
	err := chiRender.Bind(r, &User)
	organisation := store.Organisation().Get(db)
	allowedWebMails := store.AllowedWebMails().GetAll(db)
	isAuthorizedMail := false
	for _, authorizedMail := range allowedWebMails {
		filter := "*" + authorizedMail.Domain
		ok, _ := regexp.MatchString(filter, User.Email)
		isAuthorizedMail = isAuthorizedMail || ok
	}
	if !isAuthorizedMail && !organisation.Public {
		render.JSON(w, 401, "You can't sign up if organisation is not public or your email domain was unauthorized.")
	}
	if err != nil || User == (models.EmptyUser) {
		log.Print("422 here. New Public User")
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.User().Save(&User, db)
			if err == nil {
				render.JSON(w, 201, User)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

// StartAPI initialise the api with provided host and port.
func StartAPI(hostname string, port string, DbConnectionInfo *configs.DbConnection) {
	router := newRouter()
	_, _, secret = configs.InitConfig()
	// Init DB connection
	user := DbConnectionInfo.User
	db := DbConnectionInfo.Database
	pass := DbConnectionInfo.Password
	host := DbConnectionInfo.Host
	dbport := DbConnectionInfo.Port
	dbStore.db = datastores.Store().InitConnection(user, db, pass, host, dbport)
	initAuth()
	initMiddleware(router)
	initVersionRouting(router)
	// Passing -routes to the program will generate docs for the above
	// router definition. See the `routes.json` file in this folder for
	// the output.
	// log.Println(docgen.JSONRoutesDoc(router))
	// log.Println(docgen.BuildDoc(router))
	// log.Println(docgen.MarkdownRoutesDoc(router, docgen.MarkdownOpts{
	// 	ProjectPath: "github.com/titouanfreville/popcubeapi",
	// 	Intro:       "Welcomme to popcube user api.",
	// }))

	http.ListenAndServe(hostname+":"+port, router)
}
