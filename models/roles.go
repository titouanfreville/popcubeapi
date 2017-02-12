package models

import (
	"encoding/json"
	"io"
	"regexp"
	u "github.com/titouanfreville/popcubeapi/utils"
)

// Role Type descibe the Role table for Popcube DB
type Role struct {
	IDRole        uint64 `gorm:"primary_key;column:idRole;AUTO_INCREMENT" json:"-"`
	RoleName      string `gorm:"column:roleName;unique_index;not null;unique" json:"name"`
	CanUsePrivate bool   `gorm:"column:canUsePrivate;not null" json:"canUsePrivate"`
	CanModerate   bool   `gorm:"column:canModerate;not null" json:"canModerate"`
	CanArchive    bool   `gorm:"column:canArchive;not null" json:"canArchive"`
	CanInvite     bool   `gorm:"column:canInvite;not null" json:"canInvite"`
	CanManage     bool   `gorm:"column:canManage;not null" json:"canManage"`
	CanManageUser bool   `gorm:"column:canManageUser;not null" json:"canManageUser"`
}

var (
	// Owner is one of the defaul roles. Have all rights.
	Owner = Role{
		RoleName:      "owner",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     true,
		CanManage:     true,
		CanManageUser: true,
	}
	// Admin is one of the defaul roles. Have all rights.
	Admin = Role{
		RoleName:      "admin",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     true,
		CanManage:     true,
		CanManageUser: true,
	}
	// Standart is one of the defaul roles. Have all channel rights.
	Standart = Role{
		RoleName:      "standart",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	// Guest is one of the defaul roles. Have no rights.
	Guest = Role{
		RoleName:      "guest",
		CanUsePrivate: false,
		CanModerate:   false,
		CanArchive:    false,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	// BasicsRoles defines the list of basics roles
	BasicsRoles = []*Role{
		&Owner,
		&Admin,
		&Standart,
		&Guest,
	}
	restrictedRoleNames = []string{
		"owner",
		"admin",
		"standart",
		"guest",
	}
	validRoleNameChars = regexp.MustCompile(`^[a-z]+$`)
)

// IsValid is used to check validity of Role objects
func (role *Role) IsValid() *u.AppError {
	if !IsValidRoleName(role.RoleName) {
		return u.NewLocAppError("Role.IsValid", "model.role.rolename.app_error", nil, "")
	}

	return nil
}

// ToJSON transfoorm an Role into JSON
func (role *Role) ToJSON() string {
	b, err := json.Marshal(role)
	if err != nil {
		return ""
	}
	return string(b)
}

// IsValidRoleName Check that provided string is correctly formed to be used as a RoleName
func IsValidRoleName(u string) bool {
	if len(u) == 0 || len(u) > 64 {
		return false
	}

	if !validRoleNameChars.MatchString(u) {
		return false
	}

	for _, restrictedRoleName := range restrictedRoleNames {
		if u == restrictedRoleName {
			return false
		}
	}

	return true
}

// RoleFromJSON Try to parse a json object as emoji
func RoleFromJSON(data io.Reader) *Role {
	decoder := json.NewDecoder(data)
	var role Role
	err := decoder.Decode(&role)
	if err == nil {
		return &role
	}
	return nil

}

// RoleListToJSON Convert a list of roles into a JSON array
func RoleListToJSON(roleList []*Role) string {
	b, err := json.Marshal(roleList)
	if err != nil {
		return ""
	}
	return string(b)

}

// RoleListFromJSON Try to parse a JSON array into a role list
func RoleListFromJSON(data io.Reader) []*Role {
	decoder := json.NewDecoder(data)
	var roleList []*Role
	err := decoder.Decode(&roleList)
	if err == nil {
		return roleList
	}
	return nil
}
