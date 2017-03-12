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
	emojiNameKey     key = "emojiName"
	emojiLinkKey     key = "emojiLink"
	emojiShortcutKey key = "emojiShortcut"
	oldEmojiKey      key = "oldEmoji"
)

func initEmojiRoute(router chi.Router) {
	router.Route("/emoji", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(Authenticator)
		// swagger:route GET /emoji Emojis getAllEmoji
		//
		// Get emojis
		//
		// This will get all the emojis available in the organisation.
		//
		// 	Responses:
		//    200: emojiArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllEmoji)
		// swagger:route POST /emoji Emojis newEmoji
		//
		// New emoji
		//
		// This will create an emoji for organisation emojis library.
		//
		// 	Responses:
		//    201: emojiObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newEmoji)
		// swagger:route GET /emoji/all Emojis getAllEmoji1
		//
		// Get emojis
		//
		// This will get all the emojis available in the organisation.
		//
		// 	Responses:
		//    200: emojiArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllEmoji)
		// swagger:route POST /emoji Emojis newEmoji1
		//
		// New emoji
		//
		// This will create an emoji for organisation emojis library.
		//
		// 	Responses:
		//    201: emojiObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newEmoji)
		r.Route("/link/", func(r chi.Router) {
			r.Route("/:emojiLink", func(r chi.Router) {
				r.Use(emojiContext)
				// swagger:route GET /emoji/link/{emojiLink} Emojis getEmojiFromLink
				//
				// Get emoji from link
				//
				// This will return the emoji object corresponding to provided link
				//
				// 	Responses:
				//    200: emojiObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getEmojiFromLink)
			})
		})
		r.Route("/name/", func(r chi.Router) {
			r.Route("/:emojiName", func(r chi.Router) {
				r.Use(emojiContext)
				// swagger:route GET /emoji/name/{emojiName} Emojis getEmojiFromName
				//
				// Get emoji from name
				//
				// This will return the emoji object corresponding to provided name
				//
				// 	Responses:
				//    200: emojiObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getEmojiFromName)
			})
		})
		r.Route("/shortcut/", func(r chi.Router) {
			r.Route("/:emojiShortcut", func(r chi.Router) {
				r.Use(emojiContext)
				// swagger:route GET /emoji/shortcut/{emojiShortcut} Emojis getEmojiFromShortcut
				//
				// Get emoji from shortcut
				//
				// This will return the emoji object corresponding to provided shortcut
				//
				// 	Responses:
				//    200: emojiObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getEmojiFromShortcut)
			})
		})
		r.Route("/:emojiID", func(r chi.Router) {
			r.Use(emojiContext)
			// swagger:route PUT /emoji/{emojiID} Emojis updateEmoji
			//
			// Update emoji
			//
			// This will return the new emoji object
			//
			// 	Responses:
			//    200: avatarObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Put("/update", updateEmoji)
			// swagger:route DELETE /emoji/{emojiID} Emojis deleteEmoji
			//
			// Delete emoji
			//
			// This will return an object describing the deletion
			//
			// 	Responses:
			//    200: deleteMessage
			// 	  503: databaseError
			// 	  default: genericError
			r.Delete("/delete", deleteEmoji)
		})
	})
}

func emojiContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		emojiID, err := strconv.ParseUint(chi.URLParam(r, "emojiID"), 10, 64)
		name := chi.URLParam(r, "emojiName")
		link := chi.URLParam(r, "emojiLink")
		shortcut := chi.URLParam(r, "emojiShortcut")
		oldEmoji := models.Emoji{}
		ctx := context.WithValue(r.Context(), emojiNameKey, name)
		ctx = context.WithValue(ctx, emojiLinkKey, link)
		ctx = context.WithValue(ctx, emojiShortcutKey, shortcut)
		if err == nil {
			oldEmoji = datastores.Store().Emoji().GetByID(emojiID, dbStore.db)
		}
		ctx = context.WithValue(ctx, oldEmojiKey, oldEmoji)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllEmoji(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Emoji().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, error503.StatusCode, error503)
	}
}

func getEmojiFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	name := r.Context().Value(emojiNameKey).(string)
	emoji := store.Emoji().GetByName(name, db)
	render.JSON(w, 200, emoji)
}

func getEmojiFromShortcut(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	link := r.Context().Value(emojiShortcutKey).(string)
	emoji := store.Emoji().GetByShortcut(link, db)
	render.JSON(w, 200, emoji)
}

func getEmojiFromLink(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	link := r.Context().Value(emojiLinkKey).(string)
	emoji := store.Emoji().GetByLink(link, db)
	render.JSON(w, 200, emoji)
}

func newEmoji(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Emoji  *models.Emoji
		OmitID interface{} `json:"id,omitempty"`
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
	if err != nil || data.Emoji == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}

	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Emoji().Save(data.Emoji, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 201, data.Emoji)
}

func updateEmoji(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Emoji  *models.Emoji
		OmitID interface{} `json:"id,omitempty"`
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
	emoji := r.Context().Value(oldEmojiKey).(models.Emoji)
	if err != nil || data.Emoji == nil {
		render.JSON(w, error422.StatusCode, error422)
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Emoji().Update(&emoji, data.Emoji, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 200, emoji)
}

func deleteEmoji(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageOrganisation(token) {
		res := error401
		res.Message = "You don't have the right to manage organisation."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	emoji := r.Context().Value(oldEmojiKey).(models.Emoji)
	store := datastores.Store()
	message := deleteMessageModel{
		Object: emoji,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Emoji().Delete(&emoji, db)
	if apperr != nil {
		message.Success = false
		message.Message = apperr.Message
		render.JSON(w, apperr.StatusCode, message.Message)
		return
	}
	message.Success = true
	message.Message = "Emoji well removed."
	render.JSON(w, 200, message)
}
