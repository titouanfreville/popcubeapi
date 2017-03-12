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
	"github.com/titouanfreville/popcubeapi/utils"
)

const (
	oldMemberKey key = "oldMember"
)

func initMemberRoute(router chi.Router) {
	router.Route("/member", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(Authenticator)
		// swagger:route GET /member Members getAllMember
		//
		// Get members
		//
		// This will get all the members available in the organisation.
		//
		// 	Responses:
		//    200: memberArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllMember)
		// swagger:route POST /member Members newMember
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
		// swagger:route GET /member/all Members getAllMember1
		//
		// Get members
		//
		// This will get all the members available in the organisation.
		//
		// 	Responses:
		//    200: memberArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllMember)
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
		r.Post("/channel", getMemberFromChannel)
		// swagger:route POST /user Members getMemberFromUser
		//
		// Get channel user is member of
		//
		// This will return all channel provided user is in
		//
		// 	Responses:
		//    200: memberObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/user", getMemberFromUser)
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
		r.Post("/role", getMemberFromRole)
		// swagger:route POST /member/new Members newMember1
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
		r.Post("/new", newMember)
	})
	router.Route("/channel/{channelID}/user/{userID}", func(r chi.Router) {
		r.Use(tokenAuth.Verifier)
		r.Use(Authenticator)
		r.Use(memberContext)
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
		r.Put("/update", updateMember)
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
		r.Delete("/delete", deleteMember)
	})
}

func memberContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		channelID, err := strconv.ParseUint(chi.URLParam(r, "channelID"), 10, 64)
		userID, err := strconv.ParseUint(chi.URLParam(r, "userID"), 10, 64)
		oldMember := models.Member{}
		if err == nil {
			oldMember = datastores.Store().Member().GetByID(channelID, userID, dbStore.db)
		}
		ctx := context.WithValue(r.Context(), oldMemberKey, oldMember)
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
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	role := store.Member().GetByUser(data.User, db)
	render.JSON(w, 200, role)
}

func getMemberFromChannel(w http.ResponseWriter, r *http.Request) {
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
		return
	}
	if err := db.DB().Ping(); err != nil {
		render.JSON(w, error503.StatusCode, error503)
		return
	}
	role := store.Member().GetByChannel(data.Channel, db)
	render.JSON(w, 200, role)
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
	role := store.Member().GetByRole(data.Role, db)
	render.JSON(w, 200, role)
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
