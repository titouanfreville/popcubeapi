package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestMemberModel(t *testing.T) {
	userTest := User{
		WebID:              NewID(),
		LastUpdate:          10,
		Deleted:            true,
		Username:           "l",
		Password:           "test",
		Email:              "test@popcube.fr",
		EmailVerified:      true,
		NickName:           "NickName",
		FirstName:          "Test",
		LastName:           "L",
		Role:               Owner,
		LastPasswordUpdate: 20,
		FailedAttempts:     1,
		Locale:             "vi",
	}

	channelTest := Channel{
		WebID:       NewID(),
		ChannelName: "electra",
		LastUpdate:   GetMillis(),
		Type:        "audio",
		Private:     false,
		Description: "Testing channel description :O",
		Subject:     "Sujet",
		Avatar:      "jesuiscool.svg",
	}

	Convey("Testing IsValid function", t, func() {
		Convey("Given a correct member. Should be validated", func() {
			member := Member{
				User:    userTest,
				Channel: channelTest,
			}
			So(member.IsValid(), ShouldBeNil)
			So(member.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
			So(member.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
		})

		Convey("Given incorrect member. Should be refused", func() {
			empty := Member{}
			member := Member{
				User:    userTest,
				Channel: channelTest,
			}
			member.User = User{}
			Convey("Empty member or member without User should return User error", func() {
				So(member.IsValid(), ShouldResemble, u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
				So(member.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
				So(empty.IsValid(), ShouldResemble, u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
				So(empty.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
			})

			member.User = userTest
			member.Channel = Channel{}
			Convey("Empty link should result in link error", func() {
				So(member.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
				So(member.IsValid(), ShouldResemble, u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
			})
		})
	})
}
