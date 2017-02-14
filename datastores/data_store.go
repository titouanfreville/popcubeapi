/*Package datastores implements the basics databases communication functions used by PopCube chat api.

Stores

The following is a list of stores described:
	Avatar: Contain all informations for avatar management
	Channel: Contain all informations for channel management
	Emojis: Contain all informations for emojis management
	Organisation: Contain all informations for organisation management
	Parameter: Contain all informations for parmeters management
	Role: Contain all informations for roles management
	User: Contain all informations for users management
*/
// Created by Titouan FREVILLE <titouanfreville@gmail.com>
//
// Inspired by mattermost project
package datastores

import (
	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"

	// Importing sql driver. They are used by gorm package and used by default from blank.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DbStore Struct to manage Db knowledge
type DbStore struct {
	Db  *gorm.DB
	Err error
}

// InitConnection init Database connection && database models
func (ds *DbStore) InitConnection(user string, dbname string, password string, host string, port string) {
	connectionChain := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	db, _ := gorm.Open("mysql", connectionChain)

	// Create correct tables
	db.AutoMigrate(&models.Avatar{}, &models.Channel{}, &models.Emoji{}, &models.Folder{},
		models.Member{}, &models.Message{}, &models.Organisation{}, &models.Parameter{},
		&models.Role{}, &models.User{})

	// Will not set CreatedAt and LastUpdate on .Create() call
	db.Callback().Create().Remove("gorm:update_time_stamp")
	db.Callback().Create().Remove("gorm:save_associations")

	// Will not update LastUpdate on .Save() call
	db.Callback().Update().Remove("gorm:update_time_stamp")
	db.Callback().Update().Remove("gorm:save_associations")

	ds.Db = db

	if db.NewRecord(models.Owner) {
		ds.roleInitSave(models.Owner)
	}
	if db.NewRecord(&models.Admin) {
		ds.roleInitSave(models.Admin)
	}
	if db.NewRecord(&models.Standart) {
		ds.roleInitSave(models.Standart)
	}
}

func (ds *DbStore) roleInitSave(role models.Role) *u.AppError {
	db := ds.Db
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

// CloseConnection close database connection
func (ds *DbStore) CloseConnection() {
	db := *ds.Db
	defer db.Close()
	ds.Db = &gorm.DB{}
}

// Store interface the Stores and usefull DB functions
type Store interface {
	Organisation() OrganisationStore
	Avatar() AvatarStore
	Emoji() EmojiStore
	InitConnection()
	CloseConnection()
}

/*AvatarStore interface the avatar communication */
type AvatarStore interface {
	Save(avatar *models.Avatar, ds DbStore) *u.AppError
	Update(avatar *models.Avatar, newAvatar *models.Avatar, ds DbStore) *u.AppError
	GetByName(avatarName string, ds DbStore) *models.Avatar
	GetByLink(avatarLink string, ds DbStore) *models.Avatar
	GetAll(ds DbStore) *[]models.Avatar
	Delete(avatar *models.Avatar, ds DbStore) *u.AppError
}

/*ChannelStore interface the channel communication*/
type ChannelStore interface {
	Save(channel *models.Channel, ds DbStore) *u.AppError
	Update(channel *models.Channel, newChannel *models.Channel, ds DbStore) *u.AppError
	GetByName(channelName string, ds DbStore) *models.Channel
	GetByType(channelType string, ds DbStore) *[]models.Channel
	GetPublic(ds DbStore) *[]models.Channel
	GetPrivate(ds DbStore) *[]models.Channel
	GetAll(ds DbStore) *[]models.Channel
	Delete(channel *models.Channel, ds DbStore) *u.AppError
}

/*EmojiStore interface the emoji communication*/
type EmojiStore interface {
	Save(emoji *models.Emoji, ds DbStore) *u.AppError
	Update(emoji *models.Emoji, newEmoji *models.Emoji, ds DbStore) *u.AppError
	GetByName(emojiName string, ds DbStore) *models.Emoji
	GetByShortcut(emojiShortcut string, ds DbStore) *models.Emoji
	GetByLink(emojiLink string, ds DbStore) *models.Emoji
	GetAll(ds DbStore) *[]models.Emoji
	Delete(emoji *models.Emoji, ds DbStore) *u.AppError
}

/*FolderStore interface communication with message table*/
type FolderStore interface {
	Save(message *models.Folder, ds DbStore) *u.AppError
	Update(message *models.Folder, newFolder *models.Folder, ds DbStore) *u.AppError
	GetByName(messageName string, ds DbStore) *[]models.Folder
	GetByType(messageType string, ds DbStore) *[]models.Folder
	GetByLink(messageLink string, ds DbStore) *[]models.Folder
	GetByMessage(message *models.Message, ds DbStore) *[]models.Folder
	GetAll(ds DbStore) *[]models.Folder
	Delete(message *models.Folder, ds DbStore) *u.AppError
}

/*MemberStore interface communication with member table*/
type MemberStore interface {
	Save(member *models.Member, ds DbStore) *u.AppError
	Update(member *models.Member, newMember *models.Member, ds DbStore) *u.AppError
	GetByUser(user *models.User, ds DbStore) *[]models.Member
	GetByChannel(channel *models.Channel, ds DbStore) *[]models.Member
	GetByRole(role *models.Role, ds DbStore) *[]models.Member
	GetAll(ds DbStore) *[]models.Member
	Delete(member *models.Member, ds DbStore) *u.AppError
}

/*MessageStore interface communication with message table*/
type MessageStore interface {
	Save(message *models.Message, ds DbStore) *u.AppError
	Update(message *models.Message, newMessage *models.Message, ds DbStore) *u.AppError
	GetByDate(messageDate int, ds DbStore) *[]models.Message
	GetByCreator(creator *models.User, ds DbStore) *[]models.Message
	GetByChannel(channel *models.Channel, ds DbStore) *[]models.Message
	GetAll(ds DbStore) *[]models.Message
	Delete(message *models.Message, ds DbStore) *u.AppError
}

/*OrganisationStore interface the organisation communication
Organisation is unique in the database. So they are no use of providing an user to get.
Delete is useless as we will down the docker stack in case an organisation leace.
*/
type OrganisationStore interface {
	Save(organisation *models.Organisation, ds DbStore) *u.AppError
	Update(organisation *models.Organisation, newOrganisation *models.Organisation, ds DbStore) *u.AppError
	Get(ds DbStore) *models.Organisation
}

/*ParameterStore interface the parameter communication*/
type ParameterStore interface {
	Save(parameter *models.Parameter, ds DbStore) *u.AppError
	Update(parameter *models.Parameter, newParameter *models.Parameter, ds DbStore) *u.AppError
	Get(ds DbStore) *models.Parameter
}

/*RoleStore interface the role communication*/
type RoleStore interface {
	Save(role *models.Role, ds DbStore) *u.AppError
	Update(role *models.Role, newRole *models.Role, ds DbStore) *u.AppError
	GetByName(roleName string, ds DbStore) *models.Role
	GetByRights(roleRights *models.Role, ds DbStore) *[]models.Role
	GetAll(ds DbStore) *[]models.Role
	Delete(role *models.Role, ds DbStore) *u.AppError
}

/*UserStore interface the user communication*/
type UserStore interface {
	Save(user *models.User, ds DbStore) *u.AppError
	Update(user *models.User, newUser *models.User, ds DbStore) *u.AppError
	GetByUserName(userName string, ds DbStore) *models.User
	GetByEmail(userEmail string, ds DbStore) *models.User
	GetOrderedByDate(userDate int, ds DbStore) *[]models.User
	GetDeleted(ds DbStore) *[]models.User
	GetByNickName(nickName string, ds DbStore) *models.User
	GetByFirstName(firstName string, ds DbStore) *[]models.User
	GetByLastName(lastName string, ds DbStore) *[]models.User
	GetByRole(role *models.Role, ds DbStore) *[]models.User
	GetAll(ds DbStore) *[]models.User
	Delete(user *models.User, ds DbStore) *u.AppError
}

// GetByChannel(channel *models.Channel, ds DbStore) *[]models.User
