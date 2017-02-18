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

func initRoleRoute(router chi.Router) {
	router.Route("/role", func(r chi.Router) {
		r.Route("/get", func(r chi.Router) {
			r.Get("/", getAllRole)
			r.Get("/all", getAllRole)
			r.Post("/fromrights", getRoleFromRight)
			r.Route("/fromname/", func(r chi.Router) {
				r.Route("/:roleName", func(r chi.Router) {
					r.Use(roleContext)
					r.Get("/", getRoleFromName)
				})
			})
		})
		r.Post("/new", newRole)
		r.Route("/:roleID", func(r chi.Router) {
			r.Use(roleContext)
			r.Put("/update", updateRole)
			r.Delete("/delete", deleteRole)
		})
	})
}

func roleContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		roleID, err := strconv.ParseUint(chi.URLParam(r, "roleID"), 10, 64)
		name := chi.URLParam(r, "roleName")
		oldRole := models.Role{}
		ctx := context.WithValue(r.Context(), "roleName", name)
		if err == nil {
			oldRole = datastores.NewStore().Role().GetByID(roleID, dbStore.db)
		}
		ctx = context.WithValue(ctx, "oldRole", oldRole)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllRole(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Role().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getRoleFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value("roleName").(string)
	role := store.Role().GetByName(name, db)
	render.JSON(w, 200, role)
}

func getRoleFromRight(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Role   *models.Role
		OmitID interface{} `json:"id,omitempty"`
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
			role := store.Role().GetByRights(data.Role, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func newRole(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Role   *models.Role
		OmitID interface{} `json:"id,omitempty"`
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
			err := store.Role().Save(data.Role, db)
			if err == nil {
				render.JSON(w, 200, data.Role)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateRole(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Role   *models.Role
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	role := r.Context().Value("oldRole").(models.Role)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Role().Update(&role, data.Role, db)
			if err == nil {
				render.JSON(w, 200, role)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func deleteRole(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("role").(models.Role)
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Role().Delete(&role, db)
		if err == nil {
			render.JSON(w, 200, "Role correctly removed.")
		} else {
			render.JSON(w, err.StatusCode, err)
		}
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}
