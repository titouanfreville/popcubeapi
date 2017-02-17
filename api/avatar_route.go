package api

import (
	"context"
	"net/http"

	"github.com/pressly/chi"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
	renderPackage "github.com/unrolled/render"
)

func initAvatarRoute(router chi.Router) {
	router.Route("/avatar", func(r chi.Router) {
		r.Route("/get", func(r chi.Router) {
			r.Use(avatarContext)
			r.Get("/", getAllAvatar)
			r.Get("/all", getAllAvatar)
			// r.Get("/fromlink/:link", getAvatarFromLink)
			r.Get("/fromname/:avatarname/", getAvatarFromName)
		})
		r.Post("/new", newAvatar)
		// r.Route("/:avatarID", func(r chi.Router) {
		// 	r.Get("/", getAvatarFromID)
		// 	r.Get("/get", getAvatarFromID)
		// 	r.Put("/update", updateAvatar)
		// 	r.Delete("/delete", deleteAvatar)
		// })
	})
}

func avatarContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// avatarID := chi.URLParam(r, "avatarID")
		name := chi.URLParam(r, "avatarname")
		// link := chi.URLParam(r, "link")
		// ctx := context.WithValue(r.Context(), "avatar_id", avatarID)
		ctx := context.WithValue(r.Context(), "avatar_name", name)
		// ctx = context.WithValue(ctx, "avatar_link", link)
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
	avatarName := r.Context().Value("avatar_name").(*string)
	if err := db.DB().Ping(); err == nil {
		result := store.Avatar().GetByName(*avatarName, db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
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

// func updateAvatar(w http.ResponseWriter, r *http.Request) {
// 	var data struct {
// 		newAvatar *models.Avatar
// 		OmitID    interface{} `json:"id,omitempty"`
// 	}
// 	store := datastores.NewStore()
// 	render := renderPackage.New()
// 	db := dbStore.db
// 	request := r.Body
// 	err := chiRender.Bind(request, &data)
// 	if err != nil {
// 		render.JSON(w, 500, "Internal server error")
// 	} else {
// 		if err := db.DB().Ping(); err == nil {
// 			err := store.Avatar().Update(data.newAvatar, db)
// 			if err == nil {
// 				render.JSON(w, 200, data.newAvatar)
// 			} else {
// 				render.JSON(w, err.StatusCode, err)
// 			}
// 		} else {
// 			render.JSON(w, 500, "Connection failure : DATABASE")
// 		}
// 	}
// }

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
// func paginate(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// just a stub.. some ideas are to look at URL query params for something like
// 		// the page number, or the limit, and send a query cursor down the chain
// 		next.ServeHTTP(w, r)
// 	})
// }
