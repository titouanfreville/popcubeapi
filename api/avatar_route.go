package api

import (
	"net/http"

	"github.com/pressly/chi"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
	renderPackage "github.com/unrolled/render"
)

func initAvatarRoute(router chi.Router) {
	router.Route("/avatar", func(r chi.Router) {
		r.Get("/", getAllAvatar)
		r.Post("/new", newAvatar)
	})
}

// func AvatarContext(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		avatarID := chi.URLParam(r, "id_avatar")
// 		ctx := context.WithValue(r.Context(), "article", article)
// 		next.ServeHTTP(w, r.WithContext())
// 	})
// }

func getAllAvatar(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Avatar().GetAll(db)
		// avatarList := []models.Avatar{}
		// for _, avatars := range *result {
		// 	avatarList = append(avatarList, avatars)
		// }
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
	// decoder := json.NewDecoder(request)
	// avatar := models.Avatar{}
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

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
// func paginate(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// just a stub.. some ideas are to look at URL query params for something like
// 		// the page number, or the limit, and send a query cursor down the chain
// 		next.ServeHTTP(w, r)
// 	})
// }
