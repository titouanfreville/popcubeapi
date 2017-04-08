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
	oldParameterKey key = "oldParameter"
)

func initParameterRoute(router chi.Router) {
	router.Route("/parameter", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(Authenticator)
		// swagger:route GET /parameter Parameters getAllParameter
		//
		// Get parameters
		//
		// This will get all the parameters available in the organisation.
		//
		// 	Responses:
		//    200: parameterObjectSuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllParameter)
		// swagger:route POST /parameter Parameters newParameter
		//
		// New parameter
		//
		// This will create an parameter for organisation parameters library.
		//
		// 	Responses:
		//    201: parameterObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newParameter)
		// swagger:route GET /parameter/all Parameters getAllParameter1
		//
		// Get parameters
		//
		// This will get all the parameters available in the organisation.
		//
		// 	Responses:
		//    200: parameterObjectSuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllParameter)
		// swagger:route POST /parameter/new Parameters newParameter1
		//
		// New parameter
		//
		// This will create an parameter for organisation parameters library.
		//
		// 	Responses:
		//    201: parameterObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newParameter)
		r.Route("/:parameterID", func(r chi.Router) {
			r.Use(parameterContext)
			// swagger:route PUT /parameter/{parameterID} Parameters updateParameter
			//
			// Update parameter
			//
			// This will return the new parameter object
			//
			// 	Responses:
			//    200: avatarObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Put("/update", updateParameter)
		})
	})
}

func parameterContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := strconv.ParseUint(chi.URLParam(r, "parameterID"), 10, 64)
		parameter := models.EmptyParameter
		if err == nil {
			parameter = datastores.Store().Parameter().Get(dbStore.db)
		}
		ctx := context.WithValue(r.Context(), oldParameterKey, parameter)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllParameter(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	result := store.Parameter().Get(db)
	render.JSON(w, 200, result)
}

func newParameter(w http.ResponseWriter, r *http.Request) {
	var Parameter models.Parameter
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageOrganisation(token) {
		res := error401
		res.Message = "You don't have the right to manage organisation."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	store := datastores.Store()
	db := dbStore.db
	
	err := chiRender.Bind(r, &Parameter)
	if err != nil || Parameter == (models.EmptyParameter) {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Parameter().Save(&Parameter, db)
	if err != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 200, Parameter)
}

func updateParameter(w http.ResponseWriter, r *http.Request) {
	var Parameter models.Parameter
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageOrganisation(token) {
		res := error401
		res.Message = "You don't have the right to manage organisation."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	store := datastores.Store()
	db := dbStore.db
	
	err := chiRender.Bind(r, &Parameter)
	parameter := r.Context().Value(oldParameterKey).(models.Parameter)
	if err != nil || Parameter == (models.EmptyParameter) {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Parameter().Update(&parameter, &Parameter, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 200, parameter)
}
