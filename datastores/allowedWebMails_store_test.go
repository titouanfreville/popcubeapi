// This file is used to test if user model is working correctly.
// A user is always linked to an allowedWebMails
// He has basic channel to join
package datastores

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func genDbError(error string) *u.AppError {
	return u.NewAPIError(500, "unxepected error", "Unexpected error while adding entry. \n ---- "+error+"----")
}

func genDuplicateError(duplicateID string) *u.AppError {
	return u.NewAPIError(409, "duplicate entry", "You already authorized "+duplicateID+" mails to sign up.")
}

func genInvalidDataError(errorID string, errorMessage string) *u.AppError {
	return u.NewAPIError(422, errorID, "Wrong data provided : \n ---- "+errorMessage+" ----")
}

func TestAllowedWebMailsStore(t *testing.T) {
	store := Store()
	db := store.InitConnection("root", "popcube_test", "popcube_dev", "database", "3306")
	asi := store.AllowedWebMails()
	Convey("Testing save function", t, func() {
		allowedWebMails := AllowedWebMails{
			Domain:        "popcube.xyz",
			Provider:      "PopCube",
			DefaultRights: "Master",
		}
		db.Delete(&allowedWebMails)
		Convey("Given a correct allowedWebMails.", func() {
			appError := asi.Save(&allowedWebMails, db)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, genDbError("nil"))
				So(appError, ShouldNotResemble, genDuplicateError(allowedWebMails.Domain))
				So(appError, ShouldNotResemble, genInvalidDataError("test", "something"))
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := asi.Save(&allowedWebMails, db)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldNotResemble, genDbError(""))
				So(appError2, ShouldResemble, genDuplicateError(allowedWebMails.Domain))
				So(appError2, ShouldNotResemble, genInvalidDataError("test", "something"))
			})
		})
		db.Delete(&allowedWebMails)
	})

	Convey("Testing update function", t, func() {
		allowedWebMails := AllowedWebMails{
			Domain:        "popcube.xyz",
			Provider:      "PopCube",
			DefaultRights: "Master",
		}
		allowedWebMailsNew := AllowedWebMails{
			Domain:        "newpopcube.xyz",
			Provider:      "NewPopCube",
			DefaultRights: "NewMaster",
		}

		appError := asi.Save(&allowedWebMails, db)
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, genDbError(""))
		So(appError, ShouldNotResemble, genDuplicateError(allowedWebMails.Domain))
		So(appError, ShouldNotResemble, genInvalidDataError("test", "something"))

		Convey("Provided correct AllowedWebMails to modify should not return errors", func() {
			appError := asi.Update(&allowedWebMails, &allowedWebMailsNew, db)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, genDbError(""))
			So(appError, ShouldNotResemble, genDuplicateError(allowedWebMails.Domain))
			So(appError, ShouldNotResemble, genInvalidDataError("test", "something"))
		})

		// Convey("Provided wrong AllowedWebMails to modify should result in old_allowedWebMails error", func() {
		// 	allowedWebMails.Name = ""
		// 	Convey("Too long or empty Name should return name error", func() {
		// 		appError := asi.Update(&allowedWebMails, &allowedWebMailsNew, db)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("allowedWebMailsStoreImpl.Update.allowedWebMailsOld.PreSave", "model.allowedWebMails.name.app_error", nil, ""))
		// 		allowedWebMails.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
		// 		appError = asi.Update(&allowedWebMails, &allowedWebMailsNew, db)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("allowedWebMailsStoreImpl.Update.allowedWebMailsOld.PreSave", "model.allowedWebMails.name.app_error", nil, ""))
		// 	})

		// 	allowedWebMails.Name = "Correct Name"
		// 	allowedWebMails.Link = ""

		// 	Convey("Empty link should result in link error", func() {
		// 		appError = asi.Update(&allowedWebMails, &allowedWebMailsNew, db)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("allowedWebMailsStoreImpl.Update.allowedWebMailsOld.PreSave", "model.allowedWebMails.link.app_error", nil, ""))
		// 	})
		// })

		// Convey("Provided wrong AllowedWebMails to modify should result in newAllowedWebMails error", func() {
		// allowedWebMailsNew.Name = ""
		// Convey("Too long or empty Name should return name error", func() {
		// 	appError := asi.Update(&allowedWebMails, &allowedWebMailsNew, db)
		// 	So(appError, ShouldNotBeNil)
		// 	So(appError, ShouldNotResemble, dbError)
		// 	So(appError, ShouldResemble, u.NewLocAppError("allowedWebMailsStoreImpl.Update.allowedWebMailsNew.PreSave", "model.allowedWebMails.name.app_error", nil, ""))
		// 	allowedWebMailsNew.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
		// 	appError = asi.Update(&allowedWebMails, &allowedWebMailsNew, db)
		// 	So(appError, ShouldNotBeNil)
		// 	So(appError, ShouldNotResemble, dbError)
		// 	So(appError, ShouldResemble, u.NewLocAppError("allowedWebMailsStoreImpl.Update.allowedWebMailsNew.PreSave", "model.allowedWebMails.name.app_error", nil, ""))
		// })

		// allowedWebMailsNew.Name = "Correct Name"
		// allowedWebMailsNew.Link = ""

		// Convey("Empty link should result in link error", func() {
		// 	appError = asi.Update(&allowedWebMails, &allowedWebMailsNew, db)
		// 	So(appError, ShouldNotBeNil)
		// 	So(appError, ShouldNotResemble, dbError)
		// 	So(appError, ShouldResemble, u.NewLocAppError("allowedWebMailsStoreImpl.Update.allowedWebMailsNew.PreSave", "model.allowedWebMails.link.app_error", nil, ""))
		// })
		// })
		db.Delete(&allowedWebMails)
		db.Delete(&allowedWebMailsNew)
	})
	// Convey("Testing Getters", t, func() {
	// 	allowedWebMails0 := AllowedWebMails{
	// 		Name: "Troll Face",
	// 		Link: "allowedWebMailss/trollface.svg",
	// 	}
	// 	allowedWebMails1 := AllowedWebMails{
	// 		Name: "Face Palm Old",
	// 		Link: "allowedWebMailss/facepalmold.svg",
	// 	}
	// 	allowedWebMails1New := AllowedWebMails{
	// 		Name: "Face Palm",
	// 		Link: "allowedWebMailss/facepalm.svg",
	// 	}
	// 	allowedWebMails2 := AllowedWebMails{
	// 		Name: "God",
	// 		Link: "allowedWebMailss/docker.svg",
	// 	}
	// 	allowedWebMails3 := AllowedWebMails{
	// 		Name: "nOOb",
	// 		Link: "allowedWebMailss/sparadra.svg",
	// 	}
	// 	asi.Save(&allowedWebMails0, db)
	// 	asi.Save(&allowedWebMails1, db)
	// 	asi.Update(&allowedWebMails1, &allowedWebMails1New, db)
	// 	asi.Save(&allowedWebMails2, db)
	// 	asi.Save(&allowedWebMails3, db)
	// 	// Have to be after save so ID are up to date :O
	// 	allowedWebMailsList := []AllowedWebMails{
	// 		allowedWebMails0,
	// 		allowedWebMails1,
	// 		allowedWebMails2,
	// 		allowedWebMails3,
	// 	}

	// 	Convey("We have to be able to find all allowedWebMailss in the db", func() {
	// 		allowedWebMailss := asi.GetAll(db)
	// 		So(allowedWebMailss, ShouldNotResemble, []AllowedWebMails{})
	// 		So(allowedWebMailss, ShouldResemble, allowedWebMailsList)
	// 	})

	// 	Convey("We have to be able to find an allowedWebMails from is name", func() {
	// 		allowedWebMails := asi.GetByName(allowedWebMails0.Name, db)
	// 		So(allowedWebMails, ShouldNotResemble, AllowedWebMails{})
	// 		So(allowedWebMails, ShouldResemble, allowedWebMails0)
	// 		allowedWebMails = asi.GetByName(allowedWebMails2.Name, db)
	// 		So(allowedWebMails, ShouldNotResemble, AllowedWebMails{})
	// 		So(allowedWebMails, ShouldResemble, allowedWebMails2)
	// 		allowedWebMails = asi.GetByName(allowedWebMails3.Name, db)
	// 		So(allowedWebMails, ShouldNotResemble, AllowedWebMails{})
	// 		So(allowedWebMails, ShouldResemble, allowedWebMails3)
	// 		Convey("Should also work from updated value", func() {
	// 			allowedWebMails = asi.GetByName(allowedWebMails1.Name, db)
	// 			So(allowedWebMails, ShouldNotResemble, AllowedWebMails{})
	// 			So(allowedWebMails, ShouldResemble, allowedWebMails1)
	// 		})
	// 	})

	// 	Convey("We have to be able to find an allowedWebMails from is link", func() {
	// 		allowedWebMails := asi.GetByLink(allowedWebMails0.Link, db)
	// 		So(allowedWebMails, ShouldNotResemble, AllowedWebMails{})
	// 		So(allowedWebMails, ShouldResemble, allowedWebMails0)
	// 		allowedWebMails = asi.GetByLink(allowedWebMails2.Link, db)
	// 		So(allowedWebMails, ShouldNotResemble, AllowedWebMails{})
	// 		So(allowedWebMails, ShouldResemble, allowedWebMails2)
	// 		allowedWebMails = asi.GetByLink(allowedWebMails3.Link, db)
	// 		So(allowedWebMails, ShouldNotResemble, AllowedWebMails{})
	// 		So(allowedWebMails, ShouldResemble, allowedWebMails3)
	// 		Convey("Should also work from updated value", func() {
	// 			allowedWebMails = asi.GetByLink(allowedWebMails1.Link, db)
	// 			So(allowedWebMails, ShouldNotResemble, AllowedWebMails{})
	// 			So(allowedWebMails, ShouldResemble, allowedWebMails1)
	// 		})
	// 	})

	// 	Convey("Searching for non existent allowedWebMails should return empty", func() {
	// 		allowedWebMails := asi.GetByLink("The Mask", db)
	// 		So(allowedWebMails, ShouldResemble, AllowedWebMails{})
	// 		allowedWebMails = asi.GetByName("Fantôme", db)
	// 		So(allowedWebMails, ShouldResemble, AllowedWebMails{})
	// 	})

	// 	db.Delete(&allowedWebMails0)
	// 	db.Delete(&allowedWebMails1)
	// 	db.Delete(&allowedWebMails2)
	// 	db.Delete(&allowedWebMails3)

	// 	Convey("Searching all in empty table should return empty", func() {
	// 		allowedWebMailss := asi.GetAll(db)
	// 		So(allowedWebMailss, ShouldResemble, []AllowedWebMails{})
	// 	})
	// })

	// Convey("Testing delete allowedWebMails", t, func() {
	// 	dberror := u.NewLocAppError("allowedWebMailsStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
	// 	allowedWebMails0 := AllowedWebMails{
	// 		Name: "Troll Face",
	// 		Link: "allowedWebMailss/trollface.svg",
	// 	}
	// 	allowedWebMails1 := AllowedWebMails{
	// 		Name: "Face Palm",
	// 		Link: "allowedWebMailss/facepal.svg",
	// 	}
	// 	allowedWebMails2 := AllowedWebMails{
	// 		Name: "God",
	// 		Link: "allowedWebMailss/docker.svg",
	// 	}
	// 	allowedWebMails3 := AllowedWebMails{
	// 		Name: "nOOb",
	// 		Link: "allowedWebMailss/sparadra.svg",
	// 	}
	// 	asi.Save(&allowedWebMails0, db)
	// 	asi.Save(&allowedWebMails1, db)
	// 	asi.Save(&allowedWebMails2, db)
	// 	asi.Save(&allowedWebMails3, db)
	// 	allowedWebMails3Old := allowedWebMails3
	// 	allowedWebMailsList1 := []AllowedWebMails{
	// 		allowedWebMails0,
	// 		allowedWebMails1,
	// 		allowedWebMails2,
	// 		allowedWebMails3Old,
	// 	}

	// 	Convey("Deleting a known allowedWebMails should work", func() {
	// 		appError := asi.Delete(&allowedWebMails2, db)
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dberror)
	// 		So(asi.GetByName("God", db), ShouldResemble, AllowedWebMails{})
	// 	})

	// 	Convey("Trying to delete from non conform allowedWebMails should return specific allowedWebMails error and should not delete allowedWebMailss.", func() {
	// 		allowedWebMails3.Name = ""
	// 		Convey("Too long or empty Name should return name error", func() {
	// 			appError := asi.Delete(&allowedWebMails3, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dberror)
	// 			So(appError, ShouldResemble, u.NewLocAppError("allowedWebMailsStoreImpl.Delete.allowedWebMails.PreSave", "model.allowedWebMails.name.app_error", nil, ""))
	// 			So(asi.GetAll(db), ShouldResemble, allowedWebMailsList1)
	// 			allowedWebMails3.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
	// 			appError = asi.Delete(&allowedWebMails3, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dberror)
	// 			So(appError, ShouldResemble, u.NewLocAppError("allowedWebMailsStoreImpl.Delete.allowedWebMails.PreSave", "model.allowedWebMails.name.app_error", nil, ""))
	// 			So(asi.GetAll(db), ShouldResemble, allowedWebMailsList1)
	// 		})

	// 		allowedWebMails3.Name = "nOOb"
	// 		allowedWebMails3.Link = ""

	// 		Convey("Empty link should result in link error", func() {
	// 			appError := asi.Delete(&allowedWebMails3, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dberror)
	// 			So(appError, ShouldResemble, u.NewLocAppError("allowedWebMailsStoreImpl.Delete.allowedWebMails.PreSave", "model.allowedWebMails.link.app_error", nil, ""))
	// 			So(asi.GetAll(db), ShouldResemble, allowedWebMailsList1)
	// 		})
	// 	})

	// 	db.Delete(&allowedWebMails0)
	// 	db.Delete(&allowedWebMails1)
	// 	db.Delete(&allowedWebMails2)
	// 	db.Delete(&allowedWebMails3)
	// })
}
