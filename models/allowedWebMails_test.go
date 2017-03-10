package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestAllowedWebMailssModel(t *testing.T) {
	Convey("Testing IsValid function", t, func() {
		Convey("Given a correct allowedWebMailss. Should be validated", func() {
			allowedWebMails := AllowedWebMails{
				Domain:        "popcube.xyz",
				Provider:      "PopCube",
				DefaultRights: "Master",
			}
			So(allowedWebMails.IsValid(), ShouldBeNil)
			So(allowedWebMails.IsValid(), ShouldNotResemble, u.NewLocAppError("allowedWebMails.isValid", "domain undefined", nil, ""))
		})

		Convey("Given incorrect allowedWebMailss. Should be refused", func() {
			allowedWebMails := AllowedWebMails{
				Provider:      "PopCube",
				DefaultRights: "Master",
			}

			Convey("Too long or empty Name should return name error", func() {
				So(allowedWebMails.IsValid(), ShouldNotBeNil)
				So(allowedWebMails.IsValid(), ShouldResemble, u.NewLocAppError("allowedWebMails.isValid", "domain undefined", nil, ""))
			})
		})
	})
}
