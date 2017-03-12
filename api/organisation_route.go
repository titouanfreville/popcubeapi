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

func canManageOrganisation(token *jwt.Token) bool {
	store := datastores.Store()
	db := dbStore.db
	userName := token.Claims.(jwt.MapClaims)["name"].(string)
	user := store.User().GetByUserName(userName, db)
	userRights := store.Role().GetByID(user.IDRole, db)
	return userRights.CanManage
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
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	result := store.Organisation().Get(db)
	render.JSON(w, 200, result)
}

func newOrganisation(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Organisation *models.Organisation
		OmitID       interface{} `json:"id,omitempty"`
	}
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageOrganisation(token) {
		res := error401
		res.Message = "You don't have the right to manage organisation."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Organisation == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Organisation().Save(data.Organisation, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 201, data.Organisation)
}

func updateOrganisation(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Organisation *models.Organisation
		OmitID       interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	db := dbStore.db
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageOrganisation(token) {
		res := error401
		res.Message = "You don't have the right to manage organisation."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	request := r.Body
	err := chiRender.Bind(request, &data)
	organisation := r.Context().Value(oldOrganisationKey).(models.Organisation)
	if err != nil || data.Organisation == nil {
		render.JSON(w, error422.StatusCode, error422)
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Organisation().Update(&organisation, data.Organisation, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 200, organisation)
}
