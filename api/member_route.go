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

func initMemberRoute(router chi.Router) {
	router.Route("/member", func(r chi.Router) {
		r.Route("/get", func(r chi.Router) {
			r.Get("/", getAllMember)
			r.Get("/all", getAllMember)
			r.Post("/channel", getMemberFromChannel)
			r.Post("/user", getMemberFromUser)
			r.Post("/role", getMemberFromRole)
		})
		r.Post("/new", newMember)
		r.Route("/:memberID", func(r chi.Router) {
			r.Use(memberContext)
			r.Put("/update", updateMember)
			r.Delete("/delete", deleteMember)
		})
	})
}

func memberContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		memberID, err := strconv.ParseUint(chi.URLParam(r, "memberID"), 10, 64)
		oldMember := models.Member{}
		if err == nil {
			oldMember = datastores.NewStore().Member().GetByID(memberID, dbStore.db)
		}
		ctx := context.WithValue(r.Context(), "oldMember", oldMember)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllMember(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Member().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getMemberFromUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		User   *models.User
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
			role := store.Member().GetByUser(data.User, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func getMemberFromChannel(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Channel *models.Channel
		OmitID  interface{} `json:"id,omitempty"`
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
			role := store.Member().GetByChannel(data.Channel, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func getMemberFromRole(w http.ResponseWriter, r *http.Request) {
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
			role := store.Member().GetByRole(data.Role, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func newMember(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Member *models.Member
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
			err := store.Member().Save(data.Member, db)
			if err == nil {
				render.JSON(w, 200, data.Member)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateMember(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Member *models.Member
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	member := r.Context().Value("oldMember").(models.Member)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Member().Update(&member, data.Member, db)
			if err == nil {
				render.JSON(w, 200, member)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	member := r.Context().Value("member").(models.Member)
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Member().Delete(&member, db)
		if err == nil {
			render.JSON(w, 200, "Member correctly removed.")
		} else {
			render.JSON(w, err.StatusCode, err)
		}
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}
