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
		r.Route("/:memberID", func(r chi.Router) {
			r.Use(memberContext)
			// swagger:route PUT /member/{memberID} Members updateMember
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
			// swagger:route PUT /member/{memberID} Members deleteMember
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
			r.Delete("/delete", deleteMember)
		})
	})
}

func memberContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		memberID, err := strconv.ParseUint(chi.URLParam(r, "memberID"), 10, 64)
		oldMember := models.Member{}
		if err == nil {
			oldMember = datastores.Store().Member().GetByID(memberID, dbStore.db)
		}
		ctx := context.WithValue(r.Context(), oldMemberKey, oldMember)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllMember(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()

	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Member().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, error503.StatusCode, error503)
	}
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
	} else {
		if err := db.DB().Ping(); err == nil {
			role := store.Member().GetByUser(data.User, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
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
	} else {
		if err := db.DB().Ping(); err == nil {
			role := store.Member().GetByChannel(data.Channel, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
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
	} else {
		if err := db.DB().Ping(); err == nil {
			role := store.Member().GetByRole(data.Role, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
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
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Member().Save(data.Member, db)
			if err == nil {
				render.JSON(w, 201, data.Member)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
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
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Member().Update(&member, data.Member, db)
			if err == nil {
				render.JSON(w, 200, member)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	member := r.Context().Value(oldMemberKey).(models.Member)
	store := datastores.Store()

	message := deleteMessageModel{
		Object: member,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Member().Delete(&member, db)
		if err == nil {
			message.Success = true
			message.Message = "Member well removed."
			render.JSON(w, 200, message)
		} else {
			message.Success = false
			message.Message = err.Message
			render.JSON(w, err.StatusCode, message.Message)
		}
	} else {
		render.JSON(w, error503.StatusCode, error503)
	}
}
