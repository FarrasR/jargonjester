package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var V20230429024047 = gormigrate.Migration{
	ID: "V20230429024047",
	Migrate: func(tx *gorm.DB) error {
		type ChannelPermission struct {
			gorm.Model
			ChannelID string `gorm:"type:varchar(255);not null; index:channel_idx" `
			Active    bool
		}
		type UserPermission struct {
			gorm.Model
			UserID string `gorm:"type:varchar(255);not null; index:user_idx" `
			Active bool
		}

		return tx.AutoMigrate(&ChannelPermission{}, &UserPermission{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("channel_permissions", "user_permissions")
	},
}
