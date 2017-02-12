// This file is used to test if user model is working correctly.
// A user is always linked to an role
// He has brsic channel to join
package datastores

import (
	. "github.com/titouanfreville/popcubeapi/models"
	"testing"
	u "github.com/titouanfreville/popcubeapi/utils"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRoleStore(t *testing.T) {
	ds := dbStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev")
	db := *ds.Db
	rsi := RoleStoreImpl{}
	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("roleStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("roleStoreImpl.Save", "save.transaction.create.already_exist", nil, "Role Name: classe")
		role := Role{
			RoleName:      "classe",
			CanUsePrivate: true,
			CanModerate:   false,
			CanArchive:    true,
			CanInvite:     false,
			CanManage:     false,
			CanManageUser: true,
		}
		Convey("Given a correct role.", func() {
			appError := rsi.Save(&role, ds)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := rsi.Save(&role, ds)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyExistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&role)
	})

	Convey("Testing update function", t, func() {
		dbError := u.NewLocAppError("roleStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("roleStoreImpl.Save", "save.transaction.create.already_exist", nil, "Role Name: Troll Face")
		role := Role{
			RoleName:      "classe",
			CanUsePrivate: true,
			CanModerate:   false,
			CanArchive:    true,
			CanInvite:     false,
			CanManage:     false,
			CanManageUser: true,
		}
		roleNew := Role{
			RoleName:      "classenew",
			CanUsePrivate: false,
			CanModerate:   true,
			CanArchive:    false,
			CanInvite:     true,
			CanManage:     true,
			CanManageUser: false,
		}

		appError := rsi.Save(&role, ds)
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		So(appError, ShouldNotResemble, alreadyExistError)

		Convey("Provided correct Role to modify should not return errors", func() {
			appError := rsi.Update(&role, &roleNew, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
			So(appError, ShouldNotResemble, alreadyExistError)
		})

		Convey("Provided wrong old Role to modify should result in old_role error", func() {
			role.RoleName = "testRole"
			Convey("If rolename is not a lower case char, it should be refused", func() {
				appError := rsi.Update(&role, &roleNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("roleStoreImpl.Update.roleOld.PreSave", "model.role.rolename.app_error", nil, ""))
				role.RoleName = "+alpha"
				appError = rsi.Update(&role, &roleNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("roleStoreImpl.Update.roleOld.PreSave", "model.role.rolename.app_error", nil, ""))
				role.RoleName = "alpha-numerique"
				appError = rsi.Update(&role, &roleNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("roleStoreImpl.Update.roleOld.PreSave", "model.role.rolename.app_error", nil, ""))
			})
		})

		Convey("Provided wrong new Role to modify should result in newRole error", func() {
			roleNew.RoleName = "testRole"
			appError := rsi.Update(&role, &roleNew, ds)
			So(appError, ShouldNotBeNil)
			So(appError, ShouldNotResemble, dbError)
			So(appError, ShouldNotResemble, alreadyExistError)
			So(appError, ShouldResemble, u.NewLocAppError("roleStoreImpl.Update.roleNew.PreSave", "model.role.rolename.app_error", nil, ""))
			roleNew.RoleName = "+alpha"
			appError = rsi.Update(&role, &roleNew, ds)
			So(appError, ShouldNotBeNil)
			So(appError, ShouldNotResemble, dbError)
			So(appError, ShouldNotResemble, alreadyExistError)
			So(appError, ShouldResemble, u.NewLocAppError("roleStoreImpl.Update.roleNew.PreSave", "model.role.rolename.app_error", nil, ""))
			roleNew.RoleName = "alpha-numerique"
			appError = rsi.Update(&role, &roleNew, ds)
			So(appError, ShouldNotBeNil)
			So(appError, ShouldNotResemble, dbError)
			So(appError, ShouldNotResemble, alreadyExistError)
			So(appError, ShouldResemble, u.NewLocAppError("roleStoreImpl.Update.roleNew.PreSave", "model.role.rolename.app_error", nil, ""))

		})
		db.Delete(&role)
		db.Delete(&roleNew)
	})

	Convey("Testing Getters", t, func() {

		rsi.GetByName("owner", ds)

		role0 := Role{
			RoleName:      "classe",
			CanUsePrivate: true,
			CanModerate:   false,
			CanArchive:    true,
			CanInvite:     false,
			CanManage:     false,
			CanManageUser: true,
		}
		role1 := Role{
			RoleName:      "deuce",
			CanUsePrivate: true,
			CanModerate:   false,
			CanArchive:    false,
			CanInvite:     false,
			CanManage:     true,
			CanManageUser: false,
		}
		role2 := Role{
			RoleName:      "fg",
			CanUsePrivate: true,
			CanModerate:   false,
			CanArchive:    false,
			CanInvite:     true,
			CanManage:     true,
			CanManageUser: true,
		}
		role1New := Role{
			RoleName:      "tierce",
			CanUsePrivate: false,
			CanModerate:   true,
			CanArchive:    true,
			CanInvite:     true,
			CanManage:     false,
			CanManageUser: true,
		}
		role3 := Role{
			RoleName:      "test",
			CanUsePrivate: true,
			CanModerate:   true,
			CanArchive:    false,
			CanInvite:     false,
			CanManage:     false,
			CanManageUser: true,
		}
		rolesCanPrivate := Role{CanUsePrivate: true}
		rolesCanPrivateNotArchive := Role{CanUsePrivate: true, CanArchive: false}

		rsi.Save(&role0, ds)
		rsi.Save(&role1, ds)
		rsi.Update(&role1, &role1New, ds)
		rsi.Save(&role2, ds)
		rsi.Save(&role3, ds)

		// Have to be after save so ID are up to date :O
		roleList := []Role{
			role0,
			role1,
			role2,
			role3,
		}

		canPrivateList := []Role{role0, role2, role3}
		emptyList := []Role{}

		Convey("We have to be able to find all roles in the db", func() {
			roles := rsi.GetAll(ds)
			So(roles, ShouldNotResemble, &emptyList)
			So(roles, ShouldResemble, &roleList)
		})

		Convey("We have to be able to find an role from is name", func() {
			role := rsi.GetByName(role0.RoleName, ds)
			So(role, ShouldNotResemble, &Role{})
			So(role, ShouldResemble, &role0)
			role = rsi.GetByName(role2.RoleName, ds)
			So(role, ShouldNotResemble, &Role{})
			So(role, ShouldResemble, &role2)
			role = rsi.GetByName(role3.RoleName, ds)
			So(role, ShouldNotResemble, &Role{})
			So(role, ShouldResemble, &role3)
			Convey("Should also work from updated value", func() {
				role = rsi.GetByName(role1New.RoleName, ds)
				So(role, ShouldNotResemble, &Role{})
				So(role, ShouldResemble, &role1)
			})
		})

		Convey("We have to be able to find an role from its rights", func() {
			roles := rsi.GetByRights(&rolesCanPrivate, ds)
			So(roles, ShouldNotResemble, emptyList)
			So(roles, ShouldResemble, &canPrivateList)
			roles = rsi.GetByRights(&rolesCanPrivateNotArchive, ds)
			So(roles, ShouldNotResemble, emptyList)
			So(roles, ShouldResemble, &canPrivateList)

		})

		Convey("Searching for non existent role should return empty", func() {
			role := rsi.GetByName("fantome", ds)
			So(role, ShouldResemble, &Role{})
		})

		db.Delete(&role0)
		db.Delete(&role1)
		db.Delete(&role2)
		db.Delete(&role3)

		Convey("Searching all in empty table should return empty", func() {
			roles := rsi.GetAll(ds)
			So(roles, ShouldResemble, &[]Role{})
		})
	})

	Convey("Testing delete role", t, func() {
		dberror := u.NewLocAppError("roleStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
		role0 := Role{
			RoleName:      "classe",
			CanUsePrivate: true,
			CanModerate:   false,
			CanArchive:    true,
			CanInvite:     false,
			CanManage:     false,
			CanManageUser: true,
		}
		role1 := Role{
			RoleName:      "deuce",
			CanUsePrivate: true,
			CanModerate:   false,
			CanArchive:    false,
			CanInvite:     false,
			CanManage:     false,
			CanManageUser: false,
		}
		role2 := Role{
			RoleName:      "fg",
			CanUsePrivate: true,
			CanModerate:   false,
			CanArchive:    true,
			CanInvite:     true,
			CanManage:     true,
			CanManageUser: true,
		}
		role3 := Role{
			RoleName:      "test",
			CanUsePrivate: true,
			CanModerate:   true,
			CanArchive:    true,
			CanInvite:     false,
			CanManage:     false,
			CanManageUser: true,
		}
		rsi.Save(&role0, ds)
		rsi.Save(&role1, ds)
		rsi.Save(&role2, ds)
		rsi.Save(&role3, ds)
		role3Old := role3
		roleList1 := []Role{
			role0,
			role1,
			role2,
			role3Old,
		}

		Convey("Deleting a known role should work", func() {
			appError := rsi.Delete(&role2, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dberror)
			So(rsi.GetByName("God", ds), ShouldResemble, &Role{})
		})

		Convey("Trying to delete from non conform role should return specific role error and should not delete roles.", func() {
			role3.RoleName = "Const"
			Convey("Too long or empty Name should return name error", func() {
				appError := rsi.Delete(&role3, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, u.NewLocAppError("roleStoreImpl.Delete.role.PreSave", "model.role.rolename.app_error", nil, ""))
				So(rsi.GetAll(ds), ShouldResemble, &roleList1)
				role3.RoleName = "+alpha"
				appError = rsi.Delete(&role3, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, u.NewLocAppError("roleStoreImpl.Delete.role.PreSave", "model.role.rolename.app_error", nil, ""))
				So(rsi.GetAll(ds), ShouldResemble, &roleList1)
				role3.RoleName = "alpha-numerique"
				appError = rsi.Delete(&role3, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, u.NewLocAppError("roleStoreImpl.Delete.role.PreSave", "model.role.rolename.app_error", nil, ""))
				So(rsi.GetAll(ds), ShouldResemble, &roleList1)
			})
		})

		db.Delete(&role0)
		db.Delete(&role1)
		db.Delete(&role2)
		db.Delete(&role3)
	})
}
