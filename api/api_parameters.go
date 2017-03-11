package api

import "github.com/titouanfreville/popcubeapi/models"

// <><><><><> AVATAR PARAMETERS <><><><><> //

// swagger:parameters getAvatarFromLink
type avatarLinkParam struct {
	//Link of the avatar in server.
	// in:path
	AvatarLink string `json:"avatarLink"`
}

// swagger:parameters getAvatarFromName
type avatarNameParam struct {
	//Link of the avatar in server.
	// in:path
	AvatarName string `json:"avatarName"`
}

// swagger:parameters updateAvatar deleteAvatar
type avatarIDParam struct {
	//Link of the avatar in server.
	// in:path
	AvatarID int `json:"avatarID"`
}

// swagger:parameters newAvatar newAvatar1 updateAvatar
type avatarObjectParam struct {
	//Link of the avatar in server.
	// in:body
	Avatar models.Avatar `json:"avatar"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> CHANNEL PARAMETERS <><><><><> //

// swagger:parameters getChannelFromType
type channelTypeParam struct {
	//Link of the channel in server.
	// in:path
	ChannelType string `json:"channelType"`
}

// swagger:parameters getChannelFromName
type channelNameParam struct {
	//Link of the channel in server.
	// in:path
	ChannelName string `json:"channelName"`
}

// swagger:parameters updateChannel deleteChannel
type channelIDParam struct {
	//Link of the channel in server.
	// in:path
	ChannelID int `json:"channelID"`
}

// swagger:parameters newChannel newChannel1 updateChannel getMemberFromChannel getMessageFromChannel
type channelObjectParam struct {
	//Link of the channel in server.
	// in:body
	Channel models.Channel `json:"channel"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> EMOJI PARAMETERS <><><><><> //

// swagger:parameters getEmojiFromLink
type emojiLinkParam struct {
	//Link of the emoji in server.
	// in:path
	EmojiLink string `json:"emojiLink"`
}

// swagger:parameters getEmojiFromName
type emojiNameParam struct {
	//Link of the emoji in server.
	// in:path
	EmojiName string `json:"emojiName"`
}

// swagger:parameters getEmojiFromShortcut
type emojiShortcutParam struct {
	//Link of the emoji in server.
	// in:path
	EmojiShortcut string `json:"emojiShortcut"`
}

// swagger:parameters updateEmoji deleteEmoji
type emojiIDParam struct {
	//Link of the emoji in server.
	// in:path
	EmojiID int `json:"emojiID"`
}

// swagger:parameters newEmoji newEmoji1 updateEmoji
type emojiObjectParam struct {
	//Link of the emoji in server.
	// in:body
	Emoji models.Emoji `json:"emoji"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> FOLDER PARAMETERS <><><><><> //

// swagger:parameters getFolderFromLink
type folderLinkParam struct {
	//Link of the folder in server.
	// in:path
	FolderLink string `json:"folderLink"`
}

// swagger:parameters getFolderFromName
type folderNameParam struct {
	//Link of the folder in server.
	// in:path
	FolderName string `json:"folderName"`
}

// swagger:parameters getFolderFromType
type folderTypeParam struct {
	//Link of the folder in server.
	// in:path
	FolderType string `json:"folderType"`
}

// swagger:parameters updateFolder deleteFolder
type folderIDParam struct {
	//Link of the folder in server.
	// in:path
	FolderID int `json:"folderID"`
}

// swagger:parameters newFolder newFolder1 updateFolder
type folderObjectParam struct {
	//Link of the folder in server.
	// in:body
	Folder models.Folder `json:"folder"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> MEMBER PARAMETERS <><><><><> //

// swagger:parameters updateMember
type memberIDParam struct {
	//Link of the member in server.
	// in:path
	MemberID int `json:"memberID"`
}

// swagger:parameters newMember newMember1 updateOrgansisation
type memberObjectParam struct {
	// Member object
	// in:body
	Member models.Member `json:"member"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> MESSAGE PARAMETERS <><><><><> //

// swagger:parameters getMessageFromDate
type messageDate struct {
	// Date of the message
	// in:path
	MessageDate int64 `json:"messageDate"`
}

// swagger:parameters updateMessage deleteMessage
type messageIDParam struct {
	//Link of the message in server.
	// in:path
	MessageID int `json:"messageID"`
}

// swagger:parameters newMessage newMessage1 updateMessage getFolderFromMessage
type messageObjectParam struct {
	//Link of the message in server.
	// in:body
	Message models.Message `json:"message"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> ORGANISATION PARAMETERS <><><><><> //

// swagger:parameters updateOrganisation
type organisationIDParam struct {
	//Link of the organisation in server.
	// in:path
	OrganisationID int `json:"organisationID"`
}

// swagger:parameters newOrganisation newOrganisation1 updateOrgansisation
type organisationObjectParam struct {
	// Organisation object
	// in:body
	Organisation models.Organisation `json:"organisation"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> PARMETER PARAMETERS <><><><><> //

// swagger:parameters updateParameter
type parameterIDParam struct {
	//Link of the parameter in server.
	// in:path
	ParameterID int `json:"parameterID"`
}

// swagger:parameters newParameter newParameter1 updateOrgansisation
type parameterObjectParam struct {
	// Parameter object
	// in:body
	Parameter models.Parameter `json:"parameter"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> ROLE PARAMETERS <><><><><> //

// swagger:parameters getRoleFromName
type roleNameParam struct {
	//Link of the role in server.
	// in:path
	RoleName string `json:"roleName"`
}

// swagger:parameters updateRole deleteRole
type roleIDParam struct {
	//Link of the role in server.
	// in:path
	RoleID int `json:"roleID"`
}

// swagger:parameters newRole newRole1 updateRole getUserFromRole getMemberFromRole
type roleObjectParam struct {
	// Role object
	// in:body
	Role models.Role `json:"role"`
}

// rightsParameterModel is the object you can pass to get roles from rights.
//
// swagger:model rightsParameterModel
type rightsParameterModel struct {
	// User can use private channel
	CanUsePrivate bool `gorm:"column:canUsePrivate;not null" json:"can_use_private,omitempty"`
	// User can moderate channels
	CanModerate bool `gorm:"column:canModerate;not null" json:"can_moderate,omitempty"`
	// User can archive channels
	CanArchive bool `gorm:"column:canArchive;not null" json:"can_archive,omitempty"`
	// User can invite others to private channel or organisation
	CanInvite bool `gorm:"column:canInvite;not null" json:"can_invite,omitempty"`
	// User can manage organisation/channel parameters and data
	CanManage bool `gorm:"column:canManage;not null" json:"can_manage,omitempty"`
	// User can manage other organisation/channel user
	CanManageUser bool `gorm:"column:canManageUser;not null" json:"can_manage_user,omitempty"`
}

// swagger:parameters getRoleFromRights
type rightsObjectParam struct {
	// Right of the role we search
	// in:body
	Rights rightsParameterModel `json:"rights"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> USER PARAMETERS <><><><><> //

// swagger:parameters getUserFromName
type userNameParam struct {
	//Link of the user in server.
	// in:path
	UserName string `json:"userName"`
}

// swagger:parameters getUserFromNickName
type nickNameParam struct {
	//Link of the user in server.
	// in:path
	NickName string `json:"nickName"`
}

// swagger:parameters getUserFromFirstName
type firstNameParam struct {
	//Link of the user in server.
	// in:path
	FirstName string `json:"firstName"`
}

// swagger:parameters getUserFromLastName
type lastNameeParam struct {
	//Link of the user in server.
	// in:path
	LastName string `json:"lastName"`
}

// swagger:parameters updateUser deleteUser
type userIDParam struct {
	//Link of the user in server.
	// in:path
	UserID int `json:"userID"`
}

// swagger:parameters newUser newUser1 updateUser getMemberFromUser getMessageFromUser
type userObjectParam struct {
	//Link of the user in server.
	// in:body
	User models.User `json:"user"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> LOGIN PARAMETERS <><><><><> //

// loginParameterModel is the object you have to pass to login
//
// swagger:model loginParameterModel
type loginParameterModel struct {
	// user name
	// required: true
	Login string `json:"login"`
	// user password hashed
	// required: true
	Password string `json:"password"`
}

// swagger:parameters login
type loginParam struct {
	// Login informations
	// in:body
	Login loginParameterModel `json:"login"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
// <><><><><> TOKENS PARAMETERS <><><><><> //

// swagger:parameters getAllAvatar getAllAvatar1 newAvatar newAvatar1 getAvatarFromLink getAvatarFromName updateAvatar deleteAvatar getAllChannel getAllChannel1 newChannel newChannel1 getPublicChannel getPrivateChannel getChannelFromType getChannelFromName updateChannel deleteChannel getAllEmoji getAllEmoji1 newEmoji1 newEmoji getEmojiFromLink getEmojiFromName getEmojiFromShortcut updateEmoji deleteEmoji getAllFolder getAllFolder1 newFolder newFolder1 getFolderFromMessage getFolderFromName getFolderFromLink getFolderFromType updateFolder deleteFolder getAllMember getAllMember1 newMember newMember1 getMemberFromChannel getMemberFromUser getMemberFromRole updateMember deleteMember getAllMessage getAllMessage1 newMessage newMessage1 getMessageFromChannel getMessageFromUser getMessageFromDate updateMessage deleteMessage getAllOrganisation getAllOrganisation1 newOrganisation newOrganisation1 updateOrganisation getAllParameter getAllParameter1 newParameter newParameter1 updateParameter getAllRole getAllRole1 newRole newRole1 getRoleFromRights getRoleFromName updateRole deleteRole getAllUser getAllUser1 newUser newUser1 inviteUser getDeletedUser getUserFromRole getOrderedByDate getUserFromName getUserFromNickName getUserFromFirstName getUserFromLastName updateUser deleteUser
type userTokenParameter struct {
	// User token you got from login call. Pass it as Authentication: bearer {{token}} in the header
	// required: true
	// in:header
	Token string `json:"token"`
}

// swagger:parameters newInvitedUser
type inviteToken struct {
	// Invite token you got from user/invite call. Pass it as Authentication: bearer {{token}} in the header
	// in:body
	Token string `json:"token"`
}

// swagger:parameters newPublicUser
type initToken struct {
	// Init token you got when creating new organisation. Pass it as Authentication: bearer {{token}} in the header
	// in:body
	Token string `json:"token"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
