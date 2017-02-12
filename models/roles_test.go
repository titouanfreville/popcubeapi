package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestRolesModel(t *testing.T) {
	Convey("Testing IsValid function", t, func() {
		Convey("Given a correct roles. Should be validated", func() {
			role := Role{
				RoleName:      "testrole",
				CanUsePrivate: true,
				CanModerate:   false,
				CanArchive:    true,
				CanInvite:     false,
				CanManage:     false,
				CanManageUser: true,
			}
			So(role.IsValid(), ShouldBeNil)
			So(role.IsValid(), ShouldNotResemble, u.NewLocAppError("Role.IsValid", "model.role.rolename.app_error", nil, ""))
		})

		Convey("Given incorrect roles. Should be refused", func() {
			role := Role{
				RoleName:      "testRole",
				CanUsePrivate: true,
				CanModerate:   false,
				CanArchive:    true,
				CanInvite:     false,
				CanManage:     false,
				CanManageUser: true,
			}
			Convey("If rolename is not a lower case char, it should be refused", func() {
				So(role.IsValid(), ShouldResemble, u.NewLocAppError("Role.IsValid", "model.role.rolename.app_error", nil, ""))
				role.RoleName = "+alpha"
				So(role.IsValid(), ShouldResemble, u.NewLocAppError("Role.IsValid", "model.role.rolename.app_error", nil, ""))
				role.RoleName = "alpha-numerique"
				So(role.IsValid(), ShouldResemble, u.NewLocAppError("Role.IsValid", "model.role.rolename.app_error", nil, ""))
			})
		})
	})

	Convey("Basics roles must not be valid roles", t, func() {
		for _, role := range BasicsRoles {
			So(role.IsValid(), ShouldResemble, u.NewLocAppError("Role.IsValid", "model.role.rolename.app_error", nil, ""))
		}
	})

	Convey("Testing json VS role transformations", t, func() {
		Convey("Given an role", func() {
			Convey("Transforming it in JSON then back to EMOJI should provide similar objects", func() {
				json := Admin.ToJSON()
				newRole := RoleFromJSON(strings.NewReader(json))
				So(newRole, ShouldResemble, &Admin)
			})
		})

		Convey("Given an role list", func() {
			Convey("Transfoming it in JSON then back to EMOJI LIST shoud give ressembling objects", func() {
				json := RoleListToJSON(BasicsRoles)
				newRoleList := RoleListFromJSON(strings.NewReader(json))
				So(newRoleList, ShouldResemble, BasicsRoles)
			})

		})
	})

	Convey("Testing IsValidRoleName", t, func() {
		Convey("Given a correct role name", func() {
			Convey("It should be validate", func() {
				So(IsValidRoleName("u"), ShouldBeTrue)
				So(IsValidRoleName("another"), ShouldBeTrue)
				So(IsValidRoleName("world"), ShouldBeTrue)
				So(IsValidRoleName("xdealdex"), ShouldBeTrue)
			})
		})

		Convey("Given an incorrect role name", func() {
			Convey("Contain CAPS should be refused", func() {
				So(IsValidRoleName("U"), ShouldBeFalse)
				So(IsValidRoleName("anoTher"), ShouldBeFalse)
				So(IsValidRoleName("worlD"), ShouldBeFalse)
				So(IsValidRoleName("xDeAldEx"), ShouldBeFalse)
			})
			Convey("EMPTY or too long be refused", func() {
				So(IsValidRoleName(""), ShouldBeFalse)
				So(IsValidRoleName("thismustbeaverylongnamecontainingonlylowercasealphabeticalcharacterstobesurelengthistoomuch"), ShouldBeFalse)
			})
			Convey("Containing non alphabetics caracters", func() {
				So(IsValidRoleName("random2"), ShouldBeFalse)
				So(IsValidRoleName("random*"), ShouldBeFalse)
				So(IsValidRoleName("random?"), ShouldBeFalse)
				So(IsValidRoleName("random/"), ShouldBeFalse)
			})
		})
	})
}
