package models

import (
	u "github.com/titouanfreville/popcubeapi/utils"
)

// Member object.
//
// Member is the link between an User and a Channel. It also state the role of the user
// in the channel if it is channel specific.
//
// swagger:model
type Member struct {
	User User `db:"-" json:"-"`
	// required: true
	IDUser  uint64  `gorm:"column:idUser; not null;" json:"id_user,omitempty"`
	Channel Channel `db:"-" json:"-"`
	// required: true
	IDChannel  uint64 `gorm:"column:idChannel; not null;" json:"id_channel,omitempty"`
	Role       Role   `db:"-" json:"-"`
	IDRole     uint64 `gorm:"column:idRole; not null;" json:"id_role,omitempty"`
	TimedOut   bool   `gorm:"column:timedOut; not null;" json:"timed_out, omitempty"`
	TimeOutEnd int64  `gorm:"column:timeOutEnd" json:"timeout_end, omitempty"`
}

// IsValid check validity of member object
func (member *Member) IsValid() *u.AppError {
	// if member.User == (User{}) {
	// 	return u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, "")
	// }
	// if member.Channel == (Channel{}) {
	// 	return u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, "")
	// }
	return nil
}
