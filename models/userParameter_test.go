package models

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func genUPlocalError(locale string) *u.AppError {
	return u.NewLocAppError("UserParameter.IsValid", "too long local", nil, "The local : "+locale+" can not be manage. Max size for local is 5.")
}

func genUPtimeZoneError(tz string) *u.AppError {
	return u.NewLocAppError("UserParameter.IsValid", "too long timeZone", nil, "The TimeZone : "+tz+" can not be manage. Max size for local is 4.")
}

func genUPsleepError(part string, time int) *u.AppError {
	return u.NewLocAppError("UserParameter.IsValid", "invalid hour", nil, "The sleep "+part+" time: "+strconv.Itoa(time)+"ms is not valable. It has to be between 0 and 1440.")
}

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
		Convey("Given a correct userParameter. UserParameter should be validate", func() {
			So(userParameter.IsValid(), ShouldBeNil)
			So(userParameter.IsValid(), ShouldNotResemble, nameError)
			So(userParameter.IsValid(), ShouldNotResemble, genUPlocalError(userParameter.Local))
			So(userParameter.IsValid(), ShouldNotResemble, genUPtimeZoneError(userParameter.TimeZone))
			So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("start", userParameter.SleepStart))
			So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("end", userParameter.SleepEnd))
		})
		Convey("Given an incorrect userParameter. UserParameter should be refused", func() {
			empty := UserParameter{}
			userParameter.ParameterName = ""
			Convey("Empty userParameter should return first error from is valid error", func() {
				So(empty.IsValid(), ShouldNotBeNil)
				So(empty.IsValid(), ShouldResemble, nameError)
				So(empty.IsValid(), ShouldNotResemble, genUPlocalError(userParameter.Local))
				So(empty.IsValid(), ShouldNotResemble, genUPtimeZoneError(userParameter.TimeZone))
				So(empty.IsValid(), ShouldNotResemble, genUPsleepError("start", userParameter.SleepStart))
				So(empty.IsValid(), ShouldNotResemble, genUPsleepError("end", userParameter.SleepEnd))
			})
			Convey("Empty userParameter name should return name error", func() {
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, genUPlocalError(userParameter.Local))
				So(userParameter.IsValid(), ShouldNotResemble, genUPtimeZoneError(userParameter.TimeZone))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("start", userParameter.SleepStart))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("end", userParameter.SleepEnd))
			})
			userParameter.ParameterName = "test"
			userParameter.Local = "en_ENG"
			Convey("Given too long local should return Local error", func() {
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldResemble, genUPlocalError(userParameter.Local))
				So(userParameter.IsValid(), ShouldNotResemble, genUPtimeZoneError(userParameter.TimeZone))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("start", userParameter.SleepStart))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("end", userParameter.SleepEnd))
			})
			userParameter.Local = "en_EN"
			userParameter.TimeZone = "UTF+134"
			Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, genUPlocalError(userParameter.Local))
				So(userParameter.IsValid(), ShouldResemble, genUPtimeZoneError(userParameter.TimeZone))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("start", userParameter.SleepStart))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("end", userParameter.SleepEnd))
			})
			userParameter.TimeZone = "UTF+12"
			userParameter.SleepStart = -1
			Convey("Given negative or too big Sleep timers should return sleep error", func() {
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, genUPlocalError(userParameter.Local))
				So(userParameter.IsValid(), ShouldNotResemble, genUPtimeZoneError(userParameter.TimeZone))
				So(userParameter.IsValid(), ShouldResemble, genUPsleepError("start", userParameter.SleepStart))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("end", userParameter.SleepEnd))
				userParameter.SleepStart = 1441
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, genUPlocalError(userParameter.Local))
				So(userParameter.IsValid(), ShouldNotResemble, genUPtimeZoneError(userParameter.TimeZone))
				So(userParameter.IsValid(), ShouldResemble, genUPsleepError("start", userParameter.SleepStart))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("end", userParameter.SleepEnd))
				userParameter.SleepStart = 10
				userParameter.SleepEnd = -10
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, genUPlocalError(userParameter.Local))
				So(userParameter.IsValid(), ShouldNotResemble, genUPtimeZoneError(userParameter.TimeZone))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("start", userParameter.SleepStart))
				So(userParameter.IsValid(), ShouldResemble, genUPsleepError("end", userParameter.SleepEnd))
				userParameter.SleepEnd = 2000
				So(userParameter.IsValid(), ShouldNotBeNil)
				So(userParameter.IsValid(), ShouldNotResemble, nameError)
				So(userParameter.IsValid(), ShouldNotResemble, genUPlocalError(userParameter.Local))
				So(userParameter.IsValid(), ShouldNotResemble, genUPtimeZoneError(userParameter.TimeZone))
				So(userParameter.IsValid(), ShouldNotResemble, genUPsleepError("start", userParameter.SleepStart))
				So(userParameter.IsValid(), ShouldResemble, genUPsleepError("end", userParameter.SleepEnd))
			})
		})
	})
}
