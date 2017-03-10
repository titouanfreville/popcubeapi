package models

import (
	u "github.com/titouanfreville/popcubeapi/utils"
)

// Read object.
//
// Read state who read a given message in a given channel.
//
// swagger:model
type Read struct {
	// required: true
	IDUser uint64 `gorm:"column:idUser; not null;" json:"id_user,omitempty"`
	// required: true
	IDChannel uint64 `gorm:"column:idChannel; not null;" json:"id_channel,omitempty"`
	// required: true
	IDMessage uint64 `gorm:"column:idMessage; not null;" json:"id_message,omitempty"`
}

// IsValid check validity of read object
func (read *Read) IsValid() *u.AppError {
	// if read.User == (User{}) {
	// 	return u.NewLocAppError("Read.IsValid", "model.read.user.app_error", nil, "")
	// }
	// if read.Channel == (Channel{}) {
	// 	return u.NewLocAppError("Read.IsValid", "model.read.channel.app_error", nil, "")
	// }
	return nil
}
