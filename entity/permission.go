package entity

import (
	"gorm.io/gorm"
)

type ChannelPermission struct {
	gorm.Model
	ChannelID string `gorm:"type:varchar(255);not null; index:channel_idx,unique" `
	Active    bool
}
type UserPermission struct {
	gorm.Model
	UserID string `gorm:"type:varchar(255);not null; index:user_idx,unique" `
	Active bool
}
