// This file is used to test if member model is working correctly.
// A member is always linked to a member
// He has basic member to join
package datastores

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestMemberStore(t *testing.T) {
	store := Store()
	db := store.InitConnection("root", "popcube_test", "popcube_dev", "database", "3306")

	msi := store.Member()
	usi := store.User()
	rsi := store.Role()
	csi := store.Channel()

	standartRole := Role{
		RoleName:      randStringBytes(10),
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	rsi.Save(&standartRole, db)

	channelRole := Role{
		RoleName:      randStringBytes(10),
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    false,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	rsi.Save(&channelRole, db)

	userTest := User{
		Username:  randStringBytes(10),
		Password:  "test",
		Email:     "test@popcube.fr",
		NickName:  "NickName",
		FirstName: "Test",
		LastName:  "L",
		IDRole:    standartRole.IDRole,
	}
	usi.Save(&userTest, db)

	channelTest := Channel{
		ChannelName: randStringBytes(10),
		Type:        "audio",
		Private:     false,
		Description: "Testing channel description :O",
		Subject:     "Sujet",
		Avatar:      "jesuiscool.svg",
	}
	csi.Save(&channelTest, db)

	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("memberStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("memberStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
		member := Member{
			IDUser:    userTest.IDUser,
			IDChannel: channelTest.IDChannel,
			IDRole:    channelRole.IDRole,
		}

		Convey("Given a correct member.", func() {
			appError := msi.Save(&member, db)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := msi.Save(&member, db)
				So(appError2, ShouldNotBeNil)
				// So(appError2, ShouldResemble, alreadyExistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&member)
	})

	// Convey("Testing update function", t, func() {
	// 	dbError := u.NewLocAppError("memberStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
	// 	alreadyExistError := u.NewLocAppError("memberStoreImpl.Save", "save.transaction.create.already_exist", nil, "Member Name: electras")
	// 	member := Member{
	// 		Membername: "TesT2",
	// 		Password:    "test",
	// 		Email:       "test2@popcube.fr",
	// 		NickName:    "NickName",
	// 		FirstName:   "Test",
	// 		LastName:    "L",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	memberNew := Member{
	// 		Membername: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		IDRole:      adminRole.IDRole,
	// 	}

	// 	appError := msi.Save(&member, db)
	// 	So(appError, ShouldBeNil)
	// 	So(appError, ShouldNotResemble, dbError)
	// 	So(appError, ShouldNotResemble, alreadyExistError)

	// 	Convey("Provided correct Member to modify should not return errors", func() {
	// 		appError := msi.Update(&member, &memberNew, db)
	// 		memberShouldResemble := memberNew
	// 		memberShouldResemble.WebID = member.WebID
	// 		memberShouldResemble.IDMember = member.IDMember
	// 		memberShouldResemble.LastUpdate = member.LastUpdate
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dbError)
	// 		So(appError, ShouldNotResemble, alreadyExistError)
	// 		So(member, ShouldResemble, memberShouldResemble)
	// 	})

	// 	Convey("Provided wrong old Member to modify should result in old_member error", func() {
	// 		member.WebID = "TesT"
	// 		Convey("Incorrect ID member should return a member invalid id", func() {
	// 			appError := msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberOld.PreSave", "model.member.is_valid.WebID.app_error", nil, ""))
	// 		})
	// 		member.WebID = NewID()
	// 		Convey("Incorrect membername member should return error Invalid membername", func() {
	// 			member.Membername = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
	// 			appError := msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberOld.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+member.WebID))
	// 			member.Membername = ""
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberOld.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+member.WebID))
	// 			member.Membername = "xD/"
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberOld.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+member.WebID))
	// 			member.Membername = "xD\\"
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberOld.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+member.WebID))
	// 			member.Membername = "xD*"
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberOld.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+member.WebID))
	// 			member.Membername = "xD{"
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberOld.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+member.WebID))
	// 		})

	// 		Convey("Password can]t be empty", func() {
	// 			member.Password = ""
	// 			appError := msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberOld.PreSave", "model.member.is_valid.auth_data_pwd.app_error", nil, "member_webID="+member.WebID))
	// 		})
	// 	})

	// 	Convey("Provided wrong new Member to modify should result in old_member error", func() {
	// 		Convey("Incorrect membername member should return error Invalid membername", func() {
	// 			memberNew.Membername = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
	// 			appError := msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberNew.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+memberNew.WebID))
	// 			memberNew.Membername = ""
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberNew.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+memberNew.WebID))
	// 			memberNew.Membername = "xD/"
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberNew.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+memberNew.WebID))
	// 			memberNew.Membername = "xD\\"
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberNew.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+memberNew.WebID))
	// 			memberNew.Membername = "xD*"
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberNew.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+memberNew.WebID))
	// 			memberNew.Membername = "xD{"
	// 			appError = msi.Update(&member, &memberNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Update.memberNew.PreSave", "model.member.is_valid.Membername.app_error", nil, "member_webID="+memberNew.WebID))
	// 		})
	// 	})

	// 	db.Delete(&member)
	// 	db.Delete(&memberNew)
	// })

	// Convey("Testing Getters", t, func() {
	// 	member0 := Member{
	// 		Membername: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		IDRole:      adminRole.IDRole,
	// 	}
	// 	member1 := Member{
	// 		Membername: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe",
	// 		LastName:    "Dalton",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	member2 := Member{
	// 		Membername: "moris",
	// 		Password:    "gossiny",
	// 		Email:       "moris&gossiny@popcube.fr",
	// 		NickName:    "Moris",
	// 		FirstName:   "Moris",
	// 		LastName:    "Gossiny",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	member1New := Member{
	// 		Membername: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe, Jack, William, Avrell",
	// 		LastName:    "Dalton",
	// 		IDRole:      guestRole.IDRole,
	// 	}
	// 	member3 := Member{
	// 		Membername: "jolly",
	// 		Password:    "jumper",
	// 		Email:       "jollyjumper@popcube.fr",
	// 		NickName:    "JJ",
	// 		FirstName:   "Jolly",
	// 		LastName:    "Jumper",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	member4 := Member{
	// 		Membername: "billythekid",
	// 		Password:    "chocolat",
	// 		Email:       "billythekid@popcube.fr",
	// 		NickName:    "Kid",
	// 		FirstName:   "Billy",
	// 		LastName:    "The Kid",
	// 		IDRole:      guestRole.IDRole,
	// 	}

	// 	msi.Save(&member0, db)
	// 	msi.Save(&member1, db)
	// 	msi.Update(&member1, &member1New, db)
	// 	msi.Save(&member2, db)
	// 	msi.Save(&member3, db)
	// 	msi.Save(&member4, db)

	// 	// Have to be after save so ID are up to date :O
	// 	memberList := []Member{
	// 		member0,
	// 		member1,
	// 		member2,
	// 		member3,
	// 		member4,
	// 	}

	// 	admins := []Member{member0}
	// 	guests := []Member{member1, member4}
	// 	emptyList := []Member{}

	// 	Convey("We have to be able to find all members in the db", func() {
	// 		members := msi.GetAll(db)
	// 		So(members, ShouldNotResemble, &emptyList)
	// 		So(members, ShouldResemble, &memberList)
	// 	})

	// 	Convey("We have to be able to find a member from is name", func() {
	// 		member := msi.GetByMemberName(member0.Membername, db)
	// 		So(member, ShouldNotResemble, &Member{})
	// 		So(member, ShouldResemble, &member0)
	// 		member = msi.GetByMemberName(member2.Membername, db)
	// 		So(member, ShouldNotResemble, &Member{})
	// 		So(member, ShouldResemble, &member2)
	// 		member = msi.GetByMemberName(member3.Membername, db)
	// 		So(member, ShouldNotResemble, &Member{})
	// 		So(member, ShouldResemble, &member3)
	// 		member = msi.GetByMemberName(member4.Membername, db)
	// 		So(member, ShouldNotResemble, &Member{})
	// 		So(member, ShouldResemble, &member4)
	// 		Convey("Should also work from updated value", func() {
	// 			member = msi.GetByMemberName(member1New.Membername, db)
	// 			So(member, ShouldNotResemble, &Member{})
	// 			So(member, ShouldResemble, &member1)
	// 		})
	// 	})

	// 	Convey("We have to be able to find a member from his email", func() {
	// 		member := msi.GetByEmail(member0.Email, db)
	// 		So(member, ShouldNotResemble, &Member{})
	// 		So(member, ShouldResemble, &member0)
	// 		member = msi.GetByEmail(member2.Email, db)
	// 		So(member, ShouldNotResemble, &Member{})
	// 		So(member, ShouldResemble, &member2)
	// 		member = msi.GetByEmail(member3.Email, db)
	// 		So(member, ShouldResemble, &member3)
	// 		member = msi.GetByEmail(member4.Email, db)
	// 		So(member, ShouldNotResemble, &Member{})
	// 		So(member, ShouldResemble, &member4)
	// 	})

	// 	Convey("We have to be able to find an member from his Role", func() {
	// 		members := msi.GetByRole(&adminRole, db)
	// 		So(members, ShouldNotResemble, &Member{})
	// 		So(members, ShouldResemble, &admins)
	// 		members = msi.GetByRole(&guestRole, db)
	// 		So(members, ShouldNotResemble, &Member{})
	// 		So(members, ShouldResemble, &guests)

	// 	})

	// 	Convey("Searching for non existent member should return empty", func() {
	// 		member := msi.GetByMemberName("fantome", db)
	// 		So(member, ShouldResemble, &Member{})
	// 	})

	// 	db.Delete(&member0)
	// 	db.Delete(&member1)
	// 	db.Delete(&member1New)
	// 	db.Delete(&member2)
	// 	db.Delete(&member3)

	// 	Convey("Searching all in empty table should return empty", func() {
	// 		members := msi.GetAll(db)
	// 		So(members, ShouldResemble, &[]Member{})
	// 	})
	// })

	// Convey("Testing delete member", t, func() {
	// 	dberror := u.NewLocAppError("memberStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
	// 	member0 := Member{
	// 		Membername: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		IDRole:      adminRole.IDRole,
	// 	}
	// 	member1 := Member{
	// 		Membername: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe",
	// 		LastName:    "Dalton",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	member2 := Member{
	// 		Membername: "moris",
	// 		Password:    "gossiny",
	// 		Email:       "moris&gossiny@popcube.fr",
	// 		NickName:    "Moris",
	// 		FirstName:   "Moris",
	// 		LastName:    "Gossiny",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	member3 := Member{
	// 		Membername: "jolly",
	// 		Password:    "jumper",
	// 		Email:       "jollyjumper@popcube.fr",
	// 		NickName:    "JJ",
	// 		FirstName:   "Jolly",
	// 		LastName:    "Jumper",
	// 		IDRole:      standartRole.IDRole,
	// 	}

	// 	msi.Save(&member0, db)
	// 	msi.Save(&member1, db)
	// 	msi.Save(&member2, db)
	// 	msi.Save(&member3, db)

	// 	// Have to be after save so ID are up to date :O
	// 	// member3Old := member3
	// 	// memberList1 := []Member{
	// 	// 	member0,
	// 	// 	member1,
	// 	// 	member2,
	// 	// 	member3Old,
	// 	// }

	// 	Convey("Deleting a known member should work", func() {
	// 		appError := msi.Delete(&member2, db)
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dberror)
	// 		So(msi.GetByMemberName("moris", db), ShouldResemble, &Member{})
	// 	})

	// 	// Convey("Trying to delete from non conform member should return specific member error and should not delete members.", func() {
	// 	// 	member3.MemberName = "Const"
	// 	// 	Convey("Too long or empty Name should return name error", func() {
	// 	// 		appError := msi.Delete(&member3, db)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Delete.member.PreSave", "model.member.membername.app_error", nil, ""))
	// 	// 		So(msi.GetAll(db), ShouldResemble, &memberList1)
	// 	// 		member3.MemberName = "+alpha"
	// 	// 		appError = msi.Delete(&member3, db)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Delete.member.PreSave", "model.member.membername.app_error", nil, ""))
	// 	// 		So(msi.GetAll(db), ShouldResemble, &memberList1)
	// 	// 		member3.MemberName = "alpha-numerique"
	// 	// 		appError = msi.Delete(&member3, db)standartRole
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("memberStoreImpl.Delete.member.PreSave", "model.member.membername.app_error", nil, ""))
	// 	// 		So(msi.GetAll(db), ShouldResemble, &memberList1)
	// 	// 	})
	// 	// })

	// 	db.Delete(&member0)
	// 	db.Delete(&member1)
	// 	db.Delete(&member2)
	// 	db.Delete(&member3)
	// })
	db.Delete(&userTest)
	db.Delete(&channelTest)
	db.Delete(&standartRole)
	db.Delete(&channelRole)
}
