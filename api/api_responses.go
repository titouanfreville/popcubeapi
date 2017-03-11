package api

import (
	"github.com/titouanfreville/popcubeapi/models"
	"github.com/titouanfreville/popcubeapi/utils"
)

// Success -------------------------------------------

// generalOk default object style to return if correct
//
// swagger:response defaultOk
type generalOk struct {
	// What you want to say
	// in:body
	Message string `json:"message,omitempty"`
}

// loginOk when login correctly proceed, return the user and an auth token.
//
// swagger:response loginOk
type loginOk struct {
	// in:body
	User models.User `json:"user,omitempty"`
	// in:body
	Token string `json:"token,omitempty"`
}

// inviteOk when invite correctly proceed, return the invite token information and the JWT token.
//
// swagger:response loginOk
type inviteOk struct {
	// in:body
	Email string `json:"email,omitempty"`
	// in:body
	Organisation string `json:"organisation,omitempty"`
	// in:body
	Token string `json:"token,omitempty"`
}

// initOk when init correctly proceed, return the organisation object and its owner.
//
// swagger:response loginOk
type initOk struct {
	// in:body
	Organisation models.Organisation `json:"organisation,omitempty"`
	// in:body
	Owner models.User `json:"user,omitempty"`
}

// ---------------------------------------------------
// Errors --------------------------------------------

// genericError general error when unexpected errors occured
//
// swagger:response genericError
type genericError struct {
	// in:body
	Error utils.AppError
}

// wrongEntityError is an error object to inform that the provided object was not correctly formated
//
// swagger:response wrongEntity
type wrongEntityError struct {
	// in:body
	Error utils.AppError
}

// databaseError is an error object to tell what is happening when we encounter issue with database
//
// swagger:response databaseError
type databaseError struct {
	// in:body
	Error utils.AppError
}

// incorrectIds return error login message
//
// swagger:response incorrectIds
type incorrectIds struct {
	// in:body
	Error utils.AppError
}

// ---------------------------------------------------
// Unknow --------------------------------------------

// deleteMessageModel is an object to confirm correct deletion of an item.
//
// swagger:model deleteMessageModel
type deleteMessageModel struct {
	// Status
	Status int `json:"status"`
	// Correctly removed ?
	Success bool `json:"success"`
	// More information about why is it or isn't it removed
	Message string `json:"message,omitempty"`
	// The object we where trying to remove
	Object interface{} `json:"removed_object, omitempty"`
}

// deleteMessage return object to confirm correct deletion of an item.
//
// swagger:model deleteMessage
type deleteMessage struct {
	// in:body
	Message deleteMessageModel
}

// ---------------------------------------------------
// Generators ----------------------------------------

func newGeneralOk(message string) generalOk {
	return generalOk{
		Message: message,
	}
}

func newDeleteMessage(succes bool, message string) deleteMessageModel {
	return deleteMessageModel{
		Status:  200,
		Message: message,
		Success: succes,
	}
}

// ---------------------------------------------------

// <><><><><> AVATAR RESPONSES <><><><><> //

// avatarSlice Array of avatars
//
// swagger:response avatarArraySuccess
type avatarArraySuccess struct {
	// in:body
	Avatars []models.Avatar
}

// avatarObjectSuccess list of avatars
//
// swagger:response avatarObjectSuccess
type avatarObjectSuccess struct {
	// in:body
	// List of avatars returned
	Avatar models.Avatar `json:"avatar"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //

// <><><><><> CHANNEL RESPONSES <><><><><> //

// channelSlice Array of channels
//
// swagger:response channelArraySuccess
type channelArraySuccess struct {
	// in:body
	Channels []models.Channel
}

// channelObjectSuccess list of channels
//
// swagger:response channelObjectSuccess
type channelObjectSuccess struct {
	// in:body
	// List of channels returned
	Channel models.Channel `json:"channel"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //

// <><><><><> EMOJI RESPONSES <><><><><> //

// emojiSlice Array of emojis
//
// swagger:response emojiArraySuccess
type emojiArraySuccess struct {
	// in:body
	Emojis []models.Emoji
}

// emojiObjectSuccess list of emojis
//
// swagger:response emojiObjectSuccess
type emojiObjectSuccess struct {
	// in:body
	// List of emojis returned
	Emoji models.Emoji `json:"emoji"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //

// <><><><><> FOLDER RESPONSES <><><><><> //

// folderSlice Array of folders
//
// swagger:response folderArraySuccess
type folderArraySuccess struct {
	// in:body
	Folders []models.Folder
}

// folderObjectSuccess list of folders
//
// swagger:response folderObjectSuccess
type folderObjectSuccess struct {
	// in:body
	// List of folders returned
	Folder models.Folder `json:"folder"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //

// <><><><><> MEMBER RESPONSES <><><><><> //

// memberSlice Array of members
//
// swagger:response memberArraySuccess
type memberArraySuccess struct {
	// in:body
	Members []models.Member
}

// memberObjectSuccess list of members
//
// swagger:response memberObjectSuccess
type memberObjectSuccess struct {
	// in:body
	// List of members returned
	Member models.Member `json:"member"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> MESSAGE RESPONSES <><><><><> //

// messageSlice Array of messages
//
// swagger:response messageArraySuccess
type messageArraySuccess struct {
	// in:body
	Messages []models.Message
}

// messageObjectSuccess list of messages
//
// swagger:response messageObjectSuccess
type messageObjectSuccess struct {
	// in:body
	// List of messages returned
	Message models.Message `json:"message"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> ORGANISATION RESPONSES <><><><><> //

// organisationObjectSuccess list of organisations
//
// swagger:response organisationObjectSuccess
type organisationObjectSuccess struct {
	// in:body
	// List of organisations returned
	Organisation models.Organisation `json:"organisation"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //

// <><><><><> PARAMETERS RESPONSES <><><><><> //

// parameterObjectSuccess list of parameters
//
// swagger:response parameterObjectSuccess
type parameterObjectSuccess struct {
	// in:body
	// List of parameters returned
	Parameter models.Parameter `json:"parameter"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> ROLE RESPONSES <><><><><> //

// roleSlice Array of roles
//
// swagger:response roleArraySuccess
type roleArraySuccess struct {
	// in:body
	Roles []models.Role
}

// roleObjectSuccess list of roles
//
// swagger:response roleObjectSuccess
type roleObjectSuccess struct {
	// in:body
	// List of roles returned
	Role models.Role `json:"role"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> USER RESPONSES <><><><><> //

// userSlice Array of users
//
// swagger:response userArraySuccess
type userArraySuccess struct {
	// in:body
	Users []models.User
}

// userObjectSuccess list of users
//
// swagger:response userObjectSuccess
type userObjectSuccess struct {
	// in:body
	// List of users returned
	User models.User `json:"user"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
