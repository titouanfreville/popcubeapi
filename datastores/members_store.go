package datastores

import (
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// MemberStoreImpl Used to implement MemberStore interface
type MemberStoreImpl struct{}

// NewMemberStore Generate the struct for member store
func NewMemberStore() MemberStore {
	return &MemberStoreImpl{}
}

// Save Use to save member in BB
func (msi MemberStoreImpl) Save(member *models.Member, ds DbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := member.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("memberStoreImpl.Save.member.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.Debug().NewRecord(member) {
		transaction.Rollback()
		return u.NewLocAppError("memberStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
	}
	if err := transaction.Debug().Create(&member).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("memberStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update member in DB
func (msi MemberStoreImpl) Update(member *models.Member, newMember *models.Member, ds DbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := member.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("memberStoreImpl.Update.memberOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if appError := newMember.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("memberStoreImpl.Update.memberNew.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Model(&member).Updates(&newMember).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("memberStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get member from DB
func (msi MemberStoreImpl) GetAll(ds DbStore) *[]models.Member {
	db := *ds.Db
	members := []models.Member{}
	db.Find(&members)
	return &members
}

// GetByUser get member from user
func (msi MemberStoreImpl) GetByUser(user *models.User, ds DbStore) *[]models.Member {
	db := *ds.Db
	members := []models.Member{}
	db.Table("members").Select("*").Joins("natural join users").Where("users.idUser = ?", user.IDUser).Find(&members)
	return &members
}

// GetByChannel get member from channel
func (msi MemberStoreImpl) GetByChannel(channel *models.Channel, ds DbStore) *[]models.Member {
	db := *ds.Db
	members := []models.Member{}
	db.Table("members").Select("*").Joins("natural join channels").Where("channels.idChannel = ?", channel.IDChannel).Find(&members)
	return &members
}

// GetByRole get member from role
func (msi MemberStoreImpl) GetByRole(role *models.Role, ds DbStore) *[]models.Member {
	db := *ds.Db
	members := []models.Member{}
	db.Table("members").Select("*").Joins("natural join roles").Where("roles.idRole = ?", role.IDRole).Find(&members)
	return &members
}

// Delete Used to get member from DB
func (msi MemberStoreImpl) Delete(member *models.Member, ds DbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := member.IsValid(); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("memberStoreImpl.Delete.member.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&member).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("memberStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
