package datastores

import (
	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// ParameterStoreImpl implement ParameterStore interface
type ParameterStoreImpl struct{}

// Parameter Generate the struct for parameter store
func (s StoreImpl) Parameter() ParameterStore {
	return &ParameterStoreImpl{}
}

// Save Use to save parameter in BB
func (psi ParameterStoreImpl) Save(parameter *models.Parameter, db *gorm.DB) *u.AppError {

	transaction := db.Begin()
	if appError := parameter.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(parameter) {
		transaction.Rollback()
		return u.NewLocAppError("parameterStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
	}
	if err := transaction.Create(&parameter).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("parameterStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update parameter in DB
func (psi ParameterStoreImpl) Update(parameter *models.Parameter, newParameter *models.Parameter, db *gorm.DB) *u.AppError {

	transaction := db.Begin()
	if appError := parameter.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if appError := newParameter.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Model(&parameter).Updates(&newParameter).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("parameterStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Get Used to get parameter from DB
func (psi ParameterStoreImpl) Get(db *gorm.DB) models.Parameter {
	parameter := models.Parameter{}
	db.First(&parameter)
	return parameter
}
