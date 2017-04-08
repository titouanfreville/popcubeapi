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
	"log"

	"github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"

	// Importing sql driver. They are used by gorm package and used by default from blank.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// StoreInterface interface the Stores and usefull DB functions
type StoreInterface interface {
	AllowedWebMails() AllowedWebMailsStore
	Avatar() AvatarStore
	Channel() ChannelStore
	Emoji() EmojiStore
	Folder() FolderStore
	Member() MemberStore
	Message() MessageStore
	Organisation() OrganisationStore
	Parameter() ParameterStore
	Read() ReadStore
	Role() RoleStore
	UserParameter() UserParameterStore
	User() UserStore
	InitConnection(user string, dbname string, password string, host string, port string) *gorm.DB
	InitDatabase(user string, dbname string, password string, host string, port string)
	CloseConnection(*gorm.DB)
}

// StoreImpl implement store interface
type StoreImpl struct{}

// Store init store
func Store() StoreInterface {
	return StoreImpl{}
}

// InitConnection init Database connection && database models
func (store StoreImpl) InitConnection(user string, dbname string, password string, host string, port string) *gorm.DB {
	connectionChain := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, _ := gorm.Open("mysql", connectionChain)

	// Will not set CreatedAt and LastUpdate on .Create() call
	db.Callback().Create().Remove("gorm:update_time_stamp")
	db.Callback().Create().Remove("gorm:save_associations")

	// Will not update LastUpdate on .Save() call
	db.Callback().Update().Remove("gorm:update_time_stamp")
	db.Callback().Update().Remove("gorm:save_associations")

	if err := db.DB().Ping(); err != nil {
		log.Print("Can't connect to database")
		log.Print(host)
		return nil
	}
	return db
}

