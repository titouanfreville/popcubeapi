package api

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pressly/chi"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
)

const (
	oldUserParameterKey key = "oldUserParameter"
)

// To be add into user routes.
func initUserParameterRoute(router chi.Router) {
	// User ID will be in the context from user_route
	router.Route("/parameters", func(r chi.Router) {
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
			r.Put("/", updateUserParameter)
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
			r.Delete("/", deleteUserParameter)
		})
	})
}

func userParameterAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		userName := ctx.Value(oldUserKey).(models.User).Username
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
		ctx := r.Context()
		parameterName := chi.URLParam(r, "parameterName")
		oldUserParameter := models.UserParameter{}
		store := datastores.Store()
		db := dbStore.db
		user := ctx.Value(oldUserKey).(models.User)
		if (user != models.User{}) {
			oldUserParameter = store.UserParameter().GetByID(user.IDUser, parameterName, db)
		}
		ctx = context.WithValue(ctx, oldUserParameterKey, oldUserParameter)
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
	user := r.Context().Value(oldUserKey).(models.User)
	result := store.UserParameter().GetByUser(&user, db)
	render.JSON(w, 200, result)
}

func newUserParameter(w http.ResponseWriter, r *http.Request) {
	var UserParameter models.UserParameter
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &UserParameter)
	if err != nil || (UserParameter == models.UserParameter{}) {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	UserParameter.IDUser = r.Context().Value(oldUserKey).(models.User).IDUser
	apperr := store.UserParameter().Save(&UserParameter, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 201, UserParameter)
}

func updateUserParameter(w http.ResponseWriter, r *http.Request) {
	var UserParameter models.UserParameter
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &UserParameter)
	userparameter := r.Context().Value(oldUserParameterKey).(models.UserParameter)
	if err != nil || (UserParameter == models.UserParameter{}) {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.UserParameter().Update(&userparameter, &UserParameter, db)
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
