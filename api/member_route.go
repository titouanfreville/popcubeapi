package api

import (
	"context"
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pressly/chi"
	chiRender "github.com/pressly/chi/render"
	"github.com/titouanfreville/popcubeapi/datastores"
	"github.com/titouanfreville/popcubeapi/models"
	"github.com/titouanfreville/popcubeapi/utils"
)

const (
	oldMemberKey key = "oldMember"
)

func initMemberOverChannel(channelRoutes chi.Router) {
	channelRoutes.Route("/member", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(Authenticator)
		// swagger:route POST channel Members getMemberFromChannel
		//
		// Get member into channel
		//
		// This will return all users in provided channel
		//
		// 	Responses:
		//    200: memberObjectSuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getMemberFromChannel)
		// swagger:route POST channel/{channelID}/member Members newMember
		//
		// New member
		//
		// This will create an member for organisation members library.
		//
		// 	Responses:
		//    201: memberObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newMember)
	})
	// swagger:route POST channel/{channelID}/role/{roleID} Members getMemberFromRole
	//
	// Get member having channel specifics roles
	//
	// This will return all members having a Specific role for a channel
	//
	// 	Responses:
	//    200: memberObjectSuccess
	// 	  503: databaseError
	// 	  default: genericError
	channelRoutes.Route("/role", func(r chi.Router) {
		r.Route("/:roleID", func(r chi.Router) {
			r.Use(memberContext)
			r.Get("/", getMemberFromRole)
		})
	})
	channelRoutes.Route("/user", func(r chi.Router) {
		r.Route("/:userID", func(r chi.Router) {
			r.Use(tokenAuth.Verifier)
			r.Use(Authenticator)
			r.Use(memberContext)
			// swagger:route GET channel/{channelID}/user/{userID} Members getMemberFromUser
			//
			// Get channel user is member of
			//
			// This will return all channel provided user is in
			//
			// 	Responses:
			//    200: memberObjectSuccess
			// 	  503: databaseError
			// 	  default: genericError
			r.Get("/", getMemberFromUser)
			// swagger:route PUT /channel/{channelID}/user/{userID} Members updateMember
			//
			// Update member
			//
			// This will return the new member object
			//
			// 	Responses:
			//    200: memberObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Put("/", updateMember)
			// swagger:route DELETE /channel/{channelID}/user/{userID} Members deleteMember
			//
			// Delete member
			//
			// This will return the new member object
			//
			// 	Responses:
			//    200: memberObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Delete("/", deleteMember)
		})
	})
}

func initMemberOverUser(userRoutes chi.Router) {
	userRoutes.Route("/channels/:channelID", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(Authenticator)
		r.Use(memberContext)
		// swagger:route POST /channel Members getMemberFromChannel
		//
		// Get member into channel
		//
		// This will return all users in provided channel
		//
		// 	Responses:
		//    200: memberObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getMemberFromChannel)
		// swagger:route POST /user Members getMemberFromRole
		//
		// Get member having channel specifics roles
		//
		// This will return all members having a Specific role for a channel
		//
		// 	Responses:
		//    200: memberObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		userRoutes.Post("/role/:roleID", getMemberFromRole)
		// swagger:route PUT user/{userID}/channel/{channelID} Members updateMember
		//
		// Update member
		//
		// This will return the new member object
		//
		// 	Responses:
		//    200: memberObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Put("/", updateMember)
		// swagger:route DELETE user/{userID}/channel/{channelID} Members deleteMember
		//
		// Delete member
		//
		// This will return the new member object
		//
		// 	Responses:
		//    200: memberObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Delete("/", deleteMember)
	})
}

func memberContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var ok bool
		var channel models.Channel
		var user models.User
		var userFromParam models.User
		channel, ok = ctx.Value(oldChannelKey).(models.Channel)
		if !ok {
			channel = models.Channel{}
		}
		user, ok = ctx.Value(oldUserKey).(models.User)
		if !ok {
			user = models.User{}
			userID, err := strconv.ParseUint(chi.URLParam(r, "userID"), 10, 64)
			if err == nil {
				userFromParam = datastores.Store().User().GetByID(userID, dbStore.db)
			} else {
				userID := chi.URLParam(r, "userID")
				userFromParam = datastores.Store().User().GetByUserName(userID, dbStore.db)
			}
		}
		oldMember := models.Member{}
		if user != (models.User{}) {
			channelID, err := strconv.ParseUint(chi.URLParam(r, "channelID"), 10, 64)
			if err != nil {
				channeName := chi.URLParam(r, "channelID")
				channel = datastores.Store().Channel().GetByName(channeName, dbStore.db)
				channelID = channel.IDChannel
			}
			oldMember = datastores.Store().Member().GetByID(channelID, user.IDUser, dbStore.db)
		} else if channel != (models.Channel{}) {
			userID, err := strconv.ParseUint(chi.URLParam(r, "userID"), 10, 64)
			if err != nil {
				userName := chi.URLParam(r, "userID")
				user = datastores.Store().User().GetByUserName(userName, dbStore.db)
				userID = user.IDUser
			}
			oldMember = datastores.Store().Member().GetByID(channel.IDChannel, userID, dbStore.db)
		}
		ctx = context.WithValue(ctx, oldMemberKey, oldMember)
		ctx = context.WithValue(ctx, oldChannelKey, channel)
		log.Print("Context user:", userFromParam)
		ctx = context.WithValue(ctx, oldUserKey, userFromParam)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllMember(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	result := store.Member().GetAll(db)
	render.JSON(w, 200, result)
}

