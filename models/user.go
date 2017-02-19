/*Package models implements the basics databases models used by PopCube chat api.

Models

The following is a list of models described:
	Avatar: Contain all informations for avatar management
	Channel: Contain all informations for channel management
	Emojis: Contain all informations for emojis management
	Organisation: Contain all informations for organisation management
	Parameter: Contain all informations for parmeters management
	Role: Contain all informations for roles management
	User: Contain all informations for users management
*/
// Created by Titouan FREVILLE <titouanfreville@gmail.com>
//
// Inspired by mattermost project
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

/*User object

- webwebID: String unique and non null to webIDentify the user on application services. - REQUIRED

- username: Store the user username used to log into the service. - REQUIRED

- email: user mail ;). - REQUIRED

- emailVerified: true if email was verified by user. - REQUIRED

- lastUpdate: Time of the last update. Used to create tag for browser cache. - REQUIRED

- deleted: True if user is deleted. - REQUIRED

- password: Hashed password. - REQUIRED

- lastpasswordUpdate: Date of the last password modification. - REQUIRED

- failedAttemps: Number of fail try to connect to account. - REQUIRED

- locale: user favorite langage. - REQUIRED

- role : int referencing a user role existing in the database. - REQUIRED

- nickname: Name to use in communication channel (by default : username).

- first name: user true first name.

- last name: user true last name.

- lastActivityAt: Date && Time of the last activity of the user.
*/
type User struct {
	IDUser             uint64 `gorm:"primary_key;column:idUser;AUTO_INCREMENT" json:"id,omitempty"`
	WebID              string `gorm:"column:webId; not null; unique;" json:"web_id,omitempty"`
	Username           string `gorm:"column:userName; not null; unique;" json:"username,omitempty"`
	Email              string `gorm:"column:email; not null; unique;" json:"email,omitempty"`
	EmailVerified      bool   `gorm:"column:emailVerified; not null;" json:"email_verified,omitempty"`
	LastUpdate         int64  `gorm:"column:lastUpdate; not null;" json:"last_update,omitempty"`
	Deleted            bool   `gorm:"column:deleted; not null;" json:"deleted,omitempty"`
	Password           string `gorm:"column:password; not null;" json:"password,omitempty"`
	LastPasswordUpdate int64  `gorm:"column:lastPasswordUpdate; not null;" json:"last_password_update,omitempty"`
	FailedAttempts     int    `gorm:"column:failedAttempts; not null;" json:"failed_attempts,omitempty"`
	Locale             string `gorm:"column:locale; not null;" json:"locale,omitempty"`
	Role               Role   `gorm:"ForeignKey:IDRole;" db:"-" json:"-"`
	IDRole             uint64 `gorm:"column:idRole; not null;" json:"id_role,omitempty"`
	Avatar             string `gorm:"column:avatar;" json:"avatar, omitempty"`
	NickName           string `gorm:"column:nickName; unique" json:"nickname, omitempty"`
	FirstName          string `gorm:"column:firstName;" json:"first_name, omitempty"`
	LastName           string `gorm:"column:lastName;" json:"last_name, omitempty"`
	LastActivityAt     int64  `db:"-" json:"last_activity_at,omitempty"`
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

	if user.Locale == "" {
		user.Locale = DefaultLocale
	}

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
