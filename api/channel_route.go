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

func initChannelRoute(router chi.Router) {
	router.Route("/channel", func(r chi.Router) {
		r.Route("/get", func(r chi.Router) {
			r.Get("/", getAllChannel)
			r.Get("/all", getAllChannel)
			r.Get("/public", getPublicChannel)
			r.Get("/private", getPrivateChannel)
			r.Route("/fromtype/", func(r chi.Router) {
				r.Route("/:channelType", func(r chi.Router) {
					r.Use(channelContext)
					r.Get("/", getChannelFromType)
				})
			})
			r.Route("/fromname/", func(r chi.Router) {
				r.Route("/:channelName", func(r chi.Router) {
					r.Use(channelContext)
					r.Get("/", getChannelFromName)
				})
			})
		})
		r.Post("/new", newChannel)
		r.Route("/:channelID", func(r chi.Router) {
			r.Use(channelContext)
			r.Put("/update", updateChannel)
			r.Delete("/delete", deleteChannel)
		})
	})
}

func channelContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		channelID, err := strconv.ParseUint(chi.URLParam(r, "channelID"), 10, 64)
		name := chi.URLParam(r, "channelName")
		channelType := chi.URLParam(r, "channelType")
		shortcut := chi.URLParam(r, "channelShortcut")
		channel := models.Channel{}
		ctx := context.WithValue(r.Context(), "channelName", name)
		ctx = context.WithValue(r.Context(), "channelType", channelType)
		ctx = context.WithValue(r.Context(), "channelShortcut", shortcut)
		if err == nil {
			channel = datastores.NewStore().Channel().GetByID(channelID, dbStore.db)
		}
		ctx = context.WithValue(r.Context(), "channel", channel)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllChannel(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Channel().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getPublicChannel(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Channel().GetPublic(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getPrivateChannel(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Channel().GetPrivate(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getChannelFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value("channelName").(string)
	channel := store.Channel().GetByName(name, db)
	render.JSON(w, 200, channel)
}

func getChannelFromShortcut(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	channelType := r.Context().Value("channelShortcut").(string)
	channel := store.Channel().GetByShortcut(channelType, db)
	render.JSON(w, 200, channel)
}

func getChannelFromType(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	channelType := r.Context().Value("channelType").(string)
	channel := store.Channel().GetByType(channelType, db)
	render.JSON(w, 200, channel)
}

func newChannel(w http.ResponseWriter, r *http.Request) {
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
			err := store.Channel().Save(data.Channel, db)
			if err == nil {
				render.JSON(w, 200, data.Channel)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateChannel(w http.ResponseWriter, r *http.Request) {
	channel := r.Context().Value("channel").(models.Channel)
	data := struct {
		newChannel *models.Channel
		OmitID     interface{} `json:"id,omitempty"`
	}{newChannel: &channel}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	log.Printf("Channel to Update : Id : %d // Name : %s // Type : %s \n", channel.IDChannel, channel.Name, channel.Type)
	err := chiRender.Bind(request, &data)
	log.Printf("New Channel : Id : %d // Name : %s // Type : %s \n", data.newChannel.IDChannel, data.newChannel.Name, data.newChannel.Type)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Channel().Update(&channel, data.newChannel, db)
			if err == nil {
				render.JSON(w, 200, channel)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func deleteChannel(w http.ResponseWriter, r *http.Request) {
	channel := r.Context().Value("channel").(models.Channel)
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Channel().Delete(&channel, db)
		if err == nil {
			render.JSON(w, 200, "Channel correctly removed.")
		} else {
			render.JSON(w, err.StatusCode, err)
		}
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}
