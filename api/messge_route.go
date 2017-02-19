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

func initMessageRoute(router chi.Router) {
	router.Route("/message", func(r chi.Router) {
		r.Route("/get", func(r chi.Router) {
			r.Get("/", getAllMessage)
			r.Get("/all", getAllMessage)
			r.Post("/channel", getMessageFromChannel)
			r.Post("/creator", getMessageFromUser)
			r.Route("/date/", func(r chi.Router) {
				r.Route("/:date", func(r chi.Router) {
					r.Use(messageContext)
					r.Get("/", getMessageFromDate)
				})
			})
		})
		r.Post("/new", newMessage)
		r.Route("/:messageID", func(r chi.Router) {
			r.Use(messageContext)
			r.Put("/update", updateMessage)
			r.Delete("/delete", deleteMessage)
		})
	})
}

func messageContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		messageID, err := strconv.ParseUint(chi.URLParam(r, "messageID"), 10, 64)
		date, _ := strconv.ParseInt(chi.URLParam(r, "date"), 10, 64)
		oldMessage := models.Message{}
		ctx := context.WithValue(r.Context(), "date", date)
		if err == nil {
			oldMessage = datastores.NewStore().Message().GetByID(messageID, dbStore.db)
		}
		ctx = context.WithValue(ctx, "oldMessage", oldMessage)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllMessage(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Message().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getMessageFromDate(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	date := r.Context().Value("date").(int)
	message := store.Message().GetByDate(date, db)
	render.JSON(w, 200, message)
}

func getMessageFromUser(w http.ResponseWriter, r *http.Request) {
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
			role := store.Message().GetByCreator(data.User, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func getMessageFromChannel(w http.ResponseWriter, r *http.Request) {
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
			role := store.Message().GetByChannel(data.Channel, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func newMessage(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Message *models.Message
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
			err := store.Message().Save(data.Message, db)
			if err == nil {
				render.JSON(w, 200, data.Message)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateMessage(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Message *models.Message
		OmitID  interface{} `json:"id,omitempty"`
	}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	message := r.Context().Value("oldMessage").(models.Message)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Message().Update(&message, data.Message, db)
			if err == nil {
				render.JSON(w, 200, message)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	message := r.Context().Value("message").(models.Message)
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Message().Delete(&message, db)
		if err == nil {
			render.JSON(w, 200, "Message correctly removed.")
		} else {
			render.JSON(w, err.StatusCode, err)
		}
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}
