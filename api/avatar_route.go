package api

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
	renderPackage "github.com/unrolled/render"
)

func initAvatarRoute(router chi.Router) {
	router.Route("/avatar", func(r chi.Router) {
		r.With(paginate).Get("/", getAllAvatar)
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
	// db := store.InitConnection("root", "popcube_test", "popcube_dev", "0.0.0.0", "3306")
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Avatar().GetAll(db)
		avatarList := []models.Avatar{}
		for _, avatars := range *result {
			avatarList = append(avatarList, avatars)
		}
		render.JSON(w, 200, avatarList)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}
