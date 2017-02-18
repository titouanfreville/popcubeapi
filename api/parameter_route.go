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

func initParameterRoute(router chi.Router) {
	router.Route("/parameter", func(r chi.Router) {
		r.Route("/get", func(r chi.Router) {
			r.Get("/", getAllParameter)
			r.Get("/all", getAllParameter)
		})
		r.Post("/new", newParameter)
		r.Route("/:parameterID", func(r chi.Router) {
			r.Use(parameterContext)
			r.Put("/update", updateParameter)
		})
	})
}

func parameterContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := strconv.ParseUint(chi.URLParam(r, "parameterID"), 10, 64)
		parameter := models.Parameter{}
		if err == nil {
			parameter = datastores.NewStore().Parameter().Get(dbStore.db)
		}
		ctx := context.WithValue(r.Context(), "parameter", parameter)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllParameter(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Parameter().Get(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func newParameter(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Parameter *models.Parameter
		OmitID    interface{} `json:"id,omitempty"`
	}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Parameter().Save(data.Parameter, db)
			if err == nil {
				render.JSON(w, 200, data.Parameter)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateParameter(w http.ResponseWriter, r *http.Request) {
	parameter := r.Context().Value("parameter").(models.Parameter)
	data := struct {
		newParameter *models.Parameter
		OmitID       interface{} `json:"id,omitempty"`
	}{newParameter: &parameter}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Parameter().Update(&parameter, data.newParameter, db)
			if err == nil {
				render.JSON(w, 200, parameter)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}
