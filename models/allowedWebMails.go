package models

import (
	"github.com/titouanfreville/popcubeapi/utils"
)

const (
	allowedWebMailsDisplayNameMaxRunes = 64
	allowedWebMailsNameMaxLength       = 64
	allowedWebMailsDescriptionMaxRunes = 1024
	allowedWebMailsSubjectMaxRunes     = 250
)

// AllowedWebMails object
//
// Describe web mails user that can join organisation without being invited. (ex: @popcube.xyz, @supinfo.com, etc.)
//
// swagger:model
type AllowedWebMails struct {
	// id of the allowedWebMails
	//
	// min: 0
	IDAllowedWebMails uint64 `gorm:"primary_key;column:idAllowedWebMails;AUTO_INCREMENT" json:"id,omitempty"`
	// Mail provider
	// required: true
	Provider string `gorm:"column:provider;not null;" json:"provider,omitempty"`
	// Domain name of the allowedWebMails
	Domain string `gorm:"column:domain" json:"domain,omitempty"`
	// Rights that will automatically atributed to user created accound from the webmail
	DefaultRights string `gorm:"column:description" json:"description,omitempty"`
}

// IsValid check validity of object before saving in DB
func (aw AllowedWebMails) IsValid() *utils.AppError {
	if aw.Domain == "" {
		return utils.NewLocAppError("allowedWebMails.isValid", "domain undefined", nil, "")
	}
	return nil
}
