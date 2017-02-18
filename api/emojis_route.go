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

func initEmojiRoute(router chi.Router) {
	router.Route("/emoji", func(r chi.Router) {
		r.Route("/get", func(r chi.Router) {
			r.Get("/", getAllEmoji)
			r.Get("/all", getAllEmoji)
			r.Route("/fromlink/", func(r chi.Router) {
				r.Route("/:emojiLink", func(r chi.Router) {
					r.Use(emojiContext)
					r.Get("/", getEmojiFromLink)
				})
			})
			r.Route("/fromname/", func(r chi.Router) {
				r.Route("/:emojiName", func(r chi.Router) {
					r.Use(emojiContext)
					r.Get("/", getEmojiFromName)
				})
			})
			r.Route("/fromshortcut/", func(r chi.Router) {
				r.Route("/:emojiShortcut", func(r chi.Router) {
					r.Use(emojiContext)
					r.Get("/", getEmojiFromShortcut)
				})
			})
		})
		r.Post("/new", newEmoji)
		r.Route("/:emojiID", func(r chi.Router) {
			r.Use(emojiContext)
			r.Put("/update", updateEmoji)
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
		emoji := models.Emoji{}
		ctx := context.WithValue(r.Context(), "emojiName", name)
		ctx = context.WithValue(ctx, "emojiLink", link)
		ctx = context.WithValue(ctx, "emojiShortcut", shortcut)
		if err == nil {
			emoji = datastores.NewStore().Emoji().GetByID(emojiID, dbStore.db)
		}
		ctx = context.WithValue(ctx, "emoji", emoji)

		log.Printf("ctx is :>>>>>>>>>>>>>>>>>>>>>>>>>>> \n ")
		log.Print(ctx)
		log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> \n ")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllEmoji(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Emoji().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getEmojiFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value("emojiName").(string)
	emoji := store.Emoji().GetByName(name, db)
	render.JSON(w, 200, emoji)
}

func getEmojiFromShortcut(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	link := r.Context().Value("emojiShortcut").(string)
	emoji := store.Emoji().GetByShortcut(link, db)
	render.JSON(w, 200, emoji)
}

func getEmojiFromLink(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	link := r.Context().Value("emojiLink").(string)
	emoji := store.Emoji().GetByLink(link, db)
	render.JSON(w, 200, emoji)
}

func newEmoji(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Emoji  *models.Emoji
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
			err := store.Emoji().Save(data.Emoji, db)
			if err == nil {
				render.JSON(w, 200, data.Emoji)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateEmoji(w http.ResponseWriter, r *http.Request) {
	emoji := r.Context().Value("emoji").(models.Emoji)
	data := struct {
		newEmoji *models.Emoji
		OmitID   interface{} `json:"id,omitempty"`
	}{newEmoji: &emoji}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	log.Printf("Emoji to Update : Id : %d // Name : %s // Link : %s \n", emoji.IDEmoji, emoji.Name, emoji.Link)
	err := chiRender.Bind(request, &data)
	log.Printf("New Emoji : Id : %d // Name : %s // Link : %s \n", data.newEmoji.IDEmoji, data.newEmoji.Name, data.newEmoji.Link)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Emoji().Update(&emoji, data.newEmoji, db)
			if err == nil {
				render.JSON(w, 200, emoji)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func deleteEmoji(w http.ResponseWriter, r *http.Request) {
	emoji := r.Context().Value("emoji").(models.Emoji)
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Emoji().Delete(&emoji, db)
		if err == nil {
			render.JSON(w, 200, "Emoji correctly removed.")
		} else {
			render.JSON(w, err.StatusCode, err)
		}
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}
