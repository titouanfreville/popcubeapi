package datastores

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// RoleStoreImpl Used to implement RoleStore interface
type RoleStoreImpl struct{}

// Role Generate the struct for role store
func (s StoreImpl) Role() RoleStore {
	return RoleStoreImpl{}
}

// Save Use to save role in BB
func (rsi RoleStoreImpl) Save(role *models.Role, db *gorm.DB) *u.AppError {
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
func (rsi RoleStoreImpl) Update(role *models.Role, newRole *models.Role, db *gorm.DB) *u.AppError {
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

// GetByID Used to get role from DB
func (rsi RoleStoreImpl) GetByID(ID uint64, db *gorm.DB) models.Role {
	role := models.Role{}
	db.Where("idRole = ?", ID).First(&role)
	return role
}

// GetAll Used to get role from DB
func (rsi RoleStoreImpl) GetAll(db *gorm.DB) []models.Role {
	roles := []models.Role{}
	db.Find(&roles)
	return roles
}

// GetByName Used to get role from DB
func (rsi RoleStoreImpl) GetByName(roleName string, db *gorm.DB) models.Role {
	role := models.Role{}
	db.Where("roleName = ?", roleName).First(&role)
	return role
}

// GetByRights Used to get role from DB
// You can only search for roles set to true.
func (rsi RoleStoreImpl) GetByRights(roleRights *models.Role, db *gorm.DB) []models.Role {
	roles := []models.Role{}
	db.Where(&roleRights).Find(&roles)
	return roles
}

// Delete Used to get role from DB
func (rsi RoleStoreImpl) Delete(role *models.Role, db *gorm.DB) *u.AppError {
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
