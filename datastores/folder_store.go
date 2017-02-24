package datastores

import (
	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// FolderStoreImpl Used to implement FolderStore interface
type FolderStoreImpl struct{}

// Folder Generate the struct for folder store
func (s StoreImpl) Folder() FolderStore {
	return &FolderStoreImpl{}
}

// Save Use to save folder in BB
func (fsi FolderStoreImpl) Save(folder *models.Folder, db *gorm.DB) *u.AppError {

	transaction := db.Begin()
	if appError := folder.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("folderStoreImpl.Save.folder.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(folder) {
		transaction.Rollback()
		return u.NewLocAppError("folderStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
	}
	if err := transaction.Create(&folder).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("folderStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update folder in DB
func (fsi FolderStoreImpl) Update(folder *models.Folder, newFolder *models.Folder, db *gorm.DB) *u.AppError {

	transaction := db.Begin()
	if appError := folder.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("folderStoreImpl.Update.folderOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	// if appError := newFolder.IsValid(); appError != nil {
	// 	transaction.Rollback()
	// 	return u.NewLocAppError("folderStoreImpl.Update.folderNew.PreSave", appError.ID, nil, appError.DetailedError)
	// }
	if err := transaction.Model(&folder).Updates(&newFolder).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("folderStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get folder from DB
func (fsi FolderStoreImpl) GetAll(db *gorm.DB) []models.Folder {
	folders := []models.Folder{}
	db.Find(&folders)
	return folders
}

// GetByID Used to get folder from DB
func (fsi FolderStoreImpl) GetByID(ID uint64, db *gorm.DB) models.Folder {
	folder := models.Folder{}
	db.Where("idFolder = ?", ID).First(&folder)
	return folder
}

// GetByName Used to get folder from DB by specific name
func (fsi FolderStoreImpl) GetByName(folderName string, db *gorm.DB) []models.Folder {
	folders := []models.Folder{}
	db.Where("name = ?", folderName).Find(&folders)
	return folders
}

// GetByType get all folders from specific type
func (fsi FolderStoreImpl) GetByType(messageType string, db *gorm.DB) []models.Folder {
	folders := []models.Folder{}
	db.Where("type = ?", messageType).Find(&folders)
	return folders
}

// GetByLink get all folders having link
func (fsi FolderStoreImpl) GetByLink(messageLink string, db *gorm.DB) []models.Folder {
	folders := []models.Folder{}
	db.Where("link = ?", messageLink).Find(&folders)
	return folders
}

// GetByMessage get folder from user
func (fsi FolderStoreImpl) GetByMessage(message *models.Message, db *gorm.DB) []models.Folder {
	folders := []models.Folder{}
	db.Table("folders").Select("*").Joins("natural join messages").Where("messages.idMessage = ?", message.IDMessage).Find(&folders)
	return folders
}

// GetByFolder get folder from folder
func (fsi FolderStoreImpl) GetByFolder(folder *models.Folder, db *gorm.DB) []models.Folder {
	folders := []models.Folder{}
	db.Table("folders").Select("*").Joins("natural join folders").Where("folders.idFolder = ?", folder.IDFolder).Find(&folders)
	return folders
}

// Delete Used to get folder from DB
func (fsi FolderStoreImpl) Delete(folder *models.Folder, db *gorm.DB) *u.AppError {

	transaction := db.Begin()
	if appError := folder.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("folderStoreImpl.Delete.folder.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&folder).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("folderStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
