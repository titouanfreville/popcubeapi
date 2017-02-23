// This file is used to test if user model is working correctly.
// A user is always linked to an avatar
// He has basic channel to join
package datastores

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestAvatarStore(t *testing.T) {
	store := Store()
	db := store.InitConnection("root", "popcube_test", "popcube_dev", "database", "3306")
	asi := store.Avatar()
	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("avatarStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyexistError := u.NewLocAppError("avatarStoreImpl.Save", "save.transaction.create.already_exist", nil, "Avatar Name: Troll Face")
		avatar := Avatar{
			Name: "Troll Face",
			Link: "avatars/trollface.svg",
		}
		Convey("Given a correct avatar.", func() {
			appError := asi.Save(&avatar, db)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := asi.Save(&avatar, db)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyexistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&avatar)
	})

	Convey("Testing update function", t, func() {
		dbError := u.NewLocAppError("avatarStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyexistError := u.NewLocAppError("avatarStoreImpl.Save", "save.transaction.create.already_exist", nil, "Avatar Name: Troll Face")
		avatar := Avatar{
			Name: "Troll Face",
			Link: "avatars/trollface.svg",
		}
		avatarNew := Avatar{
			Name: "TrollFace2",
			Link: "avatars/trollface2.svg",
		}

		appError := asi.Save(&avatar, db)
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		So(appError, ShouldNotResemble, alreadyexistError)

		Convey("Provided correct Avatar to modify should not return errors", func() {
			appError := asi.Update(&avatar, &avatarNew, db)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
			So(appError, ShouldNotResemble, alreadyexistError)
		})

		Convey("Provided wrong Avatar to modify should result in old_avatar error", func() {
			avatar.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				appError := asi.Update(&avatar, &avatarNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("avatarStoreImpl.Update.avatarOld.PreSave", "model.avatar.name.app_error", nil, ""))
				avatar.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				appError = asi.Update(&avatar, &avatarNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("avatarStoreImpl.Update.avatarOld.PreSave", "model.avatar.name.app_error", nil, ""))
			})

			avatar.Name = "Correct Name"
			avatar.Link = ""

			Convey("Empty link should result in link error", func() {
				appError = asi.Update(&avatar, &avatarNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("avatarStoreImpl.Update.avatarOld.PreSave", "model.avatar.link.app_error", nil, ""))
			})
		})

		Convey("Provided wrong Avatar to modify should result in newAvatar error", func() {
			avatarNew.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				appError := asi.Update(&avatar, &avatarNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("avatarStoreImpl.Update.avatarNew.PreSave", "model.avatar.name.app_error", nil, ""))
				avatarNew.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				appError = asi.Update(&avatar, &avatarNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("avatarStoreImpl.Update.avatarNew.PreSave", "model.avatar.name.app_error", nil, ""))
			})

			avatarNew.Name = "Correct Name"
			avatarNew.Link = ""

			Convey("Empty link should result in link error", func() {
				appError = asi.Update(&avatar, &avatarNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("avatarStoreImpl.Update.avatarNew.PreSave", "model.avatar.link.app_error", nil, ""))
			})
		})
		db.Delete(&avatar)
		db.Delete(&avatarNew)
	})

	Convey("Testing Getters", t, func() {
		avatar0 := Avatar{
			Name: "Troll Face",
			Link: "avatars/trollface.svg",
		}
		avatar1 := Avatar{
			Name: "Face Palm Old",
			Link: "avatars/facepalmold.svg",
		}
		avatar1New := Avatar{
			Name: "Face Palm",
			Link: "avatars/facepalm.svg",
		}
		avatar2 := Avatar{
			Name: "God",
			Link: "avatars/docker.svg",
		}
		avatar3 := Avatar{
			Name: "nOOb",
			Link: "avatars/sparadra.svg",
		}
		asi.Save(&avatar0, db)
		asi.Save(&avatar1, db)
		asi.Update(&avatar1, &avatar1New, db)
		asi.Save(&avatar2, db)
		asi.Save(&avatar3, db)
		// Have to be after save so ID are up to date :O
		avatarList := []Avatar{
			avatar0,
			avatar1,
			avatar2,
			avatar3,
		}

		Convey("We have to be able to find all avatars in the db", func() {
			avatars := asi.GetAll(db)
			So(avatars, ShouldNotResemble, []Avatar{})
			So(avatars, ShouldResemble, avatarList)
		})

		Convey("We have to be able to find an avatar from is name", func() {
			avatar := asi.GetByName(avatar0.Name, db)
			So(avatar, ShouldNotResemble, Avatar{})
			So(avatar, ShouldResemble, avatar0)
			avatar = asi.GetByName(avatar2.Name, db)
			So(avatar, ShouldNotResemble, Avatar{})
			So(avatar, ShouldResemble, avatar2)
			avatar = asi.GetByName(avatar3.Name, db)
			So(avatar, ShouldNotResemble, Avatar{})
			So(avatar, ShouldResemble, avatar3)
			Convey("Should also work from updated value", func() {
				avatar = asi.GetByName(avatar1.Name, db)
				So(avatar, ShouldNotResemble, Avatar{})
				So(avatar, ShouldResemble, avatar1)
			})
		})

		Convey("We have to be able to find an avatar from is link", func() {
			avatar := asi.GetByLink(avatar0.Link, db)
			So(avatar, ShouldNotResemble, Avatar{})
			So(avatar, ShouldResemble, avatar0)
			avatar = asi.GetByLink(avatar2.Link, db)
			So(avatar, ShouldNotResemble, Avatar{})
			So(avatar, ShouldResemble, avatar2)
			avatar = asi.GetByLink(avatar3.Link, db)
			So(avatar, ShouldNotResemble, Avatar{})
			So(avatar, ShouldResemble, avatar3)
			Convey("Should also work from updated value", func() {
				avatar = asi.GetByLink(avatar1.Link, db)
				So(avatar, ShouldNotResemble, Avatar{})
				So(avatar, ShouldResemble, avatar1)
			})
		})

		Convey("Searching for non existent avatar should return empty", func() {
			avatar := asi.GetByLink("The Mask", db)
			So(avatar, ShouldResemble, Avatar{})
			avatar = asi.GetByName("Fantôme", db)
			So(avatar, ShouldResemble, Avatar{})
		})

		db.Delete(&avatar0)
		db.Delete(&avatar1)
		db.Delete(&avatar2)
		db.Delete(&avatar3)

		Convey("Searching all in empty table should return empty", func() {
			avatars := asi.GetAll(db)
			So(avatars, ShouldResemble, []Avatar{})
		})
	})

	Convey("Testing delete avatar", t, func() {
		dberror := u.NewLocAppError("avatarStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
		avatar0 := Avatar{
			Name: "Troll Face",
			Link: "avatars/trollface.svg",
		}
		avatar1 := Avatar{
			Name: "Face Palm",
			Link: "avatars/facepal.svg",
		}
		avatar2 := Avatar{
			Name: "God",
			Link: "avatars/docker.svg",
		}
		avatar3 := Avatar{
			Name: "nOOb",
			Link: "avatars/sparadra.svg",
		}
		asi.Save(&avatar0, db)
		asi.Save(&avatar1, db)
		asi.Save(&avatar2, db)
		asi.Save(&avatar3, db)
		avatar3Old := avatar3
		avatarList1 := []Avatar{
			avatar0,
			avatar1,
			avatar2,
			avatar3Old,
		}

		Convey("Deleting a known avatar should work", func() {
			appError := asi.Delete(&avatar2, db)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dberror)
			So(asi.GetByName("God", db), ShouldResemble, Avatar{})
		})

		Convey("Trying to delete from non conform avatar should return specific avatar error and should not delete avatars.", func() {
			avatar3.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				appError := asi.Delete(&avatar3, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, u.NewLocAppError("avatarStoreImpl.Delete.avatar.PreSave", "model.avatar.name.app_error", nil, ""))
				So(asi.GetAll(db), ShouldResemble, avatarList1)
				avatar3.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				appError = asi.Delete(&avatar3, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, u.NewLocAppError("avatarStoreImpl.Delete.avatar.PreSave", "model.avatar.name.app_error", nil, ""))
				So(asi.GetAll(db), ShouldResemble, avatarList1)
			})

			avatar3.Name = "nOOb"
			avatar3.Link = ""

			Convey("Empty link should result in link error", func() {
				appError := asi.Delete(&avatar3, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, u.NewLocAppError("avatarStoreImpl.Delete.avatar.PreSave", "model.avatar.link.app_error", nil, ""))
				So(asi.GetAll(db), ShouldResemble, avatarList1)
			})
		})

		db.Delete(&avatar0)
		db.Delete(&avatar1)
		db.Delete(&avatar2)
		db.Delete(&avatar3)
	})
}
