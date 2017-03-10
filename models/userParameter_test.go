package models

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestUserParameterModel(t *testing.T) {
	Convey("Testing IsValid function", t, func() {
		userParameter := UserParameter{
			IDUser:        1,
			ParameterName: "test",
			Local:         "en_EN",
			TimeZone:      "UTC+2",
			SleepStart:    280,
			SleepEnd:      12,
		}
		nameError := u.NewLocAppError("UserParameter.IsValid", "parameter name Undefined", nil, "")
		localError := u.NewLocAppError("UserParameter.IsValid", "too long local", nil, "The local :"+userParameter.Local+" can not be manage. Max size for local is 5")
		timeZoneError := u.NewLocAppError("UserParameter.IsValid", "too long timeZone", nil, "The TimeZone :"+userParameter.TimeZone+" can not be manage. Max size for local is 4")
		sleepStartError := u.NewLocAppError("UserParameter.IsValid", "invalid hour", nil, "The sleep start time: "+strconv.Itoa(userParameter.SleepStart)+"ms is not valable. It has to be between 0 and 1440.")
		sleepEndError := u.NewLocAppError("UserParameter.IsValid", "invalid hour", nil, "The sleep end time: "+strconv.Itoa(userParameter.SleepEnd)+"ms is not valable. It has to be between 0 and 1440.")
		Convey("Given a correct userParameter. UserParameter should be validate", func() {
			So(userParameter.IsValid(), ShouldBeNil)
			So(userParameter.IsValid(), ShouldNotResemble, nameError)
			So(userParameter.IsValid(), ShouldNotResemble, localError)
			So(userParameter.IsValid(), ShouldNotResemble, timeZoneError)
			So(userParameter.IsValid(), ShouldNotResemble, sleepStartError)
			So(userParameter.IsValid(), ShouldNotResemble, sleepEndError)
		})
		Convey("Given an incorrect userParameter. UserParameter should be refused", func() {
			empty := UserParameter{}
			userParameter.ParameterName = ""
			Convey("Empty userParameter should return first error from is valid error", func() {
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, localError)
				So(userParameter.IsValid(), ShouldNotResemble, timeZoneError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepStartError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepEndError)
			})
			Convey("Empty userParameter name should return local error", func() {
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, localError)
				So(userParameter.IsValid(), ShouldNotResemble, timeZoneError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepStartError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepEndError)
			})
			userParameter.Local = "en_ENG"
			Convey("Given too long local should return Local error", func() {
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldResemble, localError)
				So(userParameter.IsValid(), ShouldNotResemble, timeZoneError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepStartError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepEndError)
			})
			userParameter.Local = "en_EN"
			userParameter.TimeZone = "UTF+134"
			Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, localError)
				So(userParameter.IsValid(), ShouldResemble, timeZoneError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepStartError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepEndError)
			})
			userParameter.TimeZone = "UTF+12"
			userParameter.SleepEnd = -1
			Convey("Given negative or too big Sleep timers should return sleep error", func() {
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, localError)
				So(userParameter.IsValid(), ShouldNotResemble, timeZoneError)
				So(userParameter.IsValid(), ShouldResemble, sleepStartError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepEndError)
				userParameter.SleepEnd = 1441
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, localError)
				So(userParameter.IsValid(), ShouldNotResemble, timeZoneError)
				So(userParameter.IsValid(), ShouldResemble, sleepStartError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepEndError)
				userParameter.SleepEnd = 10
				userParameter.SleepStart = -10
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, localError)
				So(userParameter.IsValid(), ShouldNotResemble, timeZoneError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepStartError)
				So(userParameter.IsValid(), ShouldResemble, sleepEndError)
				userParameter.SleepStart = 2000
				So(userParameter.IsValid(), ShouldResemble, u.NewLocAppError("UserParameter.IsValid", "model.userParameter.is_valid.userParameter_sleep_start.app_error", nil, ""))
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, localError)
				So(userParameter.IsValid(), ShouldNotResemble, timeZoneError)
				So(userParameter.IsValid(), ShouldNotResemble, sleepStartError)
				So(userParameter.IsValid(), ShouldResemble, sleepEndError)
			})
		})
	})
}
