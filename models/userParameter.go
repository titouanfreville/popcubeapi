package models

import (
	"strconv"

	u "github.com/titouanfreville/popcubeapi/utils"
)

// Defined in Parameter.
// const (
// 	localMaxSize    = 5
// 	timeZoneMaxSize = 6
// 	maxTime         = 1440
// )

// UserParameter object
//
// User parameter store all the parameters set an user can use.
//
// swagger:model
type UserParameter struct {
	// id of the user who parameter can be applied to
	IDUser uint64 `gorm:"column:idUser; not null" json:"id,omitempty"`
	// required true
	ParameterName string `gorm:"column: parameterName" json:"parameter_name, omitempty"`
	// Default langage
	//
	// required: true
	Local string `gorm:"column:local;not null; unique" json:"local,omitempty"`
	// Default time zone
	//
	// required: true
	TimeZone string `gorm:"column:timeZone;not null; unique;" json:"time_zone,omitempty"`
	// Default start of non notification period
	//
	// required: true
	SleepStart int `gorm:"column:sleepStart;not null;unique" json:"sleep_start,omitempty"`
	// Default end of non notification period
	//
	// required: true
	SleepEnd int `gorm:"column:sleepEnd;not null;unique" json:"sleep_end,omitempty"`
}

// IsValid is used to check validity of UserParameter objects
func (userParameter *UserParameter) IsValid() *u.AppError {
	if userParameter.ParameterName == "" {
		return u.NewLocAppError("UserParameter.IsValid", "parameter name Undefined", nil, "")
	}

	if len(userParameter.Local) > localMaxSize {
		return u.NewLocAppError("UserParameter.IsValid", "too long local", nil, "The local : "+userParameter.Local+" can not be manage. Max size for local is 5.")
	}

	if len(userParameter.TimeZone) > timeZoneMaxSize {
		return u.NewLocAppError("UserParameter.IsValid", "too long timeZone", nil, "The TimeZone : "+userParameter.TimeZone+" can not be manage. Max size for local is 4.")
	}

	if userParameter.SleepStart < 0 || userParameter.SleepStart > maxTime {
		return u.NewLocAppError("UserParameter.IsValid", "invalid hour", nil, "The sleep start time: "+strconv.Itoa(userParameter.SleepStart)+"ms is not valable. It has to be between 0 and 1440.")
	}

	if userParameter.SleepEnd < 0 || userParameter.SleepEnd > maxTime {
		return u.NewLocAppError("UserParameter.IsValid", "invalid hour", nil, "The sleep end time: "+strconv.Itoa(userParameter.SleepEnd)+"ms is not valable. It has to be between 0 and 1440.")
	}

	return nil
}
