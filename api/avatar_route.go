package api

import (
	"context"
	"log"
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
		r.Route("/get", func(r chi.Router) {
			r.Get("/", getAllAvatar)
			r.Get("/all", getAllAvatar)
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
		})
		r.Post("/new", newAvatar)
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
		avatar := models.Avatar{}
		ctx := context.WithValue(r.Context(), "avatarName", name)
		ctx = context.WithValue(r.Context(), "avatarLink", link)
		if err == nil {
			avatar = datastores.NewStore().Avatar().GetByID(avatarID, dbStore.db)
		}
		ctx = context.WithValue(r.Context(), "avatar", avatar)
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
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getAvatarFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value("avatarName").(string)
	avatar := store.Avatar().GetByName(name, db)
	render.JSON(w, 200, avatar)
}

func getAvatarFromLink(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	link := r.Context().Value("avatarLink").(string)
	avatar := store.Avatar().GetByLink(link, db)
	render.JSON(w, 200, avatar)
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
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Avatar().Save(data.Avatar, db)
			if err == nil {
				render.JSON(w, 200, data.Avatar)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateAvatar(w http.ResponseWriter, r *http.Request) {
	avatar := r.Context().Value("avatar").(models.Avatar)
	data := struct {
		newAvatar *models.Avatar
		OmitID    interface{} `json:"id,omitempty"`
	}{newAvatar: &avatar}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	log.Printf("Avatar to Update : Id : %d // Name : %s // Link : %s \n", avatar.IDAvatar, avatar.Name, avatar.Link)
	err := chiRender.Bind(request, &data)
	log.Printf("New Avatar : Id : %d // Name : %s // Link : %s \n", data.newAvatar.IDAvatar, data.newAvatar.Name, data.newAvatar.Link)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Avatar().Update(&avatar, data.newAvatar, db)
			if err == nil {
				render.JSON(w, 200, avatar)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
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
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}
