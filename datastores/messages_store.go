package datastores

import (
	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// MessageStoreImpl Used to implement MessageStore interface
type MessageStoreImpl struct{}

// Message Generate the struct for message store
func (s StoreImpl) Message() MessageStore {
	return &MessageStoreImpl{}
}

// Save Use to save message in BB
func (msi MessageStoreImpl) Save(message *models.Message, db *gorm.DB) *u.AppError {

	transaction := db.Begin()
	message.PreSave()
	if appError := message.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("messageStoreImpl.Save.message.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(message) {
		transaction.Rollback()
		return u.NewLocAppError("messageStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
	}
	if err := transaction.Create(&message).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("messageStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update message in DB
func (msi MessageStoreImpl) Update(message *models.Message, newMessage *models.Message, db *gorm.DB) *u.AppError {

	transaction := db.Begin()
	newMessage.PreSave()
	if appError := message.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("messageStoreImpl.Update.messageOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	// if appError := newMessage.IsValid(); appError != nil {
	// 	transaction.Rollback()
	// 	return u.NewLocAppError("messageStoreImpl.Update.messageNew.PreSave", appError.ID, nil, appError.DetailedError)
	// }
	if err := transaction.Model(&message).Updates(&newMessage).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("messageStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get message from DB
func (msi MessageStoreImpl) GetAll(db *gorm.DB) []models.Message {
	messages := []models.Message{}
	db.Find(&messages)
	return messages
}

// GetByID Used to get message from DB
func (msi MessageStoreImpl) GetByID(ID uint64, db *gorm.DB) models.Message {
	message := models.EmptyMessage
	db.Where("idMessage = ?", ID).First(&message)
	return message
}

// GetByDate Used to get message from DB by specific date
func (msi MessageStoreImpl) GetByDate(messageDate int, db *gorm.DB) []models.Message {
	messages := []models.Message{}
	db.Where("date = ?", messageDate).Find(&messages)
	return messages
}

// GetOrderedByDate get all messages ordered by date
func (msi MessageStoreImpl) GetOrderedByDate(messageDate int, db *gorm.DB) []models.Message {
	messages := []models.Message{}
	db.Order("lastUpdate, messageName, email").Find(&messages)
	return messages
}

// GetByCreator get message from user
func (msi MessageStoreImpl) GetByCreator(creator *models.User, db *gorm.DB) []models.Message {
	messages := []models.Message{}
	db.Table("messages").Select("*").Joins("natural join users").Where("users.idUser = ?", creator.IDUser).Find(&messages)
	return messages
}

// GetByChannel get message from channel
func (msi MessageStoreImpl) GetByChannel(channel *models.Channel, db *gorm.DB) []models.Message {
	messages := []models.Message{}
	db.Table("messages").Select("*").Joins("natural join channels").Where("channels.idChannel = ?", channel.IDChannel).Find(&messages)
	return messages
}

// Delete Used to get message from DB
func (msi MessageStoreImpl) Delete(message *models.Message, db *gorm.DB) *u.AppError {

	transaction := db.Begin()
	if appError := message.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("messageStoreImpl.Delete.message.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&message).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("messageStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
