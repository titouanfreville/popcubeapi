package models

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"

	u "github.com/titouanfreville/popcubeapi/utils"
)

const (
	organisationDisplayNameMaxRunes = 64
	organisationNameMaxLength       = 64
	organisationDescriptionMaxRunes = 1024
	organisationSubjectMaxRunes     = 250
)

// Organisation object
//
// Describe organisation you are in. It is an unique object in the database.
//
// swagger:model
type Organisation struct {
	// id of the organisation
	//
	// min: 0
	IDOrganisation uint64 `gorm:"primary_key;column:idOrganisation;AUTO_INCREMENT" json:"id,omitempty"`
	// Stack into docker swarm
	//
	// required: true
	//min: 0
	DockerStack int `gorm:"column:dockerStack;not null;unique" json:"docker_stack,omitempty"`
	// required: true
	OrganisationName string `gorm:"column:organisationName;not null;unique" json:"name,omitempty"`
	// State if organisation is free to join or not. Default is private (false).
	Public      bool   `gorm:"column:public; not null" json:"public"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	Avatar      string `gorm:"column:avatar" json:"avatar,omitempty"`
	// Domain name of the organisation
	Domain string `gorm:"column:domain" json:"domain,omitempty"`
}

// ToJSON transfoorm an Organisation into JSON
func (organisation *Organisation) ToJSON() string {
	b, err := json.Marshal(organisation)
	if err != nil {
		return ""
	}
	return string(b)
}

// OganisationFromJSON Try to parse a json object as emoji
func OganisationFromJSON(data io.Reader) *Organisation {
	decoder := json.NewDecoder(data)
	var organisation Organisation
	err := decoder.Decode(&organisation)
	if err == nil {
		return &organisation
	}
	return nil
}

// IsValid is used to check validity of Organisation objects
func (organisation *Organisation) IsValid() *u.AppError {

	if len(organisation.OrganisationName) == 0 || utf8.RuneCountInString(organisation.OrganisationName) > organisationDisplayNameMaxRunes {
		return u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10))
	}

	if !IsValidOrganisationIdentifier(organisation.OrganisationName) {
		return u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10))
	}

	if utf8.RuneCountInString(organisation.Description) > organisationDescriptionMaxRunes {
		return u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.description.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10))
	}

	return nil
}

// PreSave is used to add some default values to organisation before saving in DB (creation).
func (organisation *Organisation) PreSave() {
	organisation.OrganisationName = strings.ToLower(organisation.OrganisationName)

	if organisation.Avatar == "" {
		organisation.Avatar = "default_organisation_avatar.svg"
	}
}
