package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestEmojisModel(t *testing.T) {
	Convey("Testing IsValid function", t, func() {
		Convey("Given a correct emojis. Should be validated", func() {
			emoji := Emoji{
				Name:     "Troll Face",
				Shortcut: ":troll-face:",
				Link:     "emojis/trollface.svg",
			}
			So(emoji.IsValid(), ShouldBeNil)
			So(emoji.IsValid(), ShouldNotResemble, u.NewLocAppError("Emoji.IsValid", "model.emoji.name.app_error", nil, ""))
			So(emoji.IsValid(), ShouldNotResemble, u.NewLocAppError("Emoji.IsValid", "model.emoji.shortcut.app_error", nil, ""))
			So(emoji.IsValid(), ShouldNotResemble, u.NewLocAppError("Emoji.IsValid", "model.emoji.link.app_error", nil, ""))
		})

		Convey("Given incorrect emojis. Should be refused", func() {
			emoji := Emoji{
				Name:     "Troll Face",
				Shortcut: ":this-is-a-tool-long-shortcut:",
				Link:     "emojis/trollface.svg",
			}

			Convey("Too long shortcut or empty shorctcut should return Shortcut error", func() {
				So(emoji.IsValid(), ShouldResemble, u.NewLocAppError("Emoji.IsValid", "model.emoji.shortcut.app_error", nil, ""))
				emoji.Shortcut = ""
				So(emoji.IsValid(), ShouldResemble, u.NewLocAppError("Emoji.IsValid", "model.emoji.shortcut.app_error", nil, ""))
			})
			emoji.Shortcut = ":goodone:"
			emoji.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				So(emoji.IsValid(), ShouldResemble, u.NewLocAppError("Emoji.IsValid", "model.emoji.name.app_error", nil, ""))
				emoji.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				So(emoji.IsValid(), ShouldResemble, u.NewLocAppError("Emoji.IsValid", "model.emoji.name.app_error", nil, ""))
			})
			emoji.Name = "Correct Name"
			emoji.Link = ""
			Convey("Empty link should result in link error", func() {
				So(emoji.IsValid(), ShouldResemble, u.NewLocAppError("Emoji.IsValid", "model.emoji.link.app_error", nil, ""))
			})
		})
	})

	Convey("Testing json VS emoji transformations", t, func() {
		Convey("Given an emoji", func() {
			emoji := Emoji{
				Name:     "Troll Face",
				Shortcut: ":troll-face:",
				Link:     "emojis/trollface.svg",
			}
			Convey("Transforming it in JSON then back to EMOJI should provide similar objects", func() {
				json := emoji.ToJSON()
				newEmoji := EmojiFromJSON(strings.NewReader(json))
				So(newEmoji, ShouldResemble, &emoji)
			})
		})

		Convey("Given an emoji list", func() {
			emoji1 := Emoji{
				Name:     "Troll Face",
				Shortcut: ":troll:",
				Link:     "emojis/trollface.svg",
			}
			emoji2 := Emoji{
				Name:     "Joy Face",
				Shortcut: ":)",
				Link:     "emojis/joyface.svg",
			}
			emoji3 := Emoji{
				Name:     "Face Palm",
				Shortcut: ":facepalm:",
				Link:     "emojis/facepalm.svg",
			}
			emojiList := []*Emoji{&emoji1, &emoji2, &emoji3}

			Convey("Transfoming it in JSON then back to EMOJI LIST shoud give ressembling objects", func() {
				json := EmojiListToJSON(emojiList)
				newEmojiList := EmojiListFromJSON(strings.NewReader(json))
				So(newEmojiList, ShouldResemble, emojiList)
			})
		})
	})

}
