package models

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"

	u "github.com/titouanfreville/popcubeapi/utils"
	"golang.org/x/crypto/bcrypt"
)

const (
	userNotifyAll           = "all"
	userNotifyMention       = "mention"
	userNotifyNone          = "none"
	userAuthServiceEmail    = "email"
	userAuthServiceUsername = "username"
)

var (
	userChannel = []string{"general", "random"}
	// Protected user name cause they are taken by system or used for special mentions.
	restrictedUsernames = []string{
		"all",
		"channel",
		"popcubebot",
		"here",
	}
	// Definition of character user can possess in there names.
	validUsernameChars = regexp.MustCompile(`^[a-z0-9\.\-_]+$`)
)

// User object.
//
// An user is an account who have an access to a specific organisation. Each user is unique inside a given organisation, but users are not shared between
// organisations. Required apply only for creation of the object.
//
// swagger:model
type User struct {
	// id of the user
	//
	// min: 0
	IDUser uint64 `gorm:"primary_key;column:idUser;AUTO_INCREMENT" json:"id,omitempty"`
	// web id for the user used only for cache and cookie purpose
	//
	// required: false
	WebID string `gorm:"column:webId; not null; unique;" json:"web_id,omitempty"`
	// User name
	//
	// required: true
	// max length: 64
	Username string `gorm:"column:userName; not null; unique;" json:"username,omitempty"`
	// User email
	//
	// required: true
	// max lenght: 128
	Email string `gorm:"column:email; not null; unique;" json:"email,omitempty"`
	// State if email was verified
	//
	// required: true
	EmailVerified bool `gorm:"column:emailVerified; not null;" json:"email_verified,omitempty"`
	// Date of the last update from user
	LastUpdate int64 `gorm:"column:lastUpdate; not null;" json:"last_update,omitempty"`
	// User is deleted from organisation but still in database
	//
	// required: true
	Deleted bool `gorm:"column:deleted; not null;" json:"deleted,omitempty"`
	// Encrypted user password
	//
	// required: true
	Password string `gorm:"column:password; not null;" json:"password,omitempty"`
	// Date of the last update of password from user
	//
	// required: true
	LastPasswordUpdate int64 `gorm:"column:lastPasswordUpdate; not null;" json:"last_password_update,omitempty"`
	// Number of attemps failed while loging in
	//
	// required: true
	FailedAttempts int  `gorm:"column:failedAttempts; not null;" json:"failed_attempts,omitempty"`
	Role           Role `gorm:"ForeignKey:IDRole;" db:"-" json:"-"`
	// Role key of user in the organisation
	//
	// required: true
	IDRole uint64 `gorm:"column:idRole; not null;" json:"id_role,omitempty"`
	// Avatar used by user
	Avatar string `gorm:"column:avatar;" json:"avatar, omitempty"`
	// User nickname
	NickName string `gorm:"column:nickName; unique" json:"nickname, omitempty"`
	// First name
	FirstName string `gorm:"column:firstName;" json:"first_name, omitempty"`
	// User Lastname
	LastName       string `gorm:"column:lastName;" json:"last_name, omitempty"`
	LastActivityAt int64  `gorm:"-" db:"-" json:"last_activity_at, omitempty"`
}

