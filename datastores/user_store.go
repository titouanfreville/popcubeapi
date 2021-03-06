package datastores

import (
	"github.com/jinzhu/gorm"
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// UserStoreImpl Used to implement UserStore interface
type UserStoreImpl struct{}

// User Generate the struct for user store
func (s StoreImpl) User() UserStore {
	return UserStoreImpl{}
}

// Save Use to save user in BB
func (usi UserStoreImpl) Save(user *models.User, db *gorm.DB) *u.AppError {

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
	if err := transaction.Create(&user).Error; err != nil {
		transaction.Rollback()
		return u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.encounterError :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Update Used to update user in DB
func (usi UserStoreImpl) Update(user *models.User, newUser *models.User, db *gorm.DB) *u.AppError {
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
func (usi UserStoreImpl) GetAll(db *gorm.DB) []models.User {
	users := []models.User{}
	db.Find(&users)
	return users
}

// GetByID Used to get user from DB
func (usi UserStoreImpl) GetByID(ID uint64, db *gorm.DB) models.User {
	user := models.User{}
	db.Where("idUser = ?", ID).First(&user)
	return user
}

// GetByUserName Used to get user from DB
func (usi UserStoreImpl) GetByUserName(userName string, db *gorm.DB) models.User {
	user := models.User{}
	db.Where("userName = ?", userName).First(&user)
	return user
}

// Login Used to log user in
func (usi UserStoreImpl) Login(userName string, pass string, db *gorm.DB) (models.User, *u.AppError) {
	user := models.User{}
	empty := models.User{}
	err := u.NewAPIError(404, "Wrong user name or password", "Can't proceed to login. Password or user name is not correct")
	db.Where("userName = ?", userName).First(&user)
	if (user == models.User{}) {
		return empty, err
	}
	if models.ComparePassword(user.Password, pass) {
		return user, nil
	}
	return empty, err
}

// GetByEmail Used to get user from DB by email
func (usi UserStoreImpl) GetByEmail(userEmail string, db *gorm.DB) models.User {
	user := models.User{}
	db.Where("email = ?", userEmail).First(&user)
	return user
}

// GetOrderedByDate get all users ordered by date
func (usi UserStoreImpl) GetOrderedByDate(userDate int, db *gorm.DB) []models.User {
	users := []models.User{}
	db.Order("lastUpdate, userName, email").Find(&users)
	return users
}

// GetDeleted get deleted users
func (usi UserStoreImpl) GetDeleted(db *gorm.DB) []models.User {
	users := []models.User{}
	db.Where("deleted = ?", true).First(&users)
	return users
}

// GetByNickName get user from nick name
func (usi UserStoreImpl) GetByNickName(nickName string, db *gorm.DB) models.User {
	user := models.User{}
	db.Where("nickName = ?", nickName).First(&user)
	return user
}

// GetByFirstName get user by first name
func (usi UserStoreImpl) GetByFirstName(firstName string, db *gorm.DB) []models.User {
	users := []models.User{}
	db.Where("firstName = ?", firstName).Find(&users)
	return users
}

// GetByLastName get user from last name
func (usi UserStoreImpl) GetByLastName(lastName string, db *gorm.DB) []models.User {
	users := []models.User{}
	db.Where("lastName = ?", lastName).Find(&users)
	return users
}

// GetByRole get user from rolme
func (usi UserStoreImpl) GetByRole(role *models.Role, db *gorm.DB) []models.User {
	users := []models.User{}
	db.Table("users").Select("*").Joins("natural join roles").Where("roles.idRole = ?", role.IDRole).Find(&users)
	return users
}

// Need MEMEBER functions to do it
// GetByChannel Get user in a channem
// func (usi UserStoreImpl) GetByChannel(channel *models.Channel, db *gorm.DB) *[]models.User {
// }

//

// Delete Used to get user from DB
func (usi UserStoreImpl) Delete(user *models.User, db *gorm.DB) *u.AppError {
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
