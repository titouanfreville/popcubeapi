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
	channelNameKey key = "channelName"
	channelTypeKey key = "channelType"
	oldChannelKey  key = "oldChannel"
)

func initChannelRoute(router chi.Router) {
	router.Route("/channel", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(Authenticator)
		// swagger:route GET /channel Channels getAllChannel
		//
		// Get channels
		//
		// This will get all the channels available in the organisation.
		//
		// 	Responses:
		//    200: channelArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllChannel)
		// swagger:route POST /channel Channels newChannel
		//
		// New channel
		//
		// This will create an channel for organisation channels library.
		//
		// 	Responses:
		//    201: channelObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newChannel)
		// swagger:route GET /channel/all Channels getAllChannel1
		//
		// Get channels
		//
		// This will get all the channels available in the organisation.
		//
		// 	Responses:
		//    200: channelArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllChannel)
		// swagger:route POST /channel/new Channels newChannel1
		//
		// New channel
		//
		// This will create an channel for organisation channels library.
		//
		// 	Responses:
		//    201: channelObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newChannel)
		// swagger:route GET /channel/public Channels getPublicChannel
		//
		// Get public channels
		//
		// This will get all the public channels available in the organisation.
		//
		// 	Responses:
		//    200: channelArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/public", getPublicChannel)
		// swagger:route GET /channel/private Channels getPrivateChannel
		//
		// Get private channels
		//
		// This will get all the private channels available in the organisation.
		//
		// 	Responses:
		//    200: channelArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/private", getPrivateChannel)
		r.Route("/type/", func(r chi.Router) {
			r.Route("/:channelType", func(r chi.Router) {
				r.Use(channelContext)
				// swagger:route GET /channel/type/{channelType} Channels getChannelFromType
				//
				// Get channels of provided type
				//
				// This will get all the channels of provided type available in the organisation.
				//
				// 	Responses:
				//    200: channelArraySuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getChannelFromType)
			})
		})
		r.Route("/name/", func(r chi.Router) {
			r.Route("/:channelName", func(r chi.Router) {
				r.Use(channelContext)
				// swagger:route GET /channel/name/{channelName} Channels getChannelFromName
				//
				// Get nammed channel
				//
				// This will get the channels having provided name in the organisation.
				//
				// 	Responses:
				//    200: channelObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getChannelFromName)
			})
		})
		r.Route("/:channelID", func(r chi.Router) {
			r.Use(channelContext)
			// swagger:route PUT /channel/{channelID} Channels updateChannel
			//
			// Update channel
			//
			// This will return the new channel object
			//
			// 	Responses:
			//    200: channelObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Put("/update", updateChannel)
			// swagger:route DELETE /channel/{channelID} Channels deleteChannel
			//
			// Delete channel
			//
			// This will return an object describing the deletion
			//
			// 	Responses:
			//    200: deleteMessage
			// 	  503: databaseError
			// 	  default: genericError
			r.Delete("/delete", deleteChannel)
		})
	})
}

func canModerate(currentChannelID uint64, token *jwt.Token) bool {
	store := datastores.Store()
	db := dbStore.db
	userName := token.Claims.(jwt.MapClaims)["name"].(string)
	user := store.User().GetByUserName(userName, db)
	userRights := store.Role().GetByID(user.IDRole, db)
	chanel := store.Channel().GetByID(currentChannelID, db)
	member := store.Member().GetChannelMember(&user, &chanel, db)
	memberRights := store.Role().GetByID(member.IDRole, db)
	return (memberRights != models.Role{} && memberRights.CanManageUser || memberRights == models.Role{} && userRights.CanManageUser)
}

func canArchive(currentChannelID uint64, token *jwt.Token) bool {
	store := datastores.Store()
	db := dbStore.db
	userName := token.Claims.(jwt.MapClaims)["name"].(string)
	user := store.User().GetByUserName(userName, db)
	userRights := store.Role().GetByID(user.IDRole, db)
	chanel := store.Channel().GetByID(currentChannelID, db)
	member := store.Member().GetChannelMember(&user, &chanel, db)
	memberRights := store.Role().GetByID(member.IDRole, db)
	return (memberRights != models.Role{} && memberRights.CanArchive || memberRights == models.Role{} && userRights.CanArchive)
}

func channelContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		channelID, err := strconv.ParseUint(chi.URLParam(r, "channelID"), 10, 64)
		name := chi.URLParam(r, "channelName")
		channelType := chi.URLParam(r, "channelType")
		oldChannel := models.Channel{}
		ctx := context.WithValue(r.Context(), channelNameKey, name)
		ctx = context.WithValue(ctx, channelTypeKey, channelType)
		if err == nil {
			oldChannel = datastores.Store().Channel().GetByID(channelID, dbStore.db)
		}
		ctx = context.WithValue(ctx, oldChannelKey, oldChannel)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllChannel(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	result := store.Channel().GetAll(db)
	render.JSON(w, 200, result)

}

func getPublicChannel(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	result := store.Channel().GetPublic(db)
	render.JSON(w, 200, result)

}

func getPrivateChannel(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	result := store.Channel().GetPrivate(db)
	render.JSON(w, 200, result)

}

func getChannelFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	name := r.Context().Value(channelNameKey).(string)
	channel := store.Channel().GetByName(name, db)
	render.JSON(w, 200, channel)
}

func getChannelFromType(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	channelType := r.Context().Value(channelTypeKey).(string)
	channel := store.Channel().GetByType(channelType, db)
	render.JSON(w, 200, channel)
}

func newChannel(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Channel *models.Channel
		OmitID  interface{} `json:"id,omitempty"`
	}
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageUser("global", false, "", token) {
		res := error401
		res.Message = "You don't have the right to manage user."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Channel == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	rerr := store.Channel().Save(data.Channel, db)
	if err != nil {
		render.JSON(w, rerr.StatusCode, rerr)
		return
	}
	render.JSON(w, 201, data.Channel)
}

func updateChannel(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Channel *models.Channel
		OmitID  interface{} `json:"id,omitempty"`
	}
	channel := r.Context().Value(oldChannelKey).(models.Channel)
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageUser(channel.ChannelName, false, "", token) {
		res := error401
		res.Message = "You don't have the right to manage user."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Channel == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	rerr := store.Channel().Update(&channel, data.Channel, db)
	if err == nil {
		render.JSON(w, rerr.StatusCode, rerr)
		return
	}
	render.JSON(w, 200, channel)
}

func deleteChannel(w http.ResponseWriter, r *http.Request) {
	channel := r.Context().Value(oldChannelKey).(models.Channel)
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	if !canManageUser(channel.ChannelName, false, "", token) {
		res := error401
		res.Message = "You don't have the right to manage user."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	store := datastores.Store()
	message := deleteMessageModel{
		Object: channel,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	err := store.Channel().Delete(&channel, db)
	if err == nil {
		message.Success = false
		message.Message = err.Message
		render.JSON(w, err.StatusCode, message.Message)
		return
	}
	message.Success = true
	message.Message = "Channel well removed."
	render.JSON(w, 200, message)
}
