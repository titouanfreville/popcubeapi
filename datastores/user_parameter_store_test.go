// This file is used to test if user model is working correctly.
// A user is always linked to an userParameter
// He has bpsic channel to join
package datastores

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func genUPlocalError(locale string) *u.AppError {
	return u.NewLocAppError("userParameterStoreImpl.Save", "too long local", nil, "The local : "+locale+" can not be manage. Max size for local is 5.")
}

func genUPtimeZoneError(tz string) *u.AppError {
	return u.NewLocAppError("userParameterStoreImpl.Save", "too long timeZone", nil, "The TimeZone : "+tz+" can not be manage. Max size for local is 4.")
}

func genUPsleepError(part string, time int) *u.AppError {
	return u.NewLocAppError("userParameterStoreImpl.Save", "invalid hour", nil, "The sleep "+part+" time: "+strconv.Itoa(time)+"ms is not valable. It has to be between 0 and 1440.")
}

func TestUserParameterStore(t *testing.T) {
	store := Store()
	db := store.InitConnection("root", "popcube_test", "popcube_dev", "database", "3306")
	psi := store.UserParameter()
	usi := store.User()
	rsi := store.Role()

	standartRole := Role{
		RoleName:      randStringBytes(10),
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	db.Delete(&standartRole)
	rsi.Save(&standartRole, db)

	userTest := User{
		Username:  randStringBytes(10),
		Password:  "test",
		Email:     "test@popcube.fr",
		NickName:  "NickName",
		FirstName: "Test",
		LastName:  "L",
		IDRole:    standartRole.IDRole,
	}
	db.Delete(&userTest)
	usi.Save(&userTest, db)
	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("userParameterStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyexistError := u.NewLocAppError("userParameterStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
		userParameter := UserParameter{
			IDUser:        userTest.IDUser,
			ParameterName: "test",
			Local:         "en_EN",
			TimeZone:      "UTC+2",
			SleepStart:    280,
			SleepEnd:      12,
		}
		db.Delete(userParameter)
		Convey("Given a correct userParameter.", func() {
			appError := psi.Save(&userParameter, db)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
			})
			// Convey("Trying to add it a second time should return duplicate error", func() {
			// 	appError2 := psi.Save(&userParameter, db)
			// 	So(appError2, ShouldNotBeNil)
			// 	So(appError2, ShouldResemble, alreadyexistError)
			// 	So(appError2, ShouldNotResemble, dbError)
			// })
		})
		Convey("Given an incorrect userParameter.", func() {
			userParameter.Local = "en_ENGA"
			Convey("Given empty local or too long local should return Local error", func() {
				appError := psi.Save(&userParameter, db)
				So(appError, ShouldResemble, genUPlocalError(userParameter.Local))
			})
			userParameter.Local = "en_EN"
			userParameter.TimeZone = "UTF+134"
			Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
				appError := psi.Save(&userParameter, db)
				So(appError, ShouldResemble, genUPtimeZoneError(userParameter.TimeZone))
			})
			userParameter.TimeZone = "UTF+12"
			userParameter.SleepEnd = -1
			Convey("Given negative or too big Sleep timers should return sleep error", func() {
				appError := psi.Save(&userParameter, db)
				So(appError, ShouldResemble, genUPsleepError("end", userParameter.SleepEnd))
				userParameter.SleepEnd = 1441
				appError = psi.Save(&userParameter, db)
				So(appError, ShouldResemble, genUPsleepError("end", userParameter.SleepEnd))
				userParameter.SleepEnd = 10
				userParameter.SleepStart = -10
				appError = psi.Save(&userParameter, db)
				So(appError, ShouldResemble, genUPsleepError("start", userParameter.SleepStart))
				userParameter.SleepStart = 2000
				appError = psi.Save(&userParameter, db)
				So(appError, ShouldResemble, genUPsleepError("start", userParameter.SleepStart))
			})
		})
		db.Delete(&userParameter)
	})

	Convey("Testing update function", t, func() {
		userParameter := UserParameter{
			IDUser:        userTest.IDUser,
			ParameterName: "test",
			Local:         "en_EN",
			TimeZone:      "UTC+2",
			SleepStart:    280,
			SleepEnd:      12,
		}
		newUserParameter := UserParameter{
			IDUser:        userTest.IDUser,
			ParameterName: "test",
			Local:         "vn_VN",
			TimeZone:      "UTC+10",
			SleepStart:    281,
			SleepEnd:      13,
		}
		db.Delete(userParameter)
		appError := psi.Save(&userParameter, db)
		dbError := u.NewLocAppError("userParameterStoreImpl.Update", "update.transaction.updates.encounterError", nil, "")
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		Convey("Providing a correct user to update", func() {
			appError := psi.Update(&userParameter, &newUserParameter, db)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
		})
		// Convey("Providing an incorrect user as new should result in errors", func() {
		// 	empty := UserParameter{}
		// 	Convey("Empty userParameter should return first error from is valid error", func() {
		// 		appError := psi.Update(&userParameter, &empty, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", "model.userParameter.is_valid.userParameter_local.app_error", nil, ""))
		// 	})
		// 	newUserParameter.Local = "en_ENG"
		// 	Convey("Given empty local or too long local should return Local error", func() {
		// 		appError := psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", "model.userParameter.is_valid.userParameter_local.app_error", nil, ""))
		// 		newUserParameter.Local = ""
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", "model.userParameter.is_valid.userParameter_local.app_error", nil, ""))
		// 	})
		// 	newUserParameter.Local = "en_EN"
		// 	newUserParameter.TimeZone = "UTF+134"
		// 	Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
		// 		appError := psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", "model.userParameter.is_valid.userParameter_timezone.app_error", nil, ""))
		// 		newUserParameter.TimeZone = ""
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", "model.userParameter.is_valid.userParameter_timezone.app_error", nil, ""))
		// 	})
		// 	newUserParameter.TimeZone = "UTF+12"
		// 	newUserParameter.SleepEnd = -1
		// 	Convey("Given negative or too big Sleep timers should return sleep error", func() {
		// 		appError := psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", "model.userParameter.is_valid.userParameter_sleep_end.app_error", nil, ""))
		// 		newUserParameter.SleepEnd = 1441
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", "model.userParameter.is_valid.userParameter_sleep_end.app_error", nil, ""))
		// 		newUserParameter.SleepEnd = 10
		// 		newUserParameter.SleepStart = -10
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", "model.userParameter.is_valid.userParameter_sleep_start.app_error", nil, ""))
		// 		newUserParameter.SleepStart = 2000
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterNew.PreSave", "model.userParameter.is_valid.userParameter_sleep_start.app_error", nil, ""))
		// 	})
		// })

		// Convey("Providing an incorrect user as old should result in errors", func() {
		// 	empty := UserParameter{}
		// 	Convey("Empty userParameter should return first error from is valid error", func() {
		// 		appError := psi.Update(&empty, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", "model.userParameter.is_valid.userParameter_local.app_error", nil, ""))
		// 	})
		// 	userParameter.Local = "en_ENG"
		// 	Convey("Given empty local or too long local should return Local error", func() {
		// 		appError := psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", "model.userParameter.is_valid.userParameter_local.app_error", nil, ""))
		// 		userParameter.Local = ""
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", "model.userParameter.is_valid.userParameter_local.app_error", nil, ""))
		// 	})
		// 	userParameter.Local = "en_EN"
		// 	userParameter.TimeZone = "UTF+134"
		// 	Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
		// 		appError := psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", "model.userParameter.is_valid.userParameter_timezone.app_error", nil, ""))
		// 		userParameter.TimeZone = ""
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", "model.userParameter.is_valid.userParameter_timezone.app_error", nil, ""))
		// 	})
		// 	userParameter.TimeZone = "UTF+12"
		// 	userParameter.SleepEnd = -1
		// 	Convey("Given negative or too big Sleep timers should return sleep error", func() {
		// 		appError := psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", "model.userParameter.is_valid.userParameter_sleep_end.app_error", nil, ""))
		// 		userParameter.SleepEnd = 1441
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", "model.userParameter.is_valid.userParameter_sleep_end.app_error", nil, ""))
		// 		userParameter.SleepEnd = 10
		// 		userParameter.SleepStart = -10
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", "model.userParameter.is_valid.userParameter_sleep_start.app_error", nil, ""))
		// 		userParameter.SleepStart = 2000
		// 		appError = psi.Update(&userParameter, &newUserParameter, db)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userParameterStoreImpl.Update.userParameterOld.PreSave", "model.userParameter.is_valid.userParameter_sleep_start.app_error", nil, ""))
		// 	})
		// })
		db.Delete(&userParameter)
		db.Delete(&newUserParameter)
	})

	// Convey("Testing Get function", t, func() {
	// 	userParameter := UserParameter{
	// 		Local:      "vi_VI",
	// 		TimeZone:   "UTC+6",
	// 		SleepStart: 260,
	// 		SleepEnd:   24,
	// 	}
	// 	Convey("Trying to get userParameter from empty DB should return empty", func() {
	// 		So(psi.GetAll(db), ShouldResemble, UserParameter{})
	// 	})
	// 	appError := psi.Save(&userParameter, db)
	// 	So(appError, ShouldBeNil)
	// 	Convey("Trying to get userParameter from non empty DB should return a correct userParameter object", func() {
	// 		got := psi.GetAll(db)
	// 		So(got, ShouldResemble, userParameter)
	// 		// So(got.IsValid(), ShouldBeNil)
	// 	})
	// 	db.Delete(&userParameter)
	// })

	db.Delete(&userTest)
	db.Delete(&standartRole)
}
