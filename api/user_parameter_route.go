package api

import (
	"context"
	"net/http"

	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pressly/chi"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
)

const (
	oldUserParameterKey key = "oldUserParameter"
	userParameterUser   key = "userParameterUser"
)

// To be add into user routes.
func initUserParameterRoute(router chi.Router) {
	router.Route("/user/:userName/parameters", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(userParameterAuthenticator)
		// swagger:route GET /user/{userName}/parameters UserParameter getAllUserParameter
		//
		// Get user parameters
		//
		// This will get all the user parameters available in the organisation.
		//
		// 	Responses:
		//    200: userParameterArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllUserParameter)
		// swagger:route POST /user/{userName}/parameters UserParameter newUserParameter
		//
		// New userparameter
		//
		// This will create an userparameter for organisation userparameters library.
		//
		// 	Responses:
		//    201: userparameterObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newUserParameter)
		// swagger:route GET /user/{userName}/parameters/all UserParameter getAllUserParameter1
		//
		// Get userparameters
		//
		// This will get all the userparameters available in the organisation.
		//
		// 	Responses:
		//    200: userparameterArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllUserParameter)
		// swagger:route POST /user/{userName}/parameters/new UserParameter newUserParameter1
		//
		// New userparameter
		//
		// This will create an userparameter for organisation userparameters library.
		//
		// 	Responses:
		//    201: userparameterObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newUserParameter)
		r.Route("/:parameterName", func(r chi.Router) {
			r.Use(tokenAuth.Verifier)
			r.Use(userParameterAuthenticator)
			r.Use(userparameterContext)
			// swagger:route PUT /user/{userName}/parameters/{parameterName} UserParameter updateUserParameter
			//
			// Update userparameter
			//
			// This will return the new userparameter object
			//
			// 	Responses:
			//    200: userparameterObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Put("/update", updateUserParameter)
			// swagger:route DELETE /user/{userName}/parameters/{parameterName} UserParameter deleteUserParameter
			//
			// Delete userparameter
			//
			// This will return the new userparameter object
			//
			// 	Responses:
			//    200: userparameterObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Delete("/delete", deleteUserParameter)
		})
	})
}

func userParameterAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if jwtErr, ok := ctx.Value(jwtErrorKey).(error); ok {
			if jwtErr != nil {
				render.JSON(w, 401, jwtErr)
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

		if tokenType != "userauth" {
			render.JSON(w, 401, "Token is not an user auth one")
			return
		}

		tokenUser, ok := jwtToken.Claims.(jwt.MapClaims)["name"].(string)
		tokenEmail, ok2 := jwtToken.Claims.(jwt.MapClaims)["email"].(string)
		userFromMail := "-*-"
		if !ok && !ok2 {
			render.JSON(w, 401, "Token is not valid. User is undifined")
			return
		}
		store := datastores.Store()
		db := dbStore.db
		userName := chi.URLParam(r, "userName")
		if ok2 {
			userFromMail = store.User().GetByEmail(tokenEmail, db).Username
		}
		if userName == "" || userName != tokenUser || userName != userFromMail {
			render.JSON(w, 401, "You are not correctly identified.")
			return
		}
		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func userparameterContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userName := chi.URLParam(r, "userName")
		parameterName := chi.URLParam(r, "parameterName")
		oldUserParameter := models.UserParameter{}
		store := datastores.Store()
		db := dbStore.db
		user := store.User().GetByUserName(userName, db)
		if (user == models.User{}) {
			oldUserParameter = store.UserParameter().GetByID(user.IDUser, parameterName, db)
		}
		ctx := context.WithValue(r.Context(), oldUserParameterKey, oldUserParameter)
		ctx = context.WithValue(ctx, userParameterUser, user)
		log.Print("New context : ----------------------- \n")
		log.Print(ctx)
		log.Print("\n------------------------------------- \n")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllUserParameter(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	user := r.Context().Value(userParameterUser).(models.User)
	result := store.UserParameter().GetByUser(&user, db)
	render.JSON(w, 200, result)
}

func newUserParameter(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserParameter *models.UserParameter
		OmitID        interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.UserParameter == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.UserParameter().Save(data.UserParameter, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 201, data.UserParameter)
}

func updateUserParameter(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserParameter *models.UserParameter
		OmitID        interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	userparameter := r.Context().Value(oldUserParameterKey).(models.UserParameter)
	if err != nil || data.UserParameter == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.UserParameter().Update(&userparameter, data.UserParameter, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 200, userparameter)
}

func deleteUserParameter(w http.ResponseWriter, r *http.Request) {
	userparameter := r.Context().Value(oldUserParameterKey).(models.UserParameter)
	store := datastores.Store()
	message := deleteMessageModel{
		Object: userparameter,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.UserParameter().Delete(&userparameter, db)
	if apperr != nil {
		message.Success = false
		message.Message = apperr.Message
		render.JSON(w, apperr.StatusCode, message.Message)
		return
	}
	message.Success = true
	message.Message = "UserParameter well removed."
	render.JSON(w, 200, message)
}
