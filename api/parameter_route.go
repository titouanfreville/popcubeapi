package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/pressly/chi"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
	renderPackage "github.com/unrolled/render"
)

const (
	oldParameterKey key = "oldParameter"
)

func initParameterRoute(router chi.Router) {
	router.Route("/parameter", func(r chi.Router) {
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
		parameter := models.Parameter{}
		if err == nil {
			parameter = datastores.Store().Parameter().Get(dbStore.db)
		}
		ctx := context.WithValue(r.Context(), oldParameterKey, parameter)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllParameter(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Parameter().Get(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, error503.StatusCode, error503)
	}
}

func newParameter(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Parameter *models.Parameter
		OmitID    interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Parameter == nil {
		render.JSON(w, 500, err)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Parameter().Save(data.Parameter, db)
			if err == nil {
				render.JSON(w, 200, data.Parameter)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func updateParameter(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Parameter *models.Parameter
		OmitID    interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	parameter := r.Context().Value(oldParameterKey).(models.Parameter)
	if err != nil || data.Parameter == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Parameter().Update(&parameter, data.Parameter, db)
			if err == nil {
				render.JSON(w, 200, parameter)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}