func getMemberFromUser(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	user := r.Context().Value(oldUserKey).(models.User)
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	member := store.Member().GetByUser(&user, db)
	render.JSON(w, 200, member)
}

func getMemberFromChannel(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	db := dbStore.db
	channel := r.Context().Value(oldChannelKey).(models.Channel)
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	member := store.Member().GetByChannel(&channel, db)
	render.JSON(w, 200, member)
}

func getMemberFromRole(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Role   *models.Role
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Role == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	member := store.Member().GetByRole(data.Role, db)
	render.JSON(w, 200, member)
}

func newMember(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Member *models.Member
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Member == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	chanel := store.Channel().GetByID(data.Member.IDChannel, db)
	if &chanel == nil {
		message := "You are trying to invite member to chanel : " + chanel.ChannelName + " but channel doesn't exist."
		apierr := utils.NewAPIError(404, "Channel don't exist", message)
		render.JSON(w, apierr.StatusCode, apierr)
		return
	}
	if !canManageUser(chanel.ChannelName, false, "", token) {
		res := error401
		res.Message = "You don't have the right to manage user from channel : " + chanel.ChannelName + "."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	apperr := store.Member().Save(data.Member, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 201, data.Member)
}

func updateMember(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Member *models.Member
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	member := r.Context().Value(oldMemberKey).(models.Member)
	if err != nil || data.Member == nil {
		render.JSON(w, error422.StatusCode, error422)
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	chanel := store.Channel().GetByID(member.IDChannel, db)
	user := store.User().GetByID(member.IDUser, db)
	rename := store.User().GetByID(data.Member.IDUser, db)
	if &chanel == nil {
		message := "You are trying to update member from chanel : " + chanel.ChannelName + " but this channel doesn't exist."
		apierr := utils.NewAPIError(404, "Channel don't exist", message)
		render.JSON(w, apierr.StatusCode, apierr)
		return
	}
	if &user == nil {
		message := "You are trying to update member : " + user.Username + "from channel :" + chanel.ChannelName + " but this user doesn't exist."
		apierr := utils.NewAPIError(404, "Channel don't exist", message)
		render.JSON(w, apierr.StatusCode, apierr)
		return
	}
	if &rename == nil {
		rename = models.User{Username: ""}
	}
	if !canManageUser(chanel.ChannelName, user.Username == rename.Username, user.Username, token) {
		res := error401
		res.Message = "You don't have the right to manage user from channel : " + chanel.ChannelName + "."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	apperr := store.Member().Update(&member, data.Member, db)
	if apperr != nil {
		render.JSON(w, apperr.StatusCode, apperr)
		return
	}
	render.JSON(w, 200, member)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	member := r.Context().Value(oldMemberKey).(models.Member)
	store := datastores.Store()
	message := deleteMessageModel{
		Object: member,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	token := r.Context().Value(jwtTokenKey).(*jwt.Token)
	chanel := store.Channel().GetByID(member.IDChannel, db)
	if &chanel == nil {
		message := "You are trying to remove member from chanel : " + chanel.ChannelName + " but channel doesn't exist."
		apierr := utils.NewAPIError(404, "Channel doesn't exist", message)
		render.JSON(w, apierr.StatusCode, apierr)
		return
	}
	if !canManageUser(chanel.ChannelName, false, "", token) {
		res := error401
		res.Message = "You don't have the right to manage user from channel : " + chanel.ChannelName + "."
		render.JSON(w, error401.StatusCode, error401)
		return
	}
	apperr := store.Member().Delete(&member, db)
	if apperr != nil {
		message.Success = false
		message.Message = apperr.Message
		render.JSON(w, apperr.StatusCode, message.Message)
		return
	}
	message.Success = true
	message.Message = "Member well removed."
	render.JSON(w, 200, message)
}
