package datastores

import (
	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// UserParameterStoreImpl implement UserParameterStore interface
type UserParameterStoreImpl struct{}

// UserParameter Generate the struct for userParameter store
func (s StoreImpl) UserParameter() UserParameterStore {
	return &UserParameterStoreImpl{}
}

// Save Use to save userParameter in BB
func (psi UserParameterStoreImpl) Save(userParameter *models.UserParameter, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	// userParameter.IsValid
	if appError := userParameter.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("userParameterStoreImpl.Save", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(userParameter) {
		transaction.Rollback()
		return u.NewLocAppError("userParameterStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
	}
	if err := transaction.Create(&userParameter).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("userParameterStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update userParameter in DB
func (psi UserParameterStoreImpl) Update(userParameter *models.UserParameter, newUserParameter *models.UserParameter, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := userParameter.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	// if appError := newUserParameter.IsValid(); appError != nil {
	// 	transaction.Rollback()
	// 	return u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", appError.ID, nil, appError.DetailedError)
	// }
	if err := transaction.Model(&userParameter).Updates(&newUserParameter).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("userParameterStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Delete Used to get userParameter from DB
func (psi UserParameterStoreImpl) Delete(userParameter *models.UserParameter, db *gorm.DB) *u.AppError {
	transaction := db.Begin()
	if appError := userParameter.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Delete.userParameter.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&userParameter).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get userParameter from DB
func (psi UserParameterStoreImpl) GetAll(db *gorm.DB) []models.UserParameter {
	userParameters := []models.UserParameter{}
	db.Find(&userParameters)
	return userParameters
}

// GetByUser get userParameter from user
func (psi UserParameterStoreImpl) GetByUser(user *models.User, db *gorm.DB) []models.UserParameter {
	userParameters := []models.UserParameter{}
	db.Table("user_parameters").Select("*").Joins("natural join users").Where("users.idUser = ?", user.IDUser).Find(&userParameters)
	return userParameters
}

// GetByName get userParameter from user
func (psi UserParameterStoreImpl) GetByName(parameterName string, db *gorm.DB) []models.UserParameter {
	userParameters := []models.UserParameter{}
	db.Where("parameterName = ?", parameterName).Find(&userParameters)
	return userParameters
}

// GetByID Used to get userParameter from DB
func (psi UserParameterStoreImpl) GetByID(userID uint64, parameterName string, db *gorm.DB) models.UserParameter {
	userParameter := models.UserParameter{}
	db.Where("idUser = ? AND parameterName= ?", userID, parameterName).First(&userParameter)
	return userParameter
}
