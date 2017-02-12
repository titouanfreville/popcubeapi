package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestAvatarsModel(t *testing.T) {
	Convey("Testing IsValid function", t, func() {
		Convey("Given a correct avatars. Should be validated", func() {
			avatar := Avatar{
				Name: "Troll Face",
				Link: "avatars/trollface.svg",
			}
			So(avatar.IsValid(), ShouldBeNil)
			So(avatar.IsValid(), ShouldNotResemble, u.NewLocAppError("Avatar.IsValid", "model.avatar.name.app_error", nil, ""))
			So(avatar.IsValid(), ShouldNotResemble, u.NewLocAppError("Avatar.IsValid", "model.avatar.link.app_error", nil, ""))
		})

		Convey("Given incorrect avatars. Should be refused", func() {
			avatar := Avatar{
				Name: "Troll Face",
				Link: "avatars/trollface.svg",
			}

			avatar.Name = ""

			Convey("Too long or empty Name should return name error", func() {
				So(avatar.IsValid(), ShouldResemble, u.NewLocAppError("Avatar.IsValid", "model.avatar.name.app_error", nil, ""))
				avatar.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				So(avatar.IsValid(), ShouldResemble, u.NewLocAppError("Avatar.IsValid", "model.avatar.name.app_error", nil, ""))
			})

			avatar.Name = "Correct Name"
			avatar.Link = ""

			Convey("Empty link should result in link error", func() {
				So(avatar.IsValid(), ShouldResemble, u.NewLocAppError("Avatar.IsValid", "model.avatar.link.app_error", nil, ""))
			})
		})
	})

	Convey("Testing json VS avatar transformations", t, func() {
		Convey("Given an avatar", func() {
			avatar := Avatar{
				Name: "Troll Face",
				Link: "avatars/trollface.svg",
			}
			Convey("Transforming it in JSON then back to EMOJI should provide similar objects", func() {
				json := avatar.ToJSON()
				newAvatar := AvatarFromJSON(strings.NewReader(json))
				So(newAvatar, ShouldResemble, &avatar)
			})
		})

		Convey("Given an avatar list", func() {
			avatar1 := Avatar{
				Name: "Troll Face",
				Link: "avatars/trollface.svg",
			}
			avatar2 := Avatar{
				Name: "Joy Face",
				Link: "avatars/joyface.svg",
			}
			avatar3 := Avatar{
				Name: "Face Palm",
				Link: "avatars/facepalm.svg",
			}
			avatarList := []*Avatar{&avatar1, &avatar2, &avatar3}

			Convey("Transfoming it in JSON then back to EMOJI LIST shoud give ressembling objects", func() {
				json := AvatarListToJSON(avatarList)
				newAvatarList := AvatarListFromJSON(strings.NewReader(json))
				So(newAvatarList, ShouldResemble, avatarList)
			})

		})
	})

}
