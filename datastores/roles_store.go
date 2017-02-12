package datastores

import (
	"github.com/titouanfreville/popcubeapi/models"
	"strings"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// RoleStoreImpl Used to implement RoleStore interface
type RoleStoreImpl struct {
	RoleStore
}

// Save Use to save role in BB
func (rsi RoleStoreImpl) Save(role *models.Role, ds dbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := role.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("roleStoreImpl.Save.role.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(role) {
		transaction.Rollback()
		return u.NewLocAppError("roleStoreImpl.Save", "save.transaction.create.already_exist", nil, "Role Name: "+role.RoleName)
	}
	if err := transaction.Create(&role).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("roleStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update role in DB
func (rsi RoleStoreImpl) Update(role *models.Role, newRole *models.Role, ds dbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := role.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("roleStoreImpl.Update.roleOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if appError := newRole.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("roleStoreImpl.Update.roleNew.PreSave", appError.ID, nil, appError.DetailedError)
	}
	transitionRole := models.Role{
		CanUsePrivate: newRole.CanUsePrivate,
		CanModerate:   newRole.CanModerate,
		CanArchive:    newRole.CanArchive,
		CanInvite:     newRole.CanInvite,
		CanManage:     newRole.CanManage,
		CanManageUser: newRole.CanManageUser,
	}
	json := transitionRole.ToJSON()
	rights := u.StringInterfaceFromJSON(strings.NewReader(json))
	if err := transaction.Model(&role).Updates(&newRole).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("roleStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	if err := transaction.Model(&role).Updates(rights).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("roleStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get role from DB
func (rsi RoleStoreImpl) GetAll(ds dbStore) *[]models.Role {
	db := *ds.Db
	roles := []models.Role{}
	db.Find(&roles)
	return &roles
}

// GetByName Used to get role from DB
func (rsi RoleStoreImpl) GetByName(roleName string, ds dbStore) *models.Role {
	db := *ds.Db
	role := models.Role{}
	db.Where("roleName = ?", roleName).First(&role)
	return &role
}

// GetByRights Used to get role from DB
// You can only search for roles set to true.
func (rsi RoleStoreImpl) GetByRights(roleRights *models.Role, ds dbStore) *[]models.Role {
	db := *ds.Db
	roles := []models.Role{}
	db.Where(&roleRights).Find(&roles)
	return &roles
}

// Delete Used to get role from DB
func (rsi RoleStoreImpl) Delete(role *models.Role, ds dbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := role.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("roleStoreImpl.Delete.role.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&role).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("roleStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
