package datastores

import (
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// EmojiStoreImpl implement EmojiStore interface
type EmojiStoreImpl struct {
	EmojiStore
}

// Save Use to save emoji in BB
func (esi EmojiStoreImpl) Save(emoji *models.Emoji, ds dbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := emoji.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("emojiStoreImpl.Save.emoji.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(emoji) {
		transaction.Rollback()
		return u.NewLocAppError("emojiStoreImpl.Save", "save.transaction.create.already_exist", nil, "Emoji Name: "+emoji.Name)
	}
	if err := transaction.Create(&emoji).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("emojiStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update emoji in DB
func (esi EmojiStoreImpl) Update(emoji *models.Emoji, newEmoji *models.Emoji, ds dbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := emoji.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("emojiStoreImpl.Update.emojiOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if appError := newEmoji.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("emojiStoreImpl.Update.emojiNew.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Model(&emoji).Updates(&newEmoji).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("emojiStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get emoji from DB
func (esi EmojiStoreImpl) GetAll(ds dbStore) *[]models.Emoji {
	db := *ds.Db
	emojis := []models.Emoji{}
	db.Find(&emojis)
	return &emojis
}

// GetByName Used to get emoji from DB
func (esi EmojiStoreImpl) GetByName(emojiName string, ds dbStore) *models.Emoji {
	db := *ds.Db
	emoji := models.Emoji{}
	db.Where("name = ?", emojiName).First(&emoji)
	return &emoji
}

// GetByShortcut Used to get emoji from DB
func (esi EmojiStoreImpl) GetByShortcut(EmojiShortcut string, ds dbStore) *models.Emoji {
	db := *ds.Db
	emoji := models.Emoji{}
	db.Where("shortcut = ?", EmojiShortcut).First(&emoji)
	return &emoji
}

// GetByLink Used to get emoji from DB
func (esi EmojiStoreImpl) GetByLink(emojiLink string, ds dbStore) *models.Emoji {
	db := *ds.Db
	emoji := models.Emoji{}
	db.Where("link = ?", emojiLink).First(&emoji)
	return &emoji
}

// Delete Used to get emoji from DB
func (esi EmojiStoreImpl) Delete(emoji *models.Emoji, ds dbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := emoji.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("emojiStoreImpl.Delete.emoji.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&emoji).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("emojiStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
