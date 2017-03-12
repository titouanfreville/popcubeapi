package api

import (
	"context"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pressly/chi"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
)

const (
	userNameKey  key = "userName"
	nickNameKey  key = "nickName"
	firstNameKey key = "firstName"
	lastNameKey  key = "lastName"
	userDateKey  key = "userDate"
	userEmailKey key = "userEmail"
	oldUserKey   key = "oldUser"
)

func initUserRoute(router chi.Router) {
	router.Route("/user", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(Authenticator)
		// swagger:route GET /user Users getAllUser
		//
		// Get users
		//
		// This will get all the users available in the organisation.
		//
		// 	Responses:
		//    200: userArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllUser)
		// swagger:route POST /user Users newUser
		//
		// New user
		//
		// This will create an user for organisation users library.
		//
		// 	Responses:
		//    201: userObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newUser)
		// swagger:route GET /user/all Users getAllUser1
		//
		// Get users
		//
		// This will get all the users available in the organisation.
		//
		// 	Responses:
		//    200: userArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllUser)
		// swagger:route POST /user/new Users newUser1
		//
		// New user
		//
		// This will create an user for organisation users library.
		//
		// 	Responses:
		//    201: userObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newUser)
		// swagger:route POST /user/invite Users inviteUser
		//
		// Invite user
		//
		// This will create an invitation token for a user.
		//
		// 	Responses:
		//    201: userObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/invite", inviteUser)
		// swagger:route GET /user/all Users getDeletedUser
		//
		// Get deleted user
		//
		// This will get all the deleted users still present in database.
		//
		// 	Responses:
		//    200: userArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/deleted", getDeletedUser)
		// swagger:route POST /user/role Users getUserFromRole
		//
		// Get users from its role
		//
		// This will return the users having provided role.
		//
		// 	Responses:
		//    200: userArraySuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/role", getUserFromRole)
		// swagger:route GET /user/date Users getOrderedByDate
		//
		// Get user ordered by date
		//
		// This will get all the users ordered by date.
		//
		// 	Responses:
		//    200: userArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/date", getOrderedByDate)
		r.Route("/email/", func(r chi.Router) {
			r.Route("/:userEmail", func(r chi.Router) {
				r.Use(userContext)
				r.Get("/", getUserFromEmail)
			})
		})
		r.Route("/username/", func(r chi.Router) {
			r.Route("/:userName", func(r chi.Router) {
				r.Use(userContext)
				// swagger:route GET /user/username/{userName} Users getUserFromName
				//
				// Get user from username
				//
				// This will return the user object corresponding to provided username
				//
				// 	Responses:
				//    200: userObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getUserFromName)
			})
		})
		r.Route("/nickname/", func(r chi.Router) {
			r.Route("/:nickName", func(r chi.Router) {
				r.Use(userContext)
				// swagger:route GET /user/nickname/{nickName} Users getUserFromNickName
				//
				// Get user from nickname
				//
				// This will return the user object corresponding to provided nickname
				//
				// 	Responses:
				//    200: userObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getUserFromNickName)
			})
		})
		r.Route("/firstname/", func(r chi.Router) {
			r.Route("/:firstName", func(r chi.Router) {
				r.Use(userContext)
				// swagger:route GET /user/firstname/{firstName} Users getUserFromFirstName
				//
				// Get user from firstname
				//
				// This will return the user object corresponding to provided firstname
				//
				// 	Responses:
				//    200: userObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getUserFromFirstName)
			})
		})
		r.Route("/lastname/", func(r chi.Router) {
			r.Route("/:lastName", func(r chi.Router) {
				r.Use(userContext)
				// swagger:route GET /user/lastname/{lastName} Users getUserFromLastName
				//
				// Get user from lastname
				//
				// This will return the user object corresponding to provided lastname
				//
				// 	Responses:
				//    200: userObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getUserFromLastName)
			})
		})
		r.Route("/:userID", func(r chi.Router) {
			r.Use(userContext)
			// swagger:route PUT /user/{userID} Users updateUser
			//
			// Update user
			//
			// This will return the new user object
			//
			// 	Responses:
			//    200: userObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Put("/update", updateUser)
			// swagger:route PUT /user/{userID} Users deleteUser
			//
			// Delete user
			//
			// This will return a delete specific mesage
			//
			// 	Responses:
			//    200: deleteMessage
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: deleteMessage
			r.Delete("/delete", deleteUser)
		})
	})
}

func userContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseUint(chi.URLParam(r, "userID"), 10, 64)
		name := chi.URLParam(r, "userName")
		nickName := chi.URLParam(r, "nickName")
		firstName := chi.URLParam(r, "firstName")
		lastName := chi.URLParam(r, "lastName")
		email := chi.URLParam(r, "email")
		date, _ := strconv.ParseInt(chi.URLParam(r, "date"), 10, 64)
		oldUser := models.User{}
		ctx := context.WithValue(r.Context(), userNameKey, name)
		ctx = context.WithValue(ctx, nickNameKey, nickName)
		ctx = context.WithValue(ctx, firstNameKey, firstName)
		ctx = context.WithValue(ctx, lastNameKey, lastName)
		ctx = context.WithValue(ctx, userEmailKey, email)
		ctx = context.WithValue(ctx, userDateKey, date)
		if err == nil {
			oldUser = datastores.Store().User().GetByID(userID, dbStore.db)
		}
		ctx = context.WithValue(ctx, oldUserKey, oldUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func canManageUser(place string, self bool, currentUser string, token *jwt.Token) bool {
	store := datastores.Store()
	db := dbStore.db
	userName := token.Claims.(jwt.MapClaims)["name"].(string)
	user := store.User().GetByUserName(userName, db)
	userRights := store.Role().GetByID(user.IDRole, db)
	if self && currentUser == userName {
		return true
	}
	if place == "organisation" || place == "global" {
		return userRights.CanManageUser
	}
	chanel := store.Channel().GetByName(place, db)
	member := store.Member().GetChannelMember(&user, &chanel, db)
	channelRights := store.Role().GetByID(member.IDRole, db)
	return channelRights.CanManageUser
}

func canInviteUser(place string, token *jwt.Token) bool {
	store := datastores.Store()
	db := dbStore.db
	userName := token.Claims.(jwt.MapClaims)["name"].(string)
	user := store.User().GetByUserName(userName, db)
	userRights := store.Role().GetByID(user.IDRole, db)
	if place == "organisation" || place == "global" {
		return userRights.CanInvite
	}
	chanel := store.Channel().GetByName(place, db)
	member := store.Member().GetChannelMember(&user, &chanel, db)
	channelRights := store.Role().GetByID(member.IDRole, db)
	return channelRights.CanInvite
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	result := store.User().GetAll(db)
	render.JSON(w, 200, result)

}

func getDeletedUser(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	result := store.User().GetDeleted(db)
	render.JSON(w, 200, result)

}

func getUserFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	name := r.Context().Value("userName").(string)
	user := store.User().GetByUserName(name, db)
	render.JSON(w, 200, user)
}

func getUserFromNickName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	name := r.Context().Value(nickNameKey).(string)
	user := store.User().GetByNickName(name, db)
	render.JSON(w, 200, user)
}

func getUserFromFirstName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	name := r.Context().Value(firstNameKey).(string)
	user := store.User().GetByFirstName(name, db)
	render.JSON(w, 200, user)
}

func getUserFromLastName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	name := r.Context().Value(lastNameKey).(string)
	user := store.User().GetByLastName(name, db)
	render.JSON(w, 200, user)
}

func getUserFromEmail(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	email := r.Context().Value(userEmailKey).(string)
	user := store.User().GetByEmail(email, db)
	render.JSON(w, 200, user)
}

func getOrderedByDate(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	date := r.Context().Value(userDateKey).(int)
	user := store.User().GetOrderedByDate(date, db)
	render.JSON(w, 200, user)
}

func getUserFromRole(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Role   *models.Role
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Role == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	role := store.User().GetByRole(data.Role, db)
	render.JSON(w, 200, role)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		User   *models.User
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageUser("global", false, "", token) {
		res := error401
		res.Message = "You don't have the right to manage user."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.User == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.User().Save(data.User, db)
	if err == nil {
		render.JSON(w, 201, data.User)
		return
	}
	render.JSON(w, apperr.StatusCode, apperr)

}

func inviteUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email   string      `json:"email"`
		Message string      `json:"message"`
		OmitID  interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	db := dbStore.db
	organisation := store.Organisation().Get(db)
	response := inviteOk{}
	request := r.Body
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageUser("global", false, "", token) || !canInviteUser("global", token) {
		res := error401
		res.Message = "You don't have the right to manage user."
		render.JSON(w, error401.StatusCode, error401)
		return
	}

	err := chiRender.Bind(request, &data)
	if err != nil || data.Email == "" {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err == nil {
		var terr error
		response.Email = data.Email
		response.Organisation = organisation.OrganisationName
		response.Token, terr = createInviteToken(data.Email, organisation.OrganisationName)
		if terr == nil {
			render.JSON(w, 201, response)
			return
		}
		render.JSON(w, 422, "Could not generate token")
		return
	}
	render.JSON(w, error503.StatusCode, error503)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		User   *models.User
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()

	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	user := r.Context().Value(oldUserKey).(models.User)
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageUser("global", true, user.Username, token) {
		res := error401
		res.Message = "You don't have the right to manage user."
		render.JSON(w, error401.StatusCode, error401)
		return
	}

	if err != nil || data.User == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err == nil {
		err := store.User().Update(&user, data.User, db)
		if err == nil {
			render.JSON(w, 200, user)
			return
		}

	}
	render.JSON(w, error503.StatusCode, error503)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(oldUserKey).(models.User)
	store := datastores.Store()
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageUser("global", true, user.Username, token) {
		res := error401
		res.Message = "You don't have the right to manage user."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	message := deleteMessageModel{
		Object: user,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.User().Delete(&user, db)
		if err == nil {
			message.Success = true
			message.Message = "User well removed."
			render.JSON(w, 200, message)
			return
		}
		message.Success = false
		message.Message = err.Message
		render.JSON(w, err.StatusCode, message.Message)
		return
	}
	render.JSON(w, 503, error503)
}
