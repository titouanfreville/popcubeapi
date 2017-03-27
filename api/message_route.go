package api

import (
	"context"
	"net/http"
	"strconv"

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
		r.Use(Authenticator)
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
			// swagger:route PUT /message/{messageID} Messages deleteMessage
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
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	result := store.Message().GetAll(db)
	render.JSON(w, 200, result)
}

func getMessageFromDate(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	date := r.Context().Value(messageDateKey).(int)
	message := store.Message().GetByDate(date, db)
	render.JSON(w, 200, message)
}

func getMessageFromUser(w http.ResponseWriter, r *http.Request) {
	var User models.User
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &User)
	if err != nil || User == (models.User{}) {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	role := store.Message().GetByCreator(&User, db)
	render.JSON(w, 200, role)
}

func getMessageFromChannel(w http.ResponseWriter, r *http.Request) {
	var Channel models.Channel
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &Channel)
	if err != nil || Channel == (models.Channel{}) {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	role := store.Message().GetByChannel(&Channel, db)
	render.JSON(w, 200, role)
}

func newMessage(w http.ResponseWriter, r *http.Request) {
	var Message models.Message
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &Message)
	if err != nil || Message == (models.Message{}) {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Message().Save(&Message, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 201, Message)
}

func updateMessage(w http.ResponseWriter, r *http.Request) {
	var Message models.Message
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &Message)
	message := r.Context().Value(oldMessageKey).(models.Message)
	if err != nil || Message == (models.Message{}) {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Message().Update(&message, &Message, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 200, message)
}

func deleteMessageFunction(w http.ResponseWriter, r *http.Request) {
	message := r.Context().Value(oldMessageKey).(models.Message)
	store := datastores.Store()
	dmessage := deleteMessageModel{
		Object: message,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	apperr := store.Message().Delete(&message, db)
	if apperr != nil {
		dmessage.Success = false
		dmessage.Message = apperr.Message
		render.JSON(w, apperr.StatusCode, dmessage.Message)
		return
	}
	dmessage.Success = true
	dmessage.Message = "Message well removed."
	render.JSON(w, 200, message)
}
