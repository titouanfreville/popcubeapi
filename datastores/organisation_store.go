package datastores

import (
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// OrganisationStoreImpl implements OrganisationSotre interface
type OrganisationStoreImpl struct {
	OrganisationStore
}

// Save Use to save data in BB
func (osi OrganisationStoreImpl) Save(organisation *models.Organisation, ds dbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	organisation.PreSave()
	if appError := organisation.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(organisation) {
		transaction.Rollback()
		return u.NewLocAppError("organisationStoreImpl.Save", "save.transaction.create.already_exist", nil, "Organisation Name: "+organisation.OrganisationName)
	}
	if err := transaction.Create(&organisation).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("organisationStoreImpl.Save", "save.transaction.create.encounterError: "+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update data in DB
func (osi OrganisationStoreImpl) Update(organisation *models.Organisation, newOrganisation *models.Organisation, ds dbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	newOrganisation.PreSave()
	if appError := organisation.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if appError := newOrganisation.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Model(&organisation).Updates(&newOrganisation).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("organisationStoreImpl.Update", "update.transaction.updates.encounterError: "+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Get Used to get organisation from DB
func (osi OrganisationStoreImpl) Get(ds dbStore) *models.Organisation {
	db := *ds.Db
	organisation := models.Organisation{}
	db.First(&organisation)
	return &organisation
}
