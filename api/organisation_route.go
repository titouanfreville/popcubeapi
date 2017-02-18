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
		r.Route("/get", func(r chi.Router) {
			r.Get("/", getAllOrganisation)
			r.Get("/all", getAllOrganisation)
		})
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
		organisation := models.Organisation{}
		if err == nil {
			organisation = datastores.NewStore().Organisation().Get(dbStore.db)
		}
		ctx := context.WithValue(r.Context(), "organisation", organisation)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllOrganisation(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
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
	store := datastores.NewStore()
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
	organisation := r.Context().Value("organisation").(models.Organisation)
	data := struct {
		newOrganisation *models.Organisation
		OmitID          interface{} `json:"id,omitempty"`
	}{newOrganisation: &organisation}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Organisation().Update(&organisation, data.newOrganisation, db)
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