func (store StoreImpl) roleInitSave(role models.Role, db *gorm.DB) *u.AppError {
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

// InitDatabase initialise a connection to the database and the database.
func (store StoreImpl) InitDatabase(user string, dbname string, password string, host string, port string) {
	db := store.InitConnection(user, dbname, password, host, port)
	db.Debug().DB().Ping()
	// Create correct tables
	// db.AutoMigrate(&models.EmptyAvatar, &models.Channel{}, &models.Emoji{}, &models.Folder{},
	// 	models.Member{}, &models.Message{}, &models.Organisation{}, &models.Parameter{},
	// 	&models.Role{}, &models.User{})

	// Will not set CreatedAt and LastUpdate on .Create() call
	db.Callback().Create().Remove("gorm:update_time_stamp")
	db.Callback().Create().Remove("gorm:save_associations")

	// Will not update LastUpdate on .Save() call
	db.Callback().Update().Remove("gorm:update_time_stamp")
	db.Callback().Update().Remove("gorm:save_associations")

	db.Debug().DB().Ping()

	if db.NewRecord(models.Owner) {
		store.roleInitSave(models.Owner, db)
	}
	if db.NewRecord(&models.Admin) {
		store.roleInitSave(models.Admin, db)
	}
	if db.NewRecord(&models.Standart) {
		store.roleInitSave(models.Standart, db)
	}
}

// CloseConnection close database connection
func (store StoreImpl) CloseConnection(db *gorm.DB) {
	defer db.Close()
}

/*AllowedWebMailsStore interface the allowedWebMails communication */
type AllowedWebMailsStore interface {
	Save(allowedWebMails *models.AllowedWebMails, db *gorm.DB) *u.AppError
	Update(allowedWebMails *models.AllowedWebMails, newAllowedWebMails *models.AllowedWebMails, db *gorm.DB) *u.AppError
	GetByDomain(domain string, db *gorm.DB) models.AllowedWebMails
	GetByProvider(provider string, db *gorm.DB) []models.AllowedWebMails
	GetByID(ID uint64, db *gorm.DB) models.AllowedWebMails
	GetAll(db *gorm.DB) []models.AllowedWebMails
	Delete(allowedWebMails *models.AllowedWebMails, db *gorm.DB) *u.AppError
}

/*AvatarStore interface the avatar communication */
type AvatarStore interface {
	Save(avatar *models.Avatar, db *gorm.DB) *u.AppError
	Update(avatar *models.Avatar, newAvatar *models.Avatar, db *gorm.DB) *u.AppError
	GetByName(avatarName string, db *gorm.DB) models.Avatar
	GetByLink(avatarLink string, db *gorm.DB) models.Avatar
	GetByID(ID uint64, db *gorm.DB) models.Avatar
	GetAll(db *gorm.DB) []models.Avatar
	Delete(avatar *models.Avatar, db *gorm.DB) *u.AppError
}

/*ChannelStore interface the channel communication*/
type ChannelStore interface {
	Save(channel *models.Channel, db *gorm.DB) *u.AppError
	Update(channel *models.Channel, newChannel *models.Channel, db *gorm.DB) *u.AppError
	GetByID(ID uint64, db *gorm.DB) models.Channel
	GetByName(channelName string, db *gorm.DB) models.Channel
	GetByType(channelType string, db *gorm.DB) []models.Channel
	GetPublic(db *gorm.DB) []models.Channel
	GetPrivate(db *gorm.DB) []models.Channel
	GetAll(db *gorm.DB) []models.Channel
	Delete(channel *models.Channel, db *gorm.DB) *u.AppError
}

/*EmojiStore interface the emoji communication*/
type EmojiStore interface {
	Save(emoji *models.Emoji, db *gorm.DB) *u.AppError
	Update(emoji *models.Emoji, newEmoji *models.Emoji, db *gorm.DB) *u.AppError
	GetByName(emojiName string, db *gorm.DB) models.Emoji
	GetByShortcut(emojiShortcut string, db *gorm.DB) models.Emoji
	GetByLink(emojiLink string, db *gorm.DB) models.Emoji
	GetByID(ID uint64, db *gorm.DB) models.Emoji
	GetAll(db *gorm.DB) []models.Emoji
	Delete(emoji *models.Emoji, db *gorm.DB) *u.AppError
}

/*FolderStore interface communication with message table*/
type FolderStore interface {
	Save(message *models.Folder, db *gorm.DB) *u.AppError
	Update(message *models.Folder, newFolder *models.Folder, db *gorm.DB) *u.AppError
	GetByID(ID uint64, db *gorm.DB) models.Folder
	GetByName(messageName string, db *gorm.DB) []models.Folder
	GetByType(messageType string, db *gorm.DB) []models.Folder
	GetByLink(messageLink string, db *gorm.DB) []models.Folder
	GetByMessage(message *models.Message, db *gorm.DB) []models.Folder
	GetAll(db *gorm.DB) []models.Folder
	Delete(message *models.Folder, db *gorm.DB) *u.AppError
}

/*MemberStore interface communication with member table*/
type MemberStore interface {
	Save(member *models.Member, db *gorm.DB) *u.AppError
	Update(member *models.Member, newMember *models.Member, db *gorm.DB) *u.AppError
	GetByID(channelID uint64, userID uint64, db *gorm.DB) models.Member
	GetChannelMember(user *models.User, channel *models.Channel, db *gorm.DB) models.Member
	GetByUser(user *models.User, db *gorm.DB) []models.Member
	GetByChannel(channel *models.Channel, db *gorm.DB) []models.Member
	GetByRole(role *models.Role, db *gorm.DB) []models.Member
	GetAll(db *gorm.DB) []models.Member
	Delete(member *models.Member, db *gorm.DB) *u.AppError
}

/*MessageStore interface communication with message table*/
type MessageStore interface {
	Save(message *models.Message, db *gorm.DB) *u.AppError
	Update(message *models.Message, newMessage *models.Message, db *gorm.DB) *u.AppError
	GetByID(ID uint64, db *gorm.DB) models.Message
	GetByDate(messageDate int, db *gorm.DB) []models.Message
	GetByCreator(creator *models.User, db *gorm.DB) []models.Message
	GetByChannel(channel *models.Channel, db *gorm.DB) []models.Message
	GetAll(db *gorm.DB) []models.Message
	Delete(message *models.Message, db *gorm.DB) *u.AppError
}

/*OrganisationStore interface the organisation communication
Organisation is unique in the database. So they are no use of providing an user to get.
Delete is useless as we will down the docker stack in case an organisation leace.
*/
type OrganisationStore interface {
	Save(organisation *models.Organisation, db *gorm.DB) *u.AppError
	Update(organisation *models.Organisation, newOrganisation *models.Organisation, db *gorm.DB) *u.AppError
	Get(db *gorm.DB) models.Organisation
}

/*ParameterStore interface the parameter communication*/
type ParameterStore interface {
	Save(parameter *models.Parameter, db *gorm.DB) *u.AppError
	Update(parameter *models.Parameter, newParameter *models.Parameter, db *gorm.DB) *u.AppError
	Get(db *gorm.DB) models.Parameter
}

/*ReadStore interface communication with read table*/
type ReadStore interface {
	Save(read *models.Read, db *gorm.DB) *u.AppError
	Update(read *models.Read, newRead *models.Read, db *gorm.DB) *u.AppError
	GetByID(ID uint64, db *gorm.DB) models.Read
	GetChannelRead(user *models.User, channel *models.Channel, db *gorm.DB) models.Read
	GetByUser(user *models.User, db *gorm.DB) []models.Read
	GetByChannel(channel *models.Channel, db *gorm.DB) []models.Read
	GetByMessage(message *models.Message, db *gorm.DB) []models.Read
	GetAll(db *gorm.DB) []models.Read
	Delete(read *models.Read, db *gorm.DB) *u.AppError
}

/*RoleStore interface the role communication*/
type RoleStore interface {
	Save(role *models.Role, db *gorm.DB) *u.AppError
	Update(role *models.Role, newRole *models.Role, db *gorm.DB) *u.AppError
	GetByID(ID uint64, db *gorm.DB) models.Role
	GetByName(roleName string, db *gorm.DB) models.Role
	GetByRights(roleRights *models.Role, db *gorm.DB) []models.Role
	GetAll(db *gorm.DB) []models.Role
	Delete(role *models.Role, db *gorm.DB) *u.AppError
}

/*UserStore interface the user communication*/
type UserStore interface {
	Save(user *models.User, db *gorm.DB) *u.AppError
	Update(user *models.User, newUser *models.User, db *gorm.DB) *u.AppError
	GetByID(ID uint64, db *gorm.DB) models.User
	GetByUserName(userName string, db *gorm.DB) models.User
	GetByEmail(userEmail string, db *gorm.DB) models.User
	GetOrderedByDate(userDate int, db *gorm.DB) []models.User
	GetDeleted(db *gorm.DB) []models.User
	GetByNickName(nickName string, db *gorm.DB) models.User
	GetByFirstName(firstName string, db *gorm.DB) []models.User
	GetByLastName(lastName string, db *gorm.DB) []models.User
	GetByRole(role *models.Role, db *gorm.DB) []models.User
	GetAll(db *gorm.DB) []models.User
	Delete(user *models.User, db *gorm.DB) *u.AppError
	Login(userName string, pass string, db *gorm.DB) (models.User, *u.AppError)
}

/*UserParameterStore interface the userUserParameter communication*/
type UserParameterStore interface {
	Save(userUserParameter *models.UserParameter, db *gorm.DB) *u.AppError
	Update(userUserParameter *models.UserParameter, newUserParameter *models.UserParameter, db *gorm.DB) *u.AppError
	Delete(userParameter *models.UserParameter, db *gorm.DB) *u.AppError
	GetAll(db *gorm.DB) []models.UserParameter
	GetByUser(user *models.User, db *gorm.DB) []models.UserParameter
	GetByName(parameterName string, db *gorm.DB) []models.UserParameter
	GetByID(userID uint64, parameterName string, db *gorm.DB) models.UserParameter
}
