package datastores

import (
	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// ChannelStoreImpl Used to implement ChannelStore interface
type ChannelStoreImpl struct{}

// Channel Generate the struct for channel store
func (s StoreImpl) Channel() ChannelStore {
	return &ChannelStoreImpl{}
}

// Save Use to save channel in BB
func (csi ChannelStoreImpl) Save(channel *models.Channel, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	channel.PreSave()
	if appError := channel.IsValid(false); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("channelStoreImpl.Save.channel.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(channel) {
		transaction.Rollback()
		return u.NewLocAppError("channelStoreImpl.Save", "save.transaction.create.already_exist", nil, "Channel Name: "+channel.ChannelName)
	}
	if err := transaction.Create(&channel).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("channelStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update channel in DB
func (csi ChannelStoreImpl) Update(channel *models.Channel, newChannel *models.Channel, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	newChannel.PreUpdate()
	newChannel.WebID = channel.WebID
	if appError := channel.IsValid(true); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	// if appError := newChannel.IsValid(true); appError != nil {
	// 	transaction.Rollback()
	// 	return u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", appError.ID, nil, appError.DetailedError)
	// }

	if err := transaction.Model(&channel).Update(&newChannel).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("channelStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}

	if !newChannel.Private {
		if err := transaction.Model(&channel).Updates("private", false).Error; err != nil {
			transaction.Rollback()
			return u.NewLocAppError("channelStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
		}
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get channel from DB
func (csi ChannelStoreImpl) GetAll(db *gorm.DB) []models.Channel {
	channels := []models.Channel{}
	db.Find(&channels)
	return channels
}

// GetByID Used to get channel from DB
func (csi ChannelStoreImpl) GetByID(ID uint64, db *gorm.DB) models.Channel {
	channel := models.Channel{}
	db.Where("idChannel = ?", ID).First(&channel)
	return channel
}

// GetByName Used to get channel from DB
func (csi ChannelStoreImpl) GetByName(channelName string, db *gorm.DB) models.Channel {

	channel := models.Channel{}
	db.Where("channelName = ?", channelName).First(&channel)
	return channel
}

// GetByType allow to find channels by types.
func (csi ChannelStoreImpl) GetByType(channelType string, db *gorm.DB) []models.Channel {
	channels := []models.Channel{}
	db.Where("type = ?", channelType).Find(&channels)
	return channels
}

// GetPublic allow to find publics channels.
func (csi ChannelStoreImpl) GetPublic(db *gorm.DB) []models.Channel {
	channels := []models.Channel{}
	db.Where("private = ?", false).Find(&channels)
	return channels
}

// GetPrivate allow to find publics channels.
func (csi ChannelStoreImpl) GetPrivate(db *gorm.DB) []models.Channel {
	channels := []models.Channel{}
	db.Where("private = ?", true).Find(&channels)
	return channels
}

// Delete Used to get channel from DB
func (csi ChannelStoreImpl) Delete(channel *models.Channel, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := channel.IsValid(true); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("channelStoreImpl.Delete.channel.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&channel).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("channelStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
