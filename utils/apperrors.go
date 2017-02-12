package utils

import (
	"encoding/json"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
	"io"
	"io/ioutil"
	"strings"
)

// AppError Type used to structure error reporting for popcube chat project.
type AppError struct {
	ID            string                 `json:"id"`
	Message       string                 `json:"message"`               // Message to be display to the end user without debugging information
	DetailedError string                 `json:"detailed_error"`        // Internal error string to help the developer
	RequestID     string                 `json:"request_id,omitempty"`  // The RequestID that's also set in the header
	StatusCode    int                    `json:"status_code,omitempty"` // The http status code
	Where         string                 `json:"-"`                     // The function where it happened in the form of Struct.Func
	IsOAuth       bool                   `json:"is_oauth,omitempty"`    // Whether the error is OAuth specific
	params        map[string]interface{} `json:"-"`
}

// Error return a string for AppError Type
func (er *AppError) Error() string {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}

func (er *AppError) translate(T goi18n.TranslateFunc) {
	if er.params == nil {
		er.Message = T(er.ID)
	}
	er.Message = T(er.ID, er.params)

}

func (er *AppError) systemMessage(T goi18n.TranslateFunc) string {
	if er.params == nil {
		return T(er.ID)
	}
	return T(er.ID, er.params)
}

// ToJSON function to transform AppError
func (er *AppError) ToJSON() string {
	b, err := json.Marshal(er)
	if err != nil {
		return ""
	}
	return string(b)
}

// AppErrorFromJSON will decode the input and return an AppError
func AppErrorFromJSON(data io.Reader) *AppError {
	str := ""
	bytes, rerr := ioutil.ReadAll(data)
	if rerr != nil {
		str = rerr.Error()
	} else {
		str = string(bytes)
	}

	decoder := json.NewDecoder(strings.NewReader(str))
	var er AppError
	err := decoder.Decode(&er)
	if err == nil {
		return &er
	}
	return NewLocAppError("AppErrorFromJSON", "model.utils.decode_json.app_error", nil, "body: "+str)

}

// NewLocAppError is used to generate server errors
func NewLocAppError(where string, id string, params map[string]interface{}, details string) *AppError {
	ap := &AppError{}
	ap.ID = id
	ap.params = params
	ap.Message = id
	ap.Where = where
	ap.DetailedError = details
	ap.StatusCode = 500
	ap.IsOAuth = false
	return ap
}
