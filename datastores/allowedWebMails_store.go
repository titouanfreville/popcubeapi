package datastores

import (
	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// AllowedWebMailsStoreImpl Used to implement AllowedWebMailsStore interface
type AllowedWebMailsStoreImpl struct{}

// AllowedWebMails Generate the struct for allowedWebMails store
func (s StoreImpl) AllowedWebMails() AllowedWebMailsStore {
	return AllowedWebMailsStoreImpl{}
}

// Save Use to save allowedWebMails in DB
func (asi AllowedWebMailsStoreImpl) Save(allowedWebMails *models.AllowedWebMails, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := allowedWebMails.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewAPIError(422, appError.ID, "Wrong data provided : \n ---- "+appError.Message+" ----")
	}
	if !transaction.NewRecord(&allowedWebMails) {
		transaction.Rollback()
		return u.NewAPIError(409, "duplicate entry", "You already authorized "+allowedWebMails.Domain+" mails to sign up.")
	}
	if err := transaction.Create(&allowedWebMails).Error; err != nil {
		transaction.Rollback()
		return u.NewAPIError(500, "unxepected error", "Unexpected error while adding entry. \n ---- "+err.Error()+"----")
	}
	transaction.Commit()
	return nil
}

// Update Used to update allowedWebMails in DB
func (asi AllowedWebMailsStoreImpl) Update(allowedWebMails *models.AllowedWebMails, newAllowedWebMails *models.AllowedWebMails, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := allowedWebMails.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewAPIError(422, appError.ID, "Wrong data provided : \n ---- "+appError.Message+" ----")
	}
	// if appError := newAllowedWebMails.IsValid(); appError != nil {
	// 	transaction.Rollback()
	// 	return u.NewLocAppError("allowedWebMailsStoreImpl.Update.allowedWebMailsNew.PreSave", appError.ID, nil, appError.DetailedError)
	// }
	if err := transaction.Model(&allowedWebMails).Updates(&newAllowedWebMails).Error; err != nil {
		transaction.Rollback()
		return u.NewAPIError(500, "unxepected error", "Unexpected error while adding entry. \n ---- "+err.Error()+"----")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get allowedWebMails from DB
func (asi AllowedWebMailsStoreImpl) GetAll(db *gorm.DB) []models.AllowedWebMails {
	allowedWebMailss := []models.AllowedWebMails{}
	db.Find(&allowedWebMailss)
	return allowedWebMailss
}

// GetByID Used to get allowedWebMails from DB
func (asi AllowedWebMailsStoreImpl) GetByID(ID uint64, db *gorm.DB) models.AllowedWebMails {
	allowedWebMails := models.AllowedWebMails{}
	db.Where("idAllowedWebMails = ?", ID).First(&allowedWebMails)
	return allowedWebMails
}

// GetByDomain Used to get allowedWebMails having providing domin
func (asi AllowedWebMailsStoreImpl) GetByDomain(domain string, db *gorm.DB) models.AllowedWebMails {
	allowedWebMails := models.AllowedWebMails{}
	db.Where("domain = ?", domain).First(&allowedWebMails)
	return allowedWebMails
}

// GetByProvider Used to get allowedWebMails provided  by ...
func (asi AllowedWebMailsStoreImpl) GetByProvider(provider string, db *gorm.DB) []models.AllowedWebMails {
	allowedWebMails := []models.AllowedWebMails{}
	db.Where("provider = ?", provider).First(&allowedWebMails)
	return allowedWebMails

}

// Delete Used to remove specified allowedWebMails from DB
func (asi AllowedWebMailsStoreImpl) Delete(allowedWebMails *models.AllowedWebMails, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := allowedWebMails.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("allowedWebMailsStoreImpl.Delete.allowedWebMails.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&allowedWebMails).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("allowedWebMailsStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
