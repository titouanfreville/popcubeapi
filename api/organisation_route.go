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

func initOrganisationRoute(router chi.Router) {
	router.Route("/organisation", func(r chi.Router) {
		r.Get("/", getAllOrganisation)
		r.Post("/", newOrganisation)
		r.Get("/all", getAllOrganisation)
		r.Post("/new", newOrganisation)
		r.Route("/:organisationID", func(r chi.Router) {
			r.Use(organisationContext)
			r.Put("/update", updateOrganisation)
		})
	})
}

func organisationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := strconv.ParseUint(chi.URLParam(r, "organisationID"), 10, 64)
		oldOrganisation := models.Organisation{}
		if err == nil {
			oldOrganisation = datastores.Store().Organisation().Get(dbStore.db)
		}
		ctx := context.WithValue(r.Context(), "oldOrganisation", oldOrganisation)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllOrganisation(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Organisation().Get(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func newOrganisation(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Organisation *models.Organisation
		OmitID       interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Organisation().Save(data.Organisation, db)
			if err == nil {
				render.JSON(w, 200, data.Organisation)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateOrganisation(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Organisation *models.Organisation
		OmitID       interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	organisation := r.Context().Value("oldOrganisation").(models.Organisation)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Organisation().Update(&organisation, data.Organisation, db)
			if err == nil {
				render.JSON(w, 200, organisation)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}
