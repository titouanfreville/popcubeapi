package models

import (
	"encoding/json"
	"io"

	u "github.com/titouanfreville/popcubeapi/utils"
)

// Avatar type is a DB model for avatar storage
type Avatar struct {
	IDAvatar uint64 `gorm:"primary_key;column:idAvatar;AUTO_INCREMENT" json:"id,omitempty"`
	Name     string `gorm:"column:name;not null;unique" json:"name,omitempty"`
	Link     string `gorm:"column:link;not null;unique" json:"link,omitempty"`
}

// IsValid check the validity of on Avatar object before saving it to DB in update or creation process
func (avatar *Avatar) IsValid() *u.AppError {
	if len(avatar.Name) == 0 || len(avatar.Name) > 64 {
		return u.NewLocAppError("Avatar.IsValid", "model.avatar.name.app_error", nil, "")
	}

	if len(avatar.Link) == 0 {
		return u.NewLocAppError("Avatar.IsValid", "model.avatar.link.app_error", nil, "")
	}

	return nil
}

// ToJSON function take an avatar and tranform it into Json object
func (avatar *Avatar) ToJSON() string {
	b, err := json.Marshal(avatar)
	if err != nil {
		return ""
	}
	return string(b)
}

// AvatarFromJSON Take a json and try to transform it into an avatar
func AvatarFromJSON(data io.Reader) *Avatar {
	decoder := json.NewDecoder(data)
	var avatar Avatar
	err := decoder.Decode(&avatar)
	if err == nil {
		return &avatar
	}
	return nil
}

// AvatarListToJSON Take an avatar list and transform it into json object
func AvatarListToJSON(avatarList []*Avatar) string {
	b, err := json.Marshal(avatarList)
	if err != nil {
		return ""
	}
	return string(b)
}

// AvatarListFromJSON Try to parse a json object as an avatar list
func AvatarListFromJSON(data io.Reader) []*Avatar {
	decoder := json.NewDecoder(data)
	var avatarList []*Avatar
	err := decoder.Decode(&avatarList)
	if err == nil {
		return avatarList
	}
	return nil
}
