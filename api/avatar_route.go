package api

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
)

func (ll Base) initAvatarRoute() {
	ll.router.Route("/avatar", func(r chi.Router) {
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
	ds := *BaseConst.ds
	AvatarStore := datastores.NewAvatarStore()
	avatarListPointer := AvatarStore.GetAll(ds)
	// avatarList := []*models.Avatar{}
	// for _, avatars := range avatarListPointer {
	// 	avatarList = append(avatarList, &avatars)
	// }
	// json := models.AvatarListToJSON(avatarList)
	render.JSON(w, r, avatarListPointer)
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
