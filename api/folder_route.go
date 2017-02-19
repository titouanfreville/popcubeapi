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

func initFolderRoute(router chi.Router) {
	router.Route("/folder", func(r chi.Router) {
		r.Route("/get", func(r chi.Router) {
			r.Get("/", getAllFolder)
			r.Get("/all", getAllFolder)
			r.Post("/message", getFolderFromMessage)
			r.Route("/foldername/", func(r chi.Router) {
				r.Route("/:folderName", func(r chi.Router) {
					r.Use(folderContext)
					r.Get("/", getFolderFromName)
				})
			})
			r.Route("/link/", func(r chi.Router) {
				r.Route("/:folderLink", func(r chi.Router) {
					r.Use(folderContext)
					r.Get("/", getFolderFromLink)
				})
			})
			r.Route("/type/", func(r chi.Router) {
				r.Route("/:folderType", func(r chi.Router) {
					r.Use(folderContext)
					r.Get("/", getFolderFromType)
				})
			})
		})
		r.Post("/new", newFolder)
		r.Route("/:folderID", func(r chi.Router) {
			r.Use(folderContext)
			r.Put("/update", updateFolder)
			r.Delete("/delete", deleteFolder)
		})
	})
}

func folderContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		folderID, err := strconv.ParseUint(chi.URLParam(r, "folderID"), 10, 64)
		name := chi.URLParam(r, "folderName")
		folderType := chi.URLParam(r, "folderType")
		folderLink := chi.URLParam(r, "folderLink")
		oldFolder := models.Folder{}
		ctx := context.WithValue(r.Context(), "folderName", name)
		ctx = context.WithValue(r.Context(), "folderType", folderType)
		ctx = context.WithValue(ctx, "folderLink", folderLink)
		if err == nil {
			oldFolder = datastores.NewStore().Folder().GetByID(folderID, dbStore.db)
		}
		ctx = context.WithValue(ctx, "oldFolder", oldFolder)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllFolder(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Folder().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}

func getFolderFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value("folderName").(string)
	folder := store.Folder().GetByName(name, db)
	render.JSON(w, 200, folder)
}

func getFolderFromType(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	folderType := r.Context().Value("folderType").(string)
	folder := store.Folder().GetByType(folderType, db)
	render.JSON(w, 200, folder)
}

func getFolderFromLink(w http.ResponseWriter, r *http.Request) {
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	link := r.Context().Value("folderLink").(string)
	folder := store.Folder().GetByLink(link, db)
	render.JSON(w, 200, folder)
}

func getFolderFromMessage(w http.ResponseWriter, r *http.Request) {
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
			role := store.Folder().GetByMessage(data.Message, db)
			render.JSON(w, 200, role)
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func newFolder(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Folder *models.Folder
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
			err := store.Folder().Save(data.Folder, db)
			if err == nil {
				render.JSON(w, 200, data.Folder)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func updateFolder(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Folder *models.Folder
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	folder := r.Context().Value("oldFolder").(models.Folder)
	if err != nil {
		render.JSON(w, 500, "Internal server error")
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Folder().Update(&folder, data.Folder, db)
			if err == nil {
				render.JSON(w, 200, folder)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, 500, "Connection failure : DATABASE")
		}
	}
}

func deleteFolder(w http.ResponseWriter, r *http.Request) {
	folder := r.Context().Value("folder").(models.Folder)
	store := datastores.NewStore()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Folder().Delete(&folder, db)
		if err == nil {
			render.JSON(w, 200, "Folder correctly removed.")
		} else {
			render.JSON(w, err.StatusCode, err)
		}
	} else {
		render.JSON(w, 500, "Connection failure : DATABASE")
	}
}