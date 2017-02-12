package datastores

import (
	"fmt"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// UserStoreImpl Used to implement UserStore interface
type UserStoreImpl struct {
	UserStore
}

// Save Use to save user in BB
func (usi UserStoreImpl) Save(user *models.User, ds DbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	user.PreSave()
	if appError := user.IsValid(false); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Save.user.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(user) {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.already_exist", nil, "User Name: "+user.Username)
	}
	fmt.Printf("IDRole is : %d, InUserRole ID is : %d", user.IDRole, user.Role.IDRole)
	if err := transaction.Debug().Create(&user).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update user in DB
func (usi UserStoreImpl) Update(user *models.User, newUser *models.User, ds DbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	newUser.PreUpdate()
	if appError := user.IsValid(false); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Update.userOld.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if appError := newUser.IsValid(true); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Model(&user).Updates(&newUser).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Update", "update.transaction.updates.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// GetAll Used to get user from DB
func (usi UserStoreImpl) GetAll(ds DbStore) *[]models.User {
	db := *ds.Db
	users := []models.User{}
	db.Find(&users)
	return &users
}

// GetByUserName Used to get user from DB
func (usi UserStoreImpl) GetByUserName(userName string, ds DbStore) *models.User {
	db := *ds.Db
	user := models.User{}
	db.Where("userName = ?", userName).First(&user)
	return &user
}

// GetByEmail Used to get user from DB by email
func (usi UserStoreImpl) GetByEmail(userEmail string, ds DbStore) *models.User {
	db := *ds.Db
	user := models.User{}
	db.Where("email = ?", userEmail).First(&user)
	return &user
}

// GetOrderedByDate get all users ordered by date
func (usi UserStoreImpl) GetOrderedByDate(userDate int, ds DbStore) *[]models.User {
	db := *ds.Db
	users := []models.User{}
	db.Order("updatedAt, userName, email").Find(&users)
	return &users
}

// GetDeleted get deleted users
func (usi UserStoreImpl) GetDeleted(ds DbStore) *[]models.User {
	db := *ds.Db
	users := []models.User{}
	db.Where("deleted = ?", true).First(&users)
	return &users
}

// GetByNickName get user from nick name
func (usi UserStoreImpl) GetByNickName(nickName string, ds DbStore) *models.User {
	db := *ds.Db
	user := models.User{}
	db.Where("nickName = ?", nickName).First(&user)
	return &user
}

// GetByFirstName get user by first name
func (usi UserStoreImpl) GetByFirstName(firstName string, ds DbStore) *[]models.User {
	db := *ds.Db
	users := []models.User{}
	db.Where("firstName = ?", firstName).Find(&users)
	return &users
}

// GetByLastName get user from last name
func (usi UserStoreImpl) GetByLastName(lastName string, ds DbStore) *[]models.User {
	db := *ds.Db
	users := []models.User{}
	db.Where("lastName = ?", lastName).Find(&users)
	return &users
}

// GetByRole get user from rolme
func (usi UserStoreImpl) GetByRole(role *models.Role, ds DbStore) *[]models.User {
	db := *ds.Db
	users := []models.User{}
	db.Debug().Model(role).Related(&users, "Role")
	// db.Debug().Model(&role).Association("Role").Find(&users)
	return &users
}

// Need MEMEBER functions to do it
// GetByChannel Get user in a channem
// func (usi UserStoreImpl) GetByChannel(channel *models.Channel, ds DbStore) *[]models.User {

// }

//

// Delete Used to get user from DB
func (usi UserStoreImpl) Delete(user *models.User, ds DbStore) *u.AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := user.IsValid(true); appError != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Delete.user.PreSave", appError.ID, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&user).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Delete", "update.transaction.delete.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
