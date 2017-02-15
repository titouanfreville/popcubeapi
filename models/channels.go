package models

import (
	"encoding/json"
	"io"
	"strings"
	"unicode/utf8"

	u "github.com/titouanfreville/popcubeapi/utils"
)

const (
	defaultChannel             = "general"
	channelDislayNameMaxRunes  = 64
	channelNameMaxLength       = 64
	channelDescriptionMaxRunes = 1024
	channelSubjectMaxRunes     = 250
)

var (
	// ChannelAvailableTypes Used to have knowsledge on type a channel can take
	ChannelAvailableTypes = []string{"direct", "text", "audio", "video"}
)

// Channel type is a model for DB Channel table
type Channel struct {
	IDChannel   uint64 `gorm:"primary_key;column:idChannel;AUTO_INCREMENT" json:"-"`
	WebID       string `gorm:"column:webId;not null;unique" json:"web_id"`
	ChannelName string `gorm:"column:channelName;not null;unique" json:"display_name"`
	Type        string `gorm:"column:type;not null" json:"type"`
	LastUpdate  int64  `gorm:"column:lastUpdate;not null;" json:"last_update"`
	Private     bool   `gorm:"column:private;not null" json:"private"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	Subject     string `gorm:"column:subject" json:"subject,omitempty"`
	Avatar      string `gorm:"column:avatar" json:"avatar,omitempty"`
}

// ToJSON Take a channel and convert it into json
func (channel *Channel) ToJSON() string {
	b, err := json.Marshal(channel)
	if err != nil {
		return ""
	}
	return string(b)
}

// ChannelFromJSON try to parse a json object as channel object
func ChannelFromJSON(data io.Reader) *Channel {
	decoder := json.NewDecoder(data)
	var channel Channel
	err := decoder.Decode(&channel)
	if err == nil {
		return &channel
	}
	return nil
}

// Etag is a small function used to create cache ID
func (channel *Channel) Etag() string {
	return Etag(channel.WebID, channel.LastUpdate)
}

// IsValid check the correctness of a channel object
func (channel *Channel) IsValid(isUpdate bool) *u.AppError {
	if !isUpdate {
		if len(channel.WebID) != 26 {
			return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.id.app_error", nil, "")
		}
	}

	if channel.LastUpdate == 0 {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.update_at.app_error", nil, "id="+channel.WebID)
	}

	if utf8.RuneCountInString(channel.ChannelName) > channelDislayNameMaxRunes || utf8.RuneCountInString(channel.ChannelName) == 0 {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.channel_name.app_error", nil, "id="+channel.WebID)
	}

	if !IsValidIDChannelentifier(channel.ChannelName) {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID)
	}

	if utf8.RuneCountInString(channel.Description) > channelDescriptionMaxRunes {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.description.app_error", nil, "id="+channel.WebID)
	}

	if utf8.RuneCountInString(channel.Subject) > channelSubjectMaxRunes {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.subject.app_error", nil, "id="+channel.WebID)
	}

	if !u.StringInArray(channel.Type, ChannelAvailableTypes) {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.type.app_error", nil, "id="+channel.WebID)
	}

	return nil
}

// PreSave Is used to add default values to channel before saving it in DB
func (channel *Channel) PreSave() {
	if channel.WebID == "" {
		channel.WebID = NewID()
	}

	channel.ChannelName = strings.ToLower(channel.ChannelName)

	channel.LastUpdate = GetMillis()

	if channel.Avatar == "" {
		channel.Avatar = "default_channel_avatar.svg"
	}

	if channel.Type == "" {
		channel.Type = "text"
	}

	if channel.Type == "direct" {
		channel.Private = true
	}
}

// PreUpdate Is used to add default values to channel before updating it in DB
func (channel *Channel) PreUpdate() {
	channel.ChannelName = strings.ToLower(channel.ChannelName)

	channel.LastUpdate = GetMillis()

	if channel.Type == "direct" {
		channel.Private = true
	}
}

// GetDMNameFromIDs Create Direct message name from 2 IDUsers
func GetDMNameFromIDs(IDUser1, IDUser2 string) string {
	if IDUser1 > IDUser2 {
		return IDUser2 + "__" + IDUser1
	}
	return IDUser1 + "__" + IDUser2

}
