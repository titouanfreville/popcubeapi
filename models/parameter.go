package models

import (
	"encoding/json"
	"io"
	"net/http"

	u "github.com/titouanfreville/popcubeapi/utils"
)

var (
	// EmptyParameter em var param
	EmptyParameter = Parameter{}
)

const (
	// DefaultLocale is a string to describe the default language used in the app
	DefaultLocale = "fr_FR"
	// DefaultTimeZone is a string to describe the default time zone used in the app
	DefaultTimeZone = "UTC-0"
	localMaxSize    = 5
	timeZoneMaxSize = 6
	maxTime         = 1440
)

// Parameter object
//
// Global parameters to apply within organisation. unique object in database
//
// swagger:model
type Parameter struct {
	// id of the parameter
	//
	// min: 0
	IDParameter uint64 `gorm:"primary_key;column:idParameter;AUTO_INCREMENT" json:"id,omitempty"`
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

// Bind method used in API
func (parameter *Parameter) Bind(r *http.Request) error {
	return nil
}

// ToJSON transfoorm an Parameter into JSON
func (parameter *Parameter) ToJSON() string {
	b, err := json.Marshal(parameter)
	if err != nil {
		return ""
	}
	return string(b)
}

// ParameterFromJSON Try to parse a json object as emoji
func ParameterFromJSON(data io.Reader) *Parameter {
	decoder := json.NewDecoder(data)
	var parameter Parameter
	err := decoder.Decode(&parameter)
	if err == nil {
		return &parameter
	}
	return nil
}

// IsValid is used to check validity of Parameter objects
func (parameter *Parameter) IsValid() *u.AppError {

	if len(parameter.Local) == 0 || len(parameter.Local) > localMaxSize {
		return u.NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_local.app_error", nil, "")
	}

	if len(parameter.TimeZone) == 0 || len(parameter.TimeZone) > timeZoneMaxSize {
		return u.NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_timezone.app_error", nil, "")
	}

	if parameter.SleepStart < 0 || parameter.SleepStart > maxTime {
		return u.NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, "")
	}

	if parameter.SleepEnd < 0 || parameter.SleepEnd > maxTime {
		return u.NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, "")
	}

	return nil
}

// PreSave is to be used before saving to add default value if needed
func (parameter *Parameter) PreSave() {
	if parameter.Local == "" {
		parameter.Local = DefaultLocale
	}
	if parameter.TimeZone == "" {
		parameter.TimeZone = DefaultTimeZone
	}
}
