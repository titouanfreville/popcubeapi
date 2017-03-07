package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/pressly/chi"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
)

const (
	oldOrganisationKey key = "oldOrganisation"
)

func initOrganisationRoute(router chi.Router) {
	router.Route("/organisation", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(Authenticator)
		// swagger:route GET /organisation Organisations getAllOrganisation
		//
		// Get organisations
		//
		// This will get all the organisations available in the organisation.
		//
		// 	Responses:
		//    200: organisationObjectSuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllOrganisation)
		// swagger:route POST /organisation Organisations newOrganisation
		//
		// New organisation
		//
		// This will create an organisation for organisation organisations library.
		//
		// 	Responses:
		//    201: organisationObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newOrganisation)
		// swagger:route GET /organisation/all Organisations getAllOrganisation1
		//
		// Get organisations
		//
		// This will get all the organisations available in the organisation.
		//
		// 	Responses:
		//    200: organisationObjectSuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllOrganisation)
		// swagger:route POST /organisation/new Organisations newOrganisation1
		//
		// New organisation
		//
		// This will create an organisation for organisation organisations library.
		//
		// 	Responses:
		//    201: organisationObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newOrganisation)
		r.Route("/:organisationID", func(r chi.Router) {
			r.Use(organisationContext)
			// swagger:route PUT /organisation/{organisationID} Organisations updateOrganisation
			//
			// Get organisation from link
			//
			// This will return the organisation object corresponding to provided link
			//
			// 	Responses:
			//    200: organisationObjectSuccess
			// 	  503: databaseError
			// 	  default: genericError
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
		ctx := context.WithValue(r.Context(), oldOrganisationKey, oldOrganisation)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllOrganisation(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()

	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Organisation().Get(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, error503.StatusCode, error503)
	}
}

func newOrganisation(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Organisation *models.Organisation
		OmitID       interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()

	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Organisation == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Organisation().Save(data.Organisation, db)
			if err == nil {
				render.JSON(w, 201, data.Organisation)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func updateOrganisation(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Organisation *models.Organisation
		OmitID       interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()

	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	organisation := r.Context().Value(oldOrganisationKey).(models.Organisation)
	if err != nil || data.Organisation == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Organisation().Update(&organisation, data.Organisation, db)
			if err == nil {
				render.JSON(w, 200, organisation)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}
