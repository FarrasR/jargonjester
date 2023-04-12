package entity

import (
	"time"

	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	ChannelID   string    `gorm:"type:varchar(255);not null; index:channel_message_idx" `
	Username    string    `gorm:"type:varchar(255);not null"`
	Content     string    `gorm:"type:varchar(255);not null"`
	MessageTime time.Time `gorm:"index:channel_message_idx"`
}
