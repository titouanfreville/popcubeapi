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

const (
	folderNameKey key = "folderName"
	folderTypeKey key = "folderType"
	folderLinkKey key = "folderLink"
	oldFolderKey  key = "oldFolder"
)

func initFolderRoute(router chi.Router) {
	router.Route("/folder", func(r chi.Router) {
		// swagger:route GET /folder Folders getAllFolder
		//
		// Get folders
		//
		// This will get all the folders available in the organisation.
		//
		// 	Responses:
		//    200: folderArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/", getAllFolder)
		// swagger:route POST /folder Folders newFolder
		//
		// New folder
		//
		// This will create an folder for organisation folders library.
		//
		// 	Responses:
		//    201: folderObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/", newFolder)
		// swagger:route GET /folder/all Folders getAllFolder1
		//
		// Get folders
		//
		// This will get all the folders available in the organisation.
		//
		// 	Responses:
		//    200: folderArraySuccess
		// 	  503: databaseError
		// 	  default: genericError
		r.Get("/all", getAllFolder)
		// swagger:route POST /folder/new Folders newFolder1
		//
		// New folder
		//
		// This will create an folder for organisation folders library.
		//
		// 	Responses:
		//    201: folderObjectSuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/new", newFolder)
		// swagger:route POST /folder/message Folders getFolderFromMessage
		//
		// Get folders linked to message
		//
		// Return folders linked to provided message.
		//
		// 	Responses:
		//    200: folderArraySuccess
		// 	  422: wrongEntity
		// 	  503: databaseError
		// 	  default: genericError
		r.Post("/message", getFolderFromMessage)
		r.Route("/name/", func(r chi.Router) {
			r.Route("/:folderName", func(r chi.Router) {
				r.Use(folderContext)
				// swagger:route GET /folder/name/{folderName} Folders getFolderFromName
				//
				// Get folder from name
				//
				// This will return the folder object corresponding to provided name
				//
				// 	Responses:
				//    200: folderObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getFolderFromName)
			})
		})
		r.Route("/link/", func(r chi.Router) {
			r.Route("/:folderLink", func(r chi.Router) {
				r.Use(folderContext)
				// swagger:route GET /folder/link/{folderLink} Folders getFolderFromLink
				//
				// Get folder from link
				//
				// This will return the folder object corresponding to provided link
				//
				// 	Responses:
				//    200: folderObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getFolderFromLink)
			})
		})
		r.Route("/type/", func(r chi.Router) {
			r.Route("/:folderType", func(r chi.Router) {
				r.Use(folderContext)
				// swagger:route GET /folder/type/{folderType} Folders getFolderFromType
				//
				// Get folder from type
				//
				// This will return the folder object corresponding to provided type
				//
				// 	Responses:
				//    200: folderObjectSuccess
				// 	  503: databaseError
				// 	  default: genericError
				r.Get("/", getFolderFromType)
			})
		})
		r.Route("/:folderID", func(r chi.Router) {
			r.Use(folderContext)
			// swagger:route PUT /folder/{folderID} Folders updateFolder
			//
			// Update folder
			//
			// This will return the new folder object
			//
			// 	Responses:
			//    200: avatarObjectSuccess
			// 	  422: wrongEntity
			// 	  503: databaseError
			// 	  default: genericError
			r.Put("/update", updateFolder)
			// swagger:route DELETE /folder/{folderID} Folders deleteFolder
			//
			// Delete folder
			//
			// This will return an object describing the deletion
			//
			// 	Responses:
			//    200: deleteMessage
			// 	  503: databaseError
			// 	  default: genericError
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
		ctx := context.WithValue(r.Context(), folderNameKey, name)
		ctx = context.WithValue(r.Context(), folderTypeKey, folderType)
		ctx = context.WithValue(ctx, folderLinkKey, folderLink)
		if err == nil {
			oldFolder = datastores.Store().Folder().GetByID(folderID, dbStore.db)
		}
		ctx = context.WithValue(ctx, oldFolderKey, oldFolder)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllFolder(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		result := store.Folder().GetAll(db)
		render.JSON(w, 200, result)
	} else {
		render.JSON(w, error503.StatusCode, error503)
	}
}

func getFolderFromName(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	name := r.Context().Value(folderNameKey).(string)
	folder := store.Folder().GetByName(name, db)
	render.JSON(w, 200, folder)
}

func getFolderFromType(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	folderType := r.Context().Value(folderTypeKey).(string)
	folder := store.Folder().GetByType(folderType, db)
	render.JSON(w, 200, folder)
}

func getFolderFromLink(w http.ResponseWriter, r *http.Request) {
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	link := r.Context().Value(folderLinkKey).(string)
	folder := store.Folder().GetByLink(link, db)
	render.JSON(w, 200, folder)
}

func getFolderFromMessage(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Message *models.Message
		OmitID  interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			folders := store.Folder().GetByMessage(data.Message, db)
			render.JSON(w, 200, folders)
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func newFolder(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Folder *models.Folder
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	if err != nil || data.Folder == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Folder().Save(data.Folder, db)
			if err == nil {
				render.JSON(w, 201, data.Folder)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func updateFolder(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Folder *models.Folder
		OmitID interface{} `json:"id,omitempty"`
	}
	store := datastores.Store()
	render := renderPackage.New()
	db := dbStore.db
	request := r.Body
	err := chiRender.Bind(request, &data)
	folder := r.Context().Value(oldFolderKey).(models.Folder)
	if err != nil || data.Folder == nil {
		render.JSON(w, error422.StatusCode, error422)
	} else {
		if err := db.DB().Ping(); err == nil {
			err := store.Folder().Update(&folder, data.Folder, db)
			if err == nil {
				render.JSON(w, 200, folder)
			} else {
				render.JSON(w, err.StatusCode, err)
			}
		} else {
			render.JSON(w, error503.StatusCode, error503)
		}
	}
}

func deleteFolder(w http.ResponseWriter, r *http.Request) {
	folder := r.Context().Value("folder").(models.Folder)
	store := datastores.Store()
	render := renderPackage.New()
	message := deleteMessageModel{
		Object: folder,
	}
	db := dbStore.db
	if err := db.DB().Ping(); err == nil {
		err := store.Folder().Delete(&folder, db)
		if err == nil {
			message.Success = true
			message.Message = "Folder well removed."
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
