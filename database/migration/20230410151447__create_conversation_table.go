package migration

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var V20230410151447 = gormigrate.Migration{
	ID: "V20230410151447",
	Migrate: func(tx *gorm.DB) error {
		type Conversation struct {
			gorm.Model
			ChannelID   string    `gorm:"type:varchar(255);not null; index:channel_message_idx" `
			Username    string    `gorm:"type:varchar(255);not null"`
			Content     string    `gorm:"type:varchar(255);not null"`
			MessageTime time.Time `gorm:"index:channel_message_idx"`
		}

		return tx.AutoMigrate(&Conversation{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("conversation")
	},
}