// IsValid valwebIDates the user and returns an error if it isn't configured
// correctly.
func (user *User) IsValid(isUpdate bool) *u.AppError {
	if !isUpdate {
		if len(user.WebID) != 26 {
			return u.NewLocAppError("user.IsValid", "model.user.is_valid.WebID.app_error", nil, "")
		}

		if len(user.Email) == 0 {
			return u.NewLocAppError("user.IsValid", "model.user.is_valid.Email.app_error", nil, "user_webID="+user.WebID)
		}

		if len(user.Password) == 0 {
			return u.NewLocAppError("user.IsValid", "model.user.is_valid.auth_data_pwd.app_error", nil, "user_webID="+user.WebID)
		}
	}

	if !IsValidUsername(user.Username) {
		return u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user.WebID)
	}

	if len(user.Email) > 128 || !IsValidEmail(user.Email) {
		return u.NewLocAppError("user.IsValid", "model.user.is_valid.Email.app_error", nil, "user_webID="+user.WebID)
	}

	if utf8.RuneCountInString(user.NickName) > 64 {
		return u.NewLocAppError("user.IsValid", "model.user.is_valid.NickName.app_error", nil, "user_webID="+user.WebID)
	}

	if utf8.RuneCountInString(user.FirstName) > 64 {
		return u.NewLocAppError("user.IsValid", "model.user.is_valid.first_name.app_error", nil, "user_webID="+user.WebID)
	}

	if utf8.RuneCountInString(user.LastName) > 64 {
		return u.NewLocAppError("user.IsValid", "model.user.is_valid.last_name.app_error", nil, "user_webID="+user.WebID)
	}

	return nil
}

// PreSave have to be run before saving user in DB. It will fill necessary information (webID, username, etc. ) and hash password
func (user *User) PreSave() {
	if user.WebID == "" {
		user.WebID = NewID()
	}

	if user.Username == "" {
		user.Username = NewID()
	}

	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	user.LastUpdate = GetMillis()
	user.LastPasswordUpdate = user.LastUpdate

	if len(user.Password) > 0 {
		user.Password = HashPassword(user.Password)
	}
}

// PreUpdate should be run before updating the user in the db.
func (user *User) PreUpdate() {
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)
	user.LastUpdate = GetMillis()

	if len(user.Password) > 0 {
		user.Password = HashPassword(user.Password)
		user.LastPasswordUpdate = user.LastUpdate
	}
}

// ToJSON convert a user to a json string
func (user *User) ToJSON() string {
	b, err := json.Marshal(user)
	if err != nil {
		return ""
	}
	return string(b)
}

// UserFromJSON will decode the input and return a user
func UserFromJSON(data io.Reader) *User {
	decoder := json.NewDecoder(data)
	var user User
	err := decoder.Decode(&user)
	if err == nil {
		return &user
	}
	return nil
}

// IsValidUsername will check if provided userName is correct
func IsValidUsername(user string) bool {
	if len(user) == 0 || len(user) > 64 {
		return false
	}

	if !validUsernameChars.MatchString(user) {
		return false
	}

	for _, restrictedUsername := range restrictedUsernames {
		if user == restrictedUsername {
			return false
		}
	}

	return true
}

// Etag Generate a valwebID strong Etag so the browser can cache the results
func (user *User) Etag(showFullName, showemail bool) string {
	return Etag(user.WebID, user.LastUpdate, showFullName, showemail)
}

// GetFullName of the user
func (user *User) GetFullName() string {
	if user.LastName == "" {
		return user.FirstName
	}
	if user.FirstName == "" {
		return user.LastName
	}
	return user.FirstName + " " + user.LastName
}

// GetDisplayName of the user
func (user *User) GetDisplayName() string {
	if user.NickName != "" {
		return user.NickName
	}
	if user.GetFullName() != "" {
		return user.GetFullName()
	}
	return user.Username
}

// HashPassword generates a hash using the bcrypt.GenerateFrompassword
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

// ComparePassword compares the hash
func ComparePassword(hash string, password string) bool {

	if len(password) == 0 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CleanUsername Transform user name to meet requirement
func CleanUsername(s string) string {
	s = strings.ToLower(strings.Replace(s, " ", "-", -1))

	for _, value := range reservedName {
		if s == value {
			s = strings.Replace(s, value, "", -1)
		}
	}

	s = strings.TrimSpace(s)

	for _, c := range s {
		char := fmt.Sprintf("%c", c)
		if !validUsernameChars.MatchString(char) {
			s = strings.Replace(s, char, "-", -1)
		}
	}

	if !IsValidUsername(s) {
		s = "a" + NewID()
	}

	return s
}
