package datastores

import (
	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// AvatarStoreImpl Used to implement AvatarStore interface
type AvatarStoreImpl struct{}

// Avatar Generate the struct for avatar store
func (s StoreImpl) Avatar() AvatarStore {
	return AvatarStoreImpl{}
}

// Save Use to save avatar in DB
func (asi AvatarStoreImpl) Save(avatar *models.Avatar, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := avatar.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("avatarStoreImpl.Save.avatar.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(avatar) {
		transaction.Rollback()
		return u.NewLocAppError("avatarStoreImpl.Save", "save.transaction.create.already_exist", nil, "Avatar Name: "+avatar.Name)
	}
	if err := transaction.Create(&avatar).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("avatarStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update avatar in DB
func (asi AvatarStoreImpl) Update(avatar *models.Avatar, newAvatar *models.Avatar, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := avatar.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("avatarStoreImpl.Update.avatarOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if appError := newAvatar.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("avatarStoreImpl.Update.avatarNew.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Model(&avatar).Updates(&newAvatar).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("avatarStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get avatar from DB
func (asi AvatarStoreImpl) GetAll(db *gorm.DB) []models.Avatar {
	avatars := []models.Avatar{}
	db.Find(&avatars)
	return avatars
}

// GetByID Used to get avatar from DB
func (asi AvatarStoreImpl) GetByID(ID uint64, db *gorm.DB) models.Avatar {
	avatar := models.Avatar{}
	db.Where("idAvatar = ?", ID).First(&avatar)
	return avatar
}

// GetByName Used to get avatar from DB
func (asi AvatarStoreImpl) GetByName(avatarName string, db *gorm.DB) models.Avatar {
	avatar := models.Avatar{}
	db.Where("name = ?", avatarName).First(&avatar)
	return avatar
}

// GetByLink Used to get avatar from DB
func (asi AvatarStoreImpl) GetByLink(avatarLink string, db *gorm.DB) models.Avatar {
	avatar := models.Avatar{}
	db.Where("link = ?", avatarLink).First(&avatar)
	return avatar
}

// Delete Used to get avatar from DB
func (asi AvatarStoreImpl) Delete(avatar *models.Avatar, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := avatar.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("avatarStoreImpl.Delete.avatar.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&avatar).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("avatarStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
