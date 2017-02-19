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

func initAvatarRoute(router chi.Router) {
	router.Route("/avatar", func(r chi.Router) {
		r.Get("/", getAllAvatar)
		r.Post("/", newAvatar)
		r.Get("/all", getAllAvatar)
		r.Post("/new", newAvatar)
		r.Route("/fromlink/", func(r chi.Router) {
			r.Route("/:avatarLink", func(r chi.Router) {
				r.Use(avatarContext)
				r.Get("/", getAvatarFromLink)
			})
		})
		r.Route("/fromname/", func(r chi.Router) {
			r.Route("/:avatarName", func(r chi.Router) {
				r.Use(avatarContext)
				r.Get("/", getAvatarFromName)
			})
		})
		r.Route("/:avatarID", func(r chi.Router) {
			r.Use(avatarContext)
			r.Put("/update", updateAvatar)
			r.Delete("/delete", deleteAvatar)
		})
	})
}

func avatarContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		avatarID, err := strconv.ParseUint(chi.URLParam(r, "avatarID"), 10, 64)
		name := chi.URLParam(r, "avatarName")
		link := chi.URLParam(r, "avatarLink")
		oldAvatar := models.Avatar{}
		ctx := context.WithValue(r.Context(), "avatarName", name)
		ctx = context.WithValue(ctx, "avatarLink", link)
		if err == nil {
			oldAvatar = datastores.NewStore().Avatar().GetByID(avatarID, dbStore.db)
		}
		ctx = context.WithValue(ctx, "oldAvatar", oldAvatar)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllAvatar(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Avatar().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 503, error503)
	}
}

func getAvatarFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		name := r.Context().Value("avatarName").(string)
		avatar := store.Avatar().GetByName(name, db)
		render.JSON(w, 200, avatar)
	} else {
		render.JSON(w, 503, error503)
	}
}

func getAvatarFromLink(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		link := r.Context().Value("avatarLink").(string)
		avatar := store.Avatar().GetByLink(link, db)
		render.JSON(w, 200, avatar)
	} else {
		render.JSON(w, 503, error503)
	}
}

func newAvatar(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Avatar *models.Avatar
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil {
		render.JSON(w, 422, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Avatar().Save(data.Avatar, db)
			if err == nil {
				render.JSON(w, 200, data.Avatar)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 503, error503)
		}
	}
}

func updateAvatar(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Avatar *models.Avatar
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	avatar := r.Context().Value("oldAvatar").(models.Avatar)
	if err != nil {
		render.JSON(w, 422, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Avatar().Update(&avatar, data.Avatar, db)
			if err == nil {
				render.JSON(w, 200, avatar)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 503, error503)
		}
	}
}

func deleteAvatar(w http.ResponseWriter, r *http.Request) {
	avatar := r.Context().Value("avatar").(models.Avatar)
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Avatar().Delete(&avatar, db)
		if err == nil {
			render.JSON(w, 200, "Avatar correctly removed.")
		} else {
			render.JSON(w, err.StatusCode, err)
		}
	} else {
		render.JSON(w, 503, error503)
	}
}
