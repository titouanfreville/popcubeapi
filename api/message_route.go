package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/goware/jwtauth"
	"github.com/pressly/chi"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
)

const (
	messageDateKey key = "messageDate"
	oldMessageKey  key = "oldMessage"
)

func initMessageRoute(router chi.Router) {
	router.Route("/message", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(jwtauth.Authenticator)
		// swagger:route GET /message Messages getAllMessage
		//
		// Get messages
		//
		// This will get all the messages available in the organisation.
		//
		// 	Responses:
		//    200: messageArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllMessage)
		// swagger:route POST /message Messages newMessage
		//
		// New message
		//
		// This will create an message for organisation messages library.
		//
		// 	Responses:
		//    201: messageObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newMessage)
		// swagger:route GET /message/all Messages getAllMessage1
		//
		// Get messages
		//
		// This will get all the messages available in the organisation.
		//
		// 	Responses:
		//    200: messageArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllMessage)
		// swagger:route POST /message/new Messages newMessage1
		//
		// New message
		//
		// This will create an message for organisation messages library.
		//
		// 	Responses:
		//    201: messageObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newMessage)
		// swagger:route POST /channel Messages getMessageFromChannel
		//
		// Get message from channel
		//
		// This will return all mesages in provided channel
		//
		// 	Responses:
		//    200: messageObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/channel", getMessageFromChannel)
		// swagger:route POST /user Messages getMessageFromUser
		//
		// Get message from user
		//
		// This will return all mesage created by provided user
		//
		// 	Responses:
		//    200: messageObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/creator", getMessageFromUser)
		r.Route("/date/", func(r chi.Router) {
			r.Route("/:messageDate", func(r chi.Router) {
				r.Use(messageContext)
				// swagger:route GET /message/date/{messageDate} Messages getMessageFromDate
				//
				// Get message from date
				//
				// This will return the message object corresponding to provided date
				//
				// 	Responses:
				//    200: messageObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getMessageFromDate)
			})
		})
		r.Route("/:messageID", func(r chi.Router) {
			r.Use(messageContext)
			// swagger:route PUT /message/{messageID} Messages updateMessage
			//
			// Update message
			//
			// This will return the new message object
			//
			// 	Responses:
			//    200: messageObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Put("/update", updateMessage)
			// swagger:route PUT /message/{messageID} Messages updateMessage
			//
			// Update message
			//
			// This will return the new message object
			//
			// 	Responses:
			//    200: messageObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Delete("/delete", deleteMessageFunction)
		})
	})
}

func messageContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		messageID, err := strconv.ParseUint(chi.URLParam(r, "messageID"), 10, 64)
		date, _ := strconv.ParseInt(chi.URLParam(r, "messageDate"), 10, 64)
		oldMessage := models.Message{}
		ctx := context.WithValue(r.Context(), messageDateKey, date)
		if err == nil {
			oldMessage = datastores.Store().Message().GetByID(messageID, dbStore.db)
		}
		ctx = context.WithValue(ctx, oldMessageKey, oldMessage)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllMessage(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()

	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Message().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, error503.StatusCode, error503)
	}
}

func getMessageFromDate(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()

	db := dbStore.db
	date := r.Context().Value(messageDateKey).(int)
	message := store.Message().GetByDate(date, db)
	render.JSON(w, 200, message)
}

func getMessageFromUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		User   *models.User
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()

	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.User == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			role := store.Message().GetByCreator(data.User, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func getMessageFromChannel(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Channel *models.Channel
		OmitID  interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()

	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Channel == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			role := store.Message().GetByChannel(data.Channel, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func newMessage(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Message *models.Message
		OmitID  interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()

	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Message == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Message().Save(data.Message, db)
			if err == nil {
				render.JSON(w, 201, data.Message)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func updateMessage(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Message *models.Message
		OmitID  interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()

	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	message := r.Context().Value(oldMessageKey).(models.Message)
	if err != nil || data.Message == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Message().Update(&message, data.Message, db)
			if err == nil {
				render.JSON(w, 200, message)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func deleteMessageFunction(w http.ResponseWriter, r *http.Request) {
	message := r.Context().Value(oldMessageKey).(models.Message)
	store := datastores.Store()

	dmessage := deleteMessageModel{
		Object: message,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Message().Delete(&message, db)
		if err == nil {
			dmessage.Success = true
			dmessage.Message = "Message well removed."
			render.JSON(w, 200, message)
		} else {
			dmessage.Success = false
			dmessage.Message = err.Message
			render.JSON(w, err.StatusCode, dmessage.Message)
		}
	} else {
		render.JSON(w, 503, error503)
	}
}
