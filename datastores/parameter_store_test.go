// This file is used to test if user model is working correctly.
// A user is always linked to an parameter
// He has bpsic channel to join
package datastores

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestParameterStore(t *testing.T) {
	store := Store()
	db := store.InitConnection("root", "popcube_test", "popcube_dev", "database", "3306")

	psi := store.Parameter()
	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("parameterStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyexistError := u.NewLocAppError("parameterStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
		parameter := Parameter{
			Local:      "en_EN",
			TimeZone:   "UTC+2",
			SleepStart: 280,
			SleepEnd:   12,
		}
		Convey("Given a correct parameter.", func() {
			appError := psi.Save(&parameter, db)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := psi.Save(&parameter, db)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyexistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		Convey("Given an incorrect parameter.", func() {
			empty := Parameter{}
			Convey("Empty parameter should return first error from is valid error", func() {
				appError := psi.Save(&empty, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_ENG"
			Convey("Given empty local or too long local should return Local error", func() {
				appError := psi.Save(&parameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
				parameter.Local = ""
				appError = psi.Save(&parameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_EN"
			parameter.TimeZone = "UTF+134"
			Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
				appError := psi.Save(&parameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
				parameter.TimeZone = ""
				appError = psi.Save(&parameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
			})
			parameter.TimeZone = "UTF+12"
			parameter.SleepEnd = -1
			Convey("Given negative or too big Sleep timers should return sleep error", func() {
				appError := psi.Save(&parameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 1441
				appError = psi.Save(&parameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 10
				parameter.SleepStart = -10
				appError = psi.Save(&parameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
				parameter.SleepStart = 2000
				appError = psi.Save(&parameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
			})
		})
		db.Delete(&parameter)
	})

	Convey("Testing update function", t, func() {
		parameter := Parameter{
			Local:      "en_EN",
			TimeZone:   "UTC+2",
			SleepStart: 280,
			SleepEnd:   12,
		}
		newParameter := Parameter{
			Local:      "vi_VI",
			TimeZone:   "UTC+6",
			SleepStart: 260,
			SleepEnd:   24,
		}
		appError := psi.Save(&parameter, db)
		dbError := u.NewLocAppError("parameterStoreImpl.Update", "update.transaction.updates.encounterError", nil, "")
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		Convey("Providing a correct user to update", func() {
			appError := psi.Update(&parameter, &newParameter, db)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
		})
		// Convey("Providing an incorrect user as new should result in errors", func() {
		// 	empty := Parameter{}
		// 	Convey("Empty parameter should return first error from is valid error", func() {
		// 		appError := psi.Update(&parameter, &empty, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
		// 	})
		// 	newParameter.Local = "en_ENG"
		// 	Convey("Given empty local or too long local should return Local error", func() {
		// 		appError := psi.Update(&parameter, &newParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
		// 		newParameter.Local = ""
		// 		appError = psi.Update(&parameter, &newParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
		// 	})
		// 	newParameter.Local = "en_EN"
		// 	newParameter.TimeZone = "UTF+134"
		// 	Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
		// 		appError := psi.Update(&parameter, &newParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
		// 		newParameter.TimeZone = ""
		// 		appError = psi.Update(&parameter, &newParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
		// 	})
		// 	newParameter.TimeZone = "UTF+12"
		// 	newParameter.SleepEnd = -1
		// 	Convey("Given negative or too big Sleep timers should return sleep error", func() {
		// 		appError := psi.Update(&parameter, &newParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
		// 		newParameter.SleepEnd = 1441
		// 		appError = psi.Update(&parameter, &newParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
		// 		newParameter.SleepEnd = 10
		// 		newParameter.SleepStart = -10
		// 		appError = psi.Update(&parameter, &newParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
		// 		newParameter.SleepStart = 2000
		// 		appError = psi.Update(&parameter, &newParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
		// 	})
		// })

		Convey("Providing an incorrect user as old should result in errors", func() {
			empty := Parameter{}
			Convey("Empty parameter should return first error from is valid error", func() {
				appError := psi.Update(&empty, &newParameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_ENG"
			Convey("Given empty local or too long local should return Local error", func() {
				appError := psi.Update(&parameter, &newParameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
				parameter.Local = ""
				appError = psi.Update(&parameter, &newParameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_EN"
			parameter.TimeZone = "UTF+134"
			Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
				appError := psi.Update(&parameter, &newParameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
				parameter.TimeZone = ""
				appError = psi.Update(&parameter, &newParameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
			})
			parameter.TimeZone = "UTF+12"
			parameter.SleepEnd = -1
			Convey("Given negative or too big Sleep timers should return sleep error", func() {
				appError := psi.Update(&parameter, &newParameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 1441
				appError = psi.Update(&parameter, &newParameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 10
				parameter.SleepStart = -10
				appError = psi.Update(&parameter, &newParameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
				parameter.SleepStart = 2000
				appError = psi.Update(&parameter, &newParameter, db)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
			})
		})
		db.Delete(&parameter)
		db.Delete(&newParameter)
	})

	Convey("Testing Get function", t, func() {
		parameter := Parameter{
			Local:      "vi_VI",
			TimeZone:   "UTC+6",
			SleepStart: 260,
			SleepEnd:   24,
		}
		Convey("Trying to get parameter from empty DB should return empty", func() {
			So(psi.Get(db), ShouldResemble, Parameter{})
		})
		appError := psi.Save(&parameter, db)
		So(appError, ShouldBeNil)
		Convey("Trying to get parameter from non empty DB should return a correct parameter object", func() {
			got := psi.Get(db)
			So(got, ShouldResemble, parameter)
			So(got.IsValid(), ShouldBeNil)
		})
		db.Delete(&parameter)
	})
}
