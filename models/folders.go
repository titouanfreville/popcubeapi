package models

import (
	"encoding/json"
	"io"

	u "github.com/titouanfreville/popcubeapi/utils"
)

// Folder Type descibe the Folder table for Popcube DB. It is linked to Message table
type Folder struct {
	IDFolder uint64  `gorm:"primary_key;column:idFolder;AUTO_INCREMENT" json:"-"`
	Link     string  `gorm:"column:link;not null;unique" json:"link"`
	Name     string  `gorm:"column:name;not null;unique" json:"name"`
	Type     string  `gorm:"column:type;not null;" json:"type"`
	Message  Message `gorm:"column:message; not null;ForeignKey:IDMessage;" json:"-"`
}

// IsValid is used to check validity of Folder objects
func (folder *Folder) IsValid() *u.AppError {
	if len(folder.Name) == 0 {
		return u.NewLocAppError("Folder.IsValid", "model.folder.name.app_error", nil, "")
	}

	if len(folder.Link) == 0 {
		return u.NewLocAppError("Folder.IsValid", "model.folder.link.app_error", nil, "")
	}
	if len(folder.Type) == 0 {
		return u.NewLocAppError("Folder.IsValid", "model.folder.type.app_error", nil, "")
	}
	if folder.Message == (Message{}) {
		return u.NewLocAppError("Folder.IsValid", "model.folder.message.app_error", nil, "")
	}
	return nil
}

// ToJSON transfoorm an Folder into JSON
func (folder *Folder) ToJSON() string {
	b, err := json.Marshal(folder)
	if err != nil {
		return ""
	}
	return string(b)
}

// FolderFromJSON Try to parse a json object as emoji
func FolderFromJSON(data io.Reader) *Folder {
	decoder := json.NewDecoder(data)
	var folder Folder
	err := decoder.Decode(&folder)
	if err == nil {
		return &folder
	}
	return nil
}

// FolderListToJSON Convert an emoji list into a json array
func FolderListToJSON(folderList []*Folder) string {
	b, err := json.Marshal(folderList)
	if err != nil {
		return ""
	}
	return string(b)
}

// FolderListFromJSON Try converting a json array into emoji list
func FolderListFromJSON(data io.Reader) []*Folder {
	decoder := json.NewDecoder(data)
	var folderList []*Folder
	err := decoder.Decode(&folderList)
	if err == nil {
		return folderList
	}
	return nil
}
