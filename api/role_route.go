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
	roleNameKey key = "roleName"
	oldRoleKey  key = "oldRole"
)

func initRoleRoute(router chi.Router) {
	router.Route("/role", func(r chi.Router) {
		// swagger:route GET /role Roles getAllRole
		//
		// Get roles
		//
		// This will get all the roles available in the organisation.
		//
		// 	Responses:
		//    200: roleArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllRole)
		// swagger:route GET /role Roles getAllRole
		//
		// Get roles
		//
		// This will get all the roles available in the organisation.
		//
		// 	Responses:
		//    201: roleArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newRole)
		// swagger:route GET /role/all Roles getAllRole1
		//
		// Get roles
		//
		// This will get all the roles available in the organisation.
		//
		// 	Responses:
		//    200: roleArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllRole)
		// swagger:route POST /role/new Roles newRole1
		//
		// New role
		//
		// This will create an role for organisation roles library.
		//
		// 	Responses:
		//    201: roleObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newRole)
		// swagger:route POST /role/rights Roles getRoleFromRights
		//
		// Get role having provided rights
		//
		// Return an array of roles corresponding to rights
		//
		// 	Responses:
		//    200: roleArraySuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/rights", getRoleFromRight)
		r.Route("/name/", func(r chi.Router) {
			r.Route("/:roleName", func(r chi.Router) {
				r.Use(roleContext)
				// swagger:route GET /role/name/{roleName} Roles getRoleFromName
				//
				// Get role from name
				//
				// This will return the role object corresponding to provided name
				//
				// 	Responses:
				//    200: roleObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getRoleFromName)
			})
		})
		r.Route("/:roleID", func(r chi.Router) {
			r.Use(roleContext)
			// swagger:route PUT /role/{roleID} Roles updateRole
			//
			// Update role
			//
			// This will return the new role object
			//
			// 	Responses:
			//    200: avatarObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Put("/update", updateRole)
			// swagger:route DELETE /role/{roleID} Roles deleteRole
			//
			// Delete role
			//
			// This will return an object describing the deletion
			//
			// 	Responses:
			//    200: deleteMessage
			// 	  503: databaseError
			// 	  default: genericError
			r.Delete("/delete", deleteRole)
		})
	})
}

func roleContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		roleID, err := strconv.ParseUint(chi.URLParam(r, "roleID"), 10, 64)
		name := chi.URLParam(r, "roleName")
		oldRole := models.Role{}
		ctx := context.WithValue(r.Context(), roleNameKey, name)
		if err == nil {
			oldRole = datastores.Store().Role().GetByID(roleID, dbStore.db)
		}
		ctx = context.WithValue(ctx, oldRoleKey, oldRole)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllRole(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Role().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, error503.StatusCode, error503)
	}
}

func getRoleFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value(roleNameKey).(string)
	role := store.Role().GetByName(name, db)
	render.JSON(w, 200, role)
}

func getRoleFromRight(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Role   *models.Role
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Role == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			role := store.Role().GetByRights(data.Role, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func newRole(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Role   *models.Role
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Role == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Role().Save(data.Role, db)
			if err == nil {
				render.JSON(w, 201, data.Role)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func updateRole(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Role   *models.Role
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	role := r.Context().Value(oldRoleKey).(models.Role)
	if err != nil || data.Role == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Role().Update(&role, data.Role, db)
			if err == nil {
				render.JSON(w, 200, role)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func deleteRole(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(oldRoleKey).(models.Role)
	store := datastores.Store()
	render := renderPackage.New()
	message := deleteMessageModel{
		Object: role,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Role().Delete(&role, db)
		if err == nil {
			message.Success = true
			message.Message = "Role well removed."
			render.JSON(w, 200, message)
		} else {
			message.Success = false
			message.Message = err.Message
			render.JSON(w, err.StatusCode, message.Message)
		}
	} else {
		render.JSON(w, 503, error503)
	}
}
