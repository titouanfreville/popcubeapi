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

func initUserRoute(router chi.Router) {
	router.Route("/user", func(r chi.Router) {
		r.Get("/", getAllUser)
		r.Post("/", newUser)
		r.Get("/all", getAllUser)
		r.Post("/new", newUser)
		r.Get("/deleted", getDeletedUser)
		r.Post("/role", getUserFromRole)
		r.Route("/date/", func(r chi.Router) {
			r.Route("/:date", func(r chi.Router) {
				r.Use(userContext)
				r.Get("/", getUserFromDate)
			})
		})
		r.Route("/email/", func(r chi.Router) {
			r.Route("/:userEmail", func(r chi.Router) {
				r.Use(userContext)
				r.Get("/", getUserFromEmail)
			})
		})
		r.Route("/username/", func(r chi.Router) {
			r.Route("/:userName", func(r chi.Router) {
				r.Use(userContext)
				r.Get("/", getUserFromName)
			})
		})
		r.Route("/nickname/", func(r chi.Router) {
			r.Route("/:nickName", func(r chi.Router) {
				r.Use(userContext)
				r.Get("/", getUserFromNickName)
			})
		})
		r.Route("/firstname/", func(r chi.Router) {
			r.Route("/:firstName", func(r chi.Router) {
				r.Use(userContext)
				r.Get("/", getUserFromFirstName)
			})
		})
		r.Route("/lastname/", func(r chi.Router) {
			r.Route("/:lastName", func(r chi.Router) {
				r.Use(userContext)
				r.Get("/", getUserFromLastName)
			})
		})
		r.Route("/:userID", func(r chi.Router) {
			r.Use(userContext)
			r.Put("/update", updateUser)
			r.Delete("/delete", deleteUser)
		})
	})
}

func userContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseUint(chi.URLParam(r, "userID"), 10, 64)
		name := chi.URLParam(r, "userName")
		nickName := chi.URLParam(r, "nickName")
		firstName := chi.URLParam(r, "firstName")
		lastName := chi.URLParam(r, "lastName")
		email := chi.URLParam(r, "email")
		date, _ := strconv.ParseInt(chi.URLParam(r, "date"), 10, 64)
		oldUser := models.User{}
		ctx := context.WithValue(r.Context(), "userName", name)
		ctx = context.WithValue(r.Context(), "nickName", nickName)
		ctx = context.WithValue(ctx, "firstName", firstName)
		ctx = context.WithValue(ctx, "lastName", lastName)
		ctx = context.WithValue(ctx, "email", email)
		ctx = context.WithValue(ctx, "date", date)
		if err == nil {
			oldUser = datastores.Store().User().GetByID(userID, dbStore.db)
		}
		ctx = context.WithValue(ctx, "oldUser", oldUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.User().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getDeletedUser(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.User().GetDeleted(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getUserFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value("userName").(string)
	user := store.User().GetByUserName(name, db)
	render.JSON(w, 200, user)
}

func getUserFromNickName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value("nickName").(string)
	user := store.User().GetByNickName(name, db)
	render.JSON(w, 200, user)
}

func getUserFromFirstName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value("firstName").(string)
	user := store.User().GetByFirstName(name, db)
	render.JSON(w, 200, user)
}

func getUserFromLastName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value("lastName").(string)
	user := store.User().GetByLastName(name, db)
	render.JSON(w, 200, user)
}

func getUserFromEmail(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	email := r.Context().Value("email").(string)
	user := store.User().GetByEmail(email, db)
	render.JSON(w, 200, user)
}

func getUserFromDate(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	date := r.Context().Value("date").(int)
	user := store.User().GetOrderedByDate(date, db)
	render.JSON(w, 200, user)
}

func getUserFromRole(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Role   *models.Role
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			role := store.User().GetByRole(data.Role, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func newUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		User   *models.User
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.User().Save(data.User, db)
			if err == nil {
				render.JSON(w, 200, data.User)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		User   *models.User
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	user := r.Context().Value("oldUser").(models.User)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.User().Update(&user, data.User, db)
			if err == nil {
				render.JSON(w, 200, user)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	store := datastores.Store()
	render := renderPackage.New()
	message := deleteMessage{
		Object: user,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.User().Delete(&user, db)
		if err == nil {
			message.Success = true
			message.Message = "User well removed."
			render.JSON(w, 200, message)
		} else {
			message.Success = false
			message.Message = err.Message
			render.JSON(w, err.StatusCode, message.Message)
		}
	} else {
		render.JSON(w, 503, error503)
	}
}
