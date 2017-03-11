package datastores

import (
	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// ReadStoreImpl Used to implement ReadStore interface
type ReadStoreImpl struct{}

// Read Generate the struct for read store
func (s StoreImpl) Read() ReadStore {
	return ReadStoreImpl{}
}

// Save Use to save read in BB
func (msi ReadStoreImpl) Save(read *models.Read, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := read.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("readStoreImpl.Save.read.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(read) {
		transaction.Rollback()
		return u.NewLocAppError("readStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
	}
	if err := transaction.Create(&read).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("readStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update read in DB
func (msi ReadStoreImpl) Update(read *models.Read, newRead *models.Read, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := read.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("readStoreImpl.Update.readOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	// if appError := newRead.IsValid(); appError != nil {
	// 	transaction.Rollback()
	// 	return u.NewLocAppError("readStoreImpl.Update.readNew.PreSave", appError.ID, nil, appError.DetailedError)
	// }
	if err := transaction.Model(&read).Updates(&newRead).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("readStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get read from DB
func (msi ReadStoreImpl) GetAll(db *gorm.DB) []models.Read {
	reads := []models.Read{}
	db.Find(&reads)
	return reads
}

// GetByID Used to get read from DB
func (msi ReadStoreImpl) GetByID(ID uint64, db *gorm.DB) models.Read {
	read := models.Read{}
	db.Where("idRead = ?", ID).First(&read)
	return read
}

// GetChannelRead get specific user in specific channel
func (msi ReadStoreImpl) GetChannelRead(user *models.User, channel *models.Channel, db *gorm.DB) models.Read {
	read := models.Read{}
	db.Table("reads").Select("*").Joins("natural join users natural join channels").Where("users.idUser = ? and channels.idChannel = ?", user.IDUser, channel.IDChannel).Find(&read)
	return read
}

// GetByUser get read from user
func (msi ReadStoreImpl) GetByUser(user *models.User, db *gorm.DB) []models.Read {
	reads := []models.Read{}
	db.Table("reads").Select("*").Joins("natural join users").Where("users.idUser = ?", user.IDUser).Find(&reads)
	return reads
}

// GetByChannel get read from channel
func (msi ReadStoreImpl) GetByChannel(channel *models.Channel, db *gorm.DB) []models.Read {
	reads := []models.Read{}
	db.Table("reads").Select("*").Joins("natural join channels").Where("channels.idChannel = ?", channel.IDChannel).Find(&reads)
	return reads
}

// GetByMessage get read from message
func (msi ReadStoreImpl) GetByMessage(message *models.Message, db *gorm.DB) []models.Read {
	reads := []models.Read{}
	db.Table("reads").Select("*").Joins("natural join messages").Where("messages.idMessage = ?", message.IDMessage).Find(&reads)
	return reads
}

// Delete Used to get read from DB
func (msi ReadStoreImpl) Delete(read *models.Read, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := read.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("readStoreImpl.Delete.read.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&read).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("readStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
