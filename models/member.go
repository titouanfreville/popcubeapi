package models

import (
	u "github.com/titouanfreville/popcubeapi/utils"
)

// Member describe the associtive table member between USER, CHANNEL, and ROLE
type Member struct {
	// IDMember uint64  `gorm:"primary_key;column:idMember;AUTO_INCREMENT" json:"-"`
	User      User    `db:"-" json:"-"`
	IDUser    uint64  `gorm:"column:idUser; not null;" json:"id_user"`
	Channel   Channel `db:"-" json:"-"`
	IDChannel uint64  `gorm:"column:idChannel; not null;" json:"id_channel"`
	Role      Role    `db:"-" json:"-"`
	IDRole    uint64  `gorm:"column:idRole; not null;" json:"id_role"`
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
