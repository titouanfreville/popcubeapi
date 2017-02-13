// This file is used to test if user model is working correctly.
// A user is always linked to a user
// He has basic user to join
package datastores

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestUserStore(t *testing.T) {
	ds := DbStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev")
	db := *ds.Db

	usi := NewUserStore()
	rsi := NewRoleStore()

	ownerRole := Role{
		RoleName:      "testowner",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     true,
		CanManage:     true,
		CanManageUser: true,
	}
	rsi.Save(&ownerRole, ds)

	adminRole := Role{
		RoleName:      "testadmin",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     true,
		CanManage:     true,
		CanManageUser: true,
	}
	rsi.Save(&adminRole, ds)
	standartRole := Role{
		RoleName:      "teststandart",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	rsi.Save(&standartRole, ds)
	guestRole := Role{
		RoleName:      "testguest",
		CanUsePrivate: false,
		CanModerate:   false,
		CanArchive:    false,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	rsi.Save(&guestRole, ds)

	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.already_exist", nil, "User Name: test")
		user := User{
			Username:  "TesT",
			Password:  "test",
			Email:     "test@popcube.fr",
			NickName:  "NickName",
			FirstName: "Test",
			LastName:  "L",
			IDRole:    ownerRole.IDRole,
		}

		Convey("Given a correct user.", func() {
			appError := usi.Save(&user, ds)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := usi.Save(&user, ds)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyExistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&user)
	})

	Convey("Testing update function", t, func() {
		dbError := u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.already_exist", nil, "User Name: electras")
		user := User{
			Username:  "TesT2",
			Password:  "test",
			Email:     "test2@popcube.fr",
			NickName:  "NickName",
			FirstName: "Test",
			LastName:  "L",
			IDRole:    ownerRole.IDRole,
		}
		userNew := User{
			Username:  "lucky",
			Password:  "lucke",
			Email:     "luckylucke@popcube.fr",
			NickName:  "LL",
			FirstName: "Luky",
			LastName:  "Luke",
			Locale:    "vn_VN",
			IDRole:    adminRole.IDRole,
		}

		appError := usi.Save(&user, ds)
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		So(appError, ShouldNotResemble, alreadyExistError)

		Convey("Provided correct User to modify should not return errors", func() {
			appError := usi.Update(&user, &userNew, ds)
			userShouldResemble := userNew
			userShouldResemble.WebID = user.WebID
			userShouldResemble.IDUser = user.IDUser
			userShouldResemble.LastUpdate = user.LastUpdate
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
			So(appError, ShouldNotResemble, alreadyExistError)
			So(user, ShouldResemble, userShouldResemble)
		})

		Convey("Provided wrong old User to modify should result in old_user error", func() {
			user.WebID = "TesT"
			Convey("Incorrect ID user should return a message invalid id", func() {
				appError := usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userOld.PreSave", "model.user.is_valid.WebID.app_error", nil, ""))
			})
			user.WebID = NewID()
			Convey("Incorrect username user should return error Invalid username", func() {
				user.Username = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
				appError := usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userOld.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+user.WebID))
				user.Username = ""
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userOld.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+user.WebID))
				user.Username = "xD/"
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userOld.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+user.WebID))
				user.Username = "xD\\"
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userOld.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+user.WebID))
				user.Username = "xD*"
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userOld.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+user.WebID))
				user.Username = "xD{"
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userOld.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+user.WebID))
			})

			Convey("Password can]t be empty", func() {
				user.Password = ""
				appError := usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userOld.PreSave", "model.user.is_valid.auth_data_pwd.app_error", nil, "user_webID="+user.WebID))
			})
		})

		Convey("Provided wrong new User to modify should result in old_user error", func() {
			Convey("Incorrect username user should return error Invalid username", func() {
				userNew.Username = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
				appError := usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+userNew.WebID))
				userNew.Username = ""
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+userNew.WebID))
				userNew.Username = "xD/"
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+userNew.WebID))
				userNew.Username = "xD\\"
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+userNew.WebID))
				userNew.Username = "xD*"
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+userNew.WebID))
				userNew.Username = "xD{"
				appError = usi.Update(&user, &userNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.Username.app_error", nil, "user_webID="+userNew.WebID))
			})
		})

		db.Delete(&user)
		db.Delete(&userNew)
	})

	Convey("Testing Getters", t, func() {
		user0 := User{
			Username:  "lucky",
			Password:  "lucke",
			Email:     "luckylucke@popcube.fr",
			NickName:  "LL",
			FirstName: "Luky",
			LastName:  "Luke",
			Locale:    "en_EN",
			IDRole:    adminRole.IDRole,
		}
		user1 := User{
			Username:  "daltons",
			Password:  "dalton",
			Email:     "daltonsbrothers@popcube.fr",
			NickName:  "thebrothers",
			FirstName: "Joe",
			LastName:  "Dalton",
			Locale:    "en_EN",
			IDRole:    standartRole.IDRole,
		}
		user2 := User{
			Username:  "moris",
			Password:  "gossiny",
			Email:     "moris&gossiny@popcube.fr",
			NickName:  "Moris",
			FirstName: "Moris",
			LastName:  "Gossiny",
			Locale:    "fr_FR",
			IDRole:    ownerRole.IDRole,
		}
		user1New := User{
			Username:  "daltons",
			Password:  "dalton",
			Email:     "daltonsbrothers@popcube.fr",
			NickName:  "thebrothers",
			FirstName: "Joe, Jack, William, Avrell",
			LastName:  "Dalton",
			Locale:    "en_EN",
			IDRole:    guestRole.IDRole,
		}
		user3 := User{
			Username:  "jolly",
			Password:  "jumper",
			Email:     "jollyjumper@popcube.fr",
			NickName:  "JJ",
			FirstName: "Jolly",
			LastName:  "Jumper",
			Locale:    "en_EN",
			IDRole:    standartRole.IDRole,
		}
		user4 := User{
			Username:  "billythekid",
			Password:  "chocolat",
			Email:     "billythekid@popcube.fr",
			NickName:  "Kid",
			FirstName: "Billy",
			LastName:  "The Kid",
			Locale:    "en_EN",
			IDRole:    guestRole.IDRole,
		}

		usi.Save(&user0, ds)
		usi.Save(&user1, ds)
		usi.Update(&user1, &user1New, ds)
		usi.Save(&user2, ds)
		usi.Save(&user3, ds)
		usi.Save(&user4, ds)

		// Have to be after save so ID are up to date :O
		userList := []User{
			user0,
			user1,
			user2,
			user3,
			user4,
		}

		admins := []User{user0}
		guests := []User{user1, user4}
		emptyList := []User{}

		Convey("We have to be able to find all users in the db", func() {
			users := usi.GetAll(ds)
			So(users, ShouldNotResemble, &emptyList)
			So(users, ShouldResemble, &userList)
		})

		Convey("We have to be able to find a user from is name", func() {
			user := usi.GetByUserName(user0.Username, ds)
			So(user, ShouldNotResemble, &User{})
			So(user, ShouldResemble, &user0)
			user = usi.GetByUserName(user2.Username, ds)
			So(user, ShouldNotResemble, &User{})
			So(user, ShouldResemble, &user2)
			user = usi.GetByUserName(user3.Username, ds)
			So(user, ShouldNotResemble, &User{})
			So(user, ShouldResemble, &user3)
			user = usi.GetByUserName(user4.Username, ds)
			So(user, ShouldNotResemble, &User{})
			So(user, ShouldResemble, &user4)
			Convey("Should also work from updated value", func() {
				user = usi.GetByUserName(user1New.Username, ds)
				So(user, ShouldNotResemble, &User{})
				So(user, ShouldResemble, &user1)
			})
		})

		Convey("We have to be able to find a user from his email", func() {
			user := usi.GetByEmail(user0.Email, ds)
			So(user, ShouldNotResemble, &User{})
			So(user, ShouldResemble, &user0)
			user = usi.GetByEmail(user2.Email, ds)
			So(user, ShouldNotResemble, &User{})
			So(user, ShouldResemble, &user2)
			user = usi.GetByEmail(user3.Email, ds)
			So(user, ShouldResemble, &user3)
			user = usi.GetByEmail(user4.Email, ds)
			So(user, ShouldNotResemble, &User{})
			So(user, ShouldResemble, &user4)
		})

		Convey("We have to be able to find an user from his Role", func() {
			users := usi.GetByRole(&adminRole, ds)
			So(users, ShouldNotResemble, &User{})
			So(users, ShouldResemble, &admins)
			users = usi.GetByRole(&guestRole, ds)
			So(users, ShouldNotResemble, &User{})
			So(users, ShouldResemble, &guests)

		})

		Convey("Searching for non existent user should return empty", func() {
			user := usi.GetByUserName("fantome", ds)
			So(user, ShouldResemble, &User{})
		})

		db.Delete(&user0)
		db.Delete(&user1)
		db.Delete(&user1New)
		db.Delete(&user2)
		db.Delete(&user3)

		Convey("Searching all in empty table should return empty", func() {
			users := usi.GetAll(ds)
			So(users, ShouldResemble, &[]User{})
		})
	})

	Convey("Testing delete user", t, func() {
		dberror := u.NewLocAppError("userStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
		user0 := User{
			Username:  "lucky",
			Password:  "lucke",
			Email:     "luckylucke@popcube.fr",
			NickName:  "LL",
			FirstName: "Luky",
			LastName:  "Luke",
			Locale:    "en_EN",
			IDRole:    adminRole.IDRole,
		}
		user1 := User{
			Username:  "daltons",
			Password:  "dalton",
			Email:     "daltonsbrothers@popcube.fr",
			NickName:  "thebrothers",
			FirstName: "Joe",
			LastName:  "Dalton",
			Locale:    "en_EN",
			IDRole:    standartRole.IDRole,
		}
		user2 := User{
			Username:  "moris",
			Password:  "gossiny",
			Email:     "moris&gossiny@popcube.fr",
			NickName:  "Moris",
			FirstName: "Moris",
			LastName:  "Gossiny",
			Locale:    "fr_FR",
			IDRole:    ownerRole.IDRole,
		}
		user3 := User{
			Username:  "jolly",
			Password:  "jumper",
			Email:     "jollyjumper@popcube.fr",
			NickName:  "JJ",
			FirstName: "Jolly",
			LastName:  "Jumper",
			Locale:    "en_EN",
			IDRole:    standartRole.IDRole,
		}

		usi.Save(&user0, ds)
		usi.Save(&user1, ds)
		usi.Save(&user2, ds)
		usi.Save(&user3, ds)

		// Have to be after save so ID are up to date :O
		// user3Old := user3
		// userList1 := []User{
		// 	user0,
		// 	user1,
		// 	user2,
		// 	user3Old,
		// }

		Convey("Deleting a known user should work", func() {
			appError := usi.Delete(&user2, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dberror)
			So(usi.GetByUserName("moris", ds), ShouldResemble, &User{})
		})

		// Convey("Trying to delete from non conform user should return specific user error and should not delete users.", func() {
		// 	user3.UserName = "Const"
		// 	Convey("Too long or empty Name should return name error", func() {
		// 		appError := usi.Delete(&user3, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dberror)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Delete.user.PreSave", "model.user.username.app_error", nil, ""))
		// 		So(usi.GetAll(ds), ShouldResemble, &userList1)
		// 		user3.UserName = "+alpha"
		// 		appError = usi.Delete(&user3, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dberror)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Delete.user.PreSave", "model.user.username.app_error", nil, ""))
		// 		So(usi.GetAll(ds), ShouldResemble, &userList1)
		// 		user3.UserName = "alpha-numerique"
		// 		appError = usi.Delete(&user3, ds)standartRole
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dberror)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Delete.user.PreSave", "model.user.username.app_error", nil, ""))
		// 		So(usi.GetAll(ds), ShouldResemble, &userList1)
		// 	})
		// })

		db.Delete(&user0)
		db.Delete(&user1)
		db.Delete(&user2)
		db.Delete(&user3)
		db.Delete(&adminRole)
		db.Delete(&standartRole)
		db.Delete(&ownerRole)
		db.Delete(&guestRole)
	})
}
