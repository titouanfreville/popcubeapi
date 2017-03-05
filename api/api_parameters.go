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

// swagger:parameters newAvatar updateAvatar
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

// swagger:parameters newChannel updateChannel
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

// swagger:parameters newEmoji updateEmoji
type emojiObjectParam struct {
	//Link of the emoji in server.
	// in:body
	Emoji models.Emoji `json:"emoji"`
}

// <><><><><> <><><><><> <><><><><> <><><><><> //
