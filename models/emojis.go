package models

import (
	"encoding/json"
	"io"

	u "github.com/titouanfreville/popcubeapi/utils"
)

// Emoji Type descibe the Emoji table for Popcube DB
type Emoji struct {
	IDEmoji  uint64 `gorm:"primary_key;column:idEmoji;AUTO_INCREMENT" json:"-"`
	Name     string `gorm:"column:name;not null;unique" json:"name"`
	Shortcut string `gorm:"column:shortcut;not null;unique" json:"shortcut"`
	Link     string `gorm:"column:link;not null;unique" json:"link"`
}

// IsValid is used to check validity of Emoji objects
func (emoji *Emoji) IsValid() *u.AppError {
	if len(emoji.Name) == 0 || len(emoji.Name) > 64 {
		return u.NewLocAppError("Emoji.IsValid", "model.emoji.name.app_error", nil, "")
	}

	if len(emoji.Shortcut) == 0 || len(emoji.Shortcut) > 20 {
		return u.NewLocAppError("Emoji.IsValid", "model.emoji.shortcut.app_error", nil, "")
	}

	if len(emoji.Link) == 0 {
		return u.NewLocAppError("Emoji.IsValid", "model.emoji.link.app_error", nil, "")
	}

	return nil
}

// ToJSON transfoorm an Emoji into JSON
func (emoji *Emoji) ToJSON() string {
	b, err := json.Marshal(emoji)
	if err != nil {
		return ""
	}
	return string(b)
}

// EmojiFromJSON Try to parse a json object as emoji
func EmojiFromJSON(data io.Reader) *Emoji {
	decoder := json.NewDecoder(data)
	var emoji Emoji
	err := decoder.Decode(&emoji)
	if err == nil {
		return &emoji
	}
	return nil
}

// EmojiListToJSON Convert an emoji list into a json array
func EmojiListToJSON(emojiList []*Emoji) string {
	b, err := json.Marshal(emojiList)
	if err != nil {
		return ""
	}
	return string(b)
}

// EmojiListFromJSON Try converting a json array into emoji list
func EmojiListFromJSON(data io.Reader) []*Emoji {
	decoder := json.NewDecoder(data)
	var emojiList []*Emoji
	err := decoder.Decode(&emojiList)
	if err == nil {
		return emojiList
	}
	return nil
}
